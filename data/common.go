package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	pbdata "github.com/StatelyCloud/go-sdk/pb/data"
)

// parsedData is a simple tuple type to hold both the raw JSON data and the parsed item.
type parsedData[T any] struct {
	jsonData   json.RawMessage
	jsonParsed T
}

func (d *parsedData[T]) getJSONParsed() T {
	if d == nil {
		return *new(T)
	}
	return d.jsonParsed
}

func (d *parsedData[T]) getJSONData() json.RawMessage {
	if d == nil {
		return nil
	}
	return d.jsonData
}

// setKeyPath is a helper function to take a string, parse it and set the key path on the item.
func setKeyPath(keyPath string, item *RawItem) (*RawItem, error) {
	if item == nil {
		item = &RawItem{}
	}

	ik, err := toItemKey(keyPath)
	if err != nil {
		return nil, err
	}

	lastPathPart := ik[len(ik)-1]

	item.Key = ik
	item.ParentKeyPath = ik[:len(ik)-1]
	item.ID = lastPathPart.ItemID
	item.ItemType = lastPathPart.ItemType

	return item, nil
}

func setProtoMetadata(md *pbdata.ItemMetadata, item *RawItem) *RawItem {
	if item == nil {
		item = &RawItem{}
	}

	item.Metadata = ItemMetadata{
		CreatedAt:             time.UnixMicro(int64(md.GetCreatedAtMicros())),
		LastModifiedAt:        time.UnixMicro(int64(md.GetLastModifiedAtMicros())),
		CreatedAtVersion:      md.GetCreatedAtVersion(),
		LastModifiedAtVersion: md.GetLastModifiedAtVersion(),
	}
	return item
}

// protoToItem is for the egress boundary of our api/client. We allow users to output struct for json or
// their well-formed proto message. This will fail if T is a string/[]byte as we discourage that usage.
func protoToItem(protoItem *pbdata.Item) (*RawItem, error) {
	var err error
	var item *RawItem

	if item, err = setKeyPath(protoItem.GetKeyPath(), item); err != nil {
		return nil, err
	}

	item = setProtoMetadata(protoItem.GetMetadata(), item)

	// nil proto item or json will result in '{}'
	jsonData, err := protoItem.GetJson().MarshalJSON()
	if err != nil {
		return nil, err
	}
	item.JSONData = jsonData

	return item, nil
}

// dataToProto is a convenience method for mapping from any type to a json/proto or in the future json+proto.
func dataToProto(data any) (*parsedData[*structpb.Struct], []byte, error) {
	// Initially test if it's a proto message
	if msg, isProto := data.(proto.Message); isProto {
		// okay we have a proper proto message
		protoBytes, err := proto.Marshal(msg)
		if err != nil {
			return nil, nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid proto message: %v", err))
		}
		return nil, protoBytes, nil
	}

	// Else it's a json struct, we'll cover some basic types, but we have no way to just type switch on "XXX struct" so
	// we just have to let the json unmarshaller let us know if it's valid.
	jsonStruct := &parsedData[*structpb.Struct]{jsonParsed: &structpb.Struct{}}
	switch v := data.(type) {
	case string:
		return nil, nil, connect.NewError(connect.CodeInvalidArgument,
			errors.New("don't use strings as your data, please use map[string]any or a json annotated struct"))
	case []byte:
		return nil, nil, connect.NewError(connect.CodeInvalidArgument,
			errors.New("don't use []byte as your data, please use map[string]any or a json annotated struct"))
	default:
		// default we assume json?
		jsonData, err := json.Marshal(v)
		jsonStruct.jsonData = jsonData
		if err != nil {
			return nil, nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid json object: %v", err))
		}
		if err = json.Unmarshal(jsonData, &jsonStruct.jsonParsed); err != nil {
			return nil, nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid json []byte: %v", err))
		}
	}

	return jsonStruct, nil, nil
}

// UnmarshalItem is a helper function to unmarshal a list of items of a specific type.
// The type of your object is the string representation of the type, e.g. "user", "todo", etc.
// Example: /list-todo/task-4 your type is "task".
func UnmarshalItem[T any](itemType string, rawItem *RawItem) (*Item[T], error) {
	if rawItem == nil {
		return nil, nil
	}
	if rawItem.ItemType != itemType {
		return nil, nil
	}
	item := &Item[T]{RawItem: rawItem}
	if err := json.Unmarshal(item.JSONData, &item.Data); err != nil {
		return nil, err
	}
	return item, nil
}

// UnmarshalItemList is a helper function to unmarshal a list of items of a specific type.
// The type of your object is the string representation of the type, e.g. "user", "todo", etc.
// Example: /list-todo/task-4 your type is "task".
// If you want to unmarshal all items, pass "*" as the itemType.
func UnmarshalItemList[T any](itemType string, rawItems ...*RawItem) ([]*Item[T], error) {
	//nolint:prealloc // we want the ability to return a nil list
	var items []*Item[T]
	for _, rawItem := range rawItems {
		if itemType != "*" && rawItem.ItemType != itemType {
			continue
		}
		item := &Item[T]{RawItem: rawItem}
		if err := json.Unmarshal(item.JSONData, &item.Data); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
