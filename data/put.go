package data

import (
	"context"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/structpb"

	pbdata "github.com/StatelyCloud/go-sdk/pb/data"
)

// PutData allows you to fill in any type of data at a desired KeyPath.
type PutData struct {
	KeyPath string
	Data    any
}

// Put adds one Item to the Store, or replaces the Item if it
// already exists at that path.
//
// This will fail if:
//   - The PutData is not under the same root item path.
//   - The PutData write conditions fails.
//   - The caller does not have permission to create Items.
//
// Additional Notes:
// The PutData is applied atomically - there are no partial
// successes. Data can be provided as a struct that can be
// marshalled/unmarshalled as JSON, or a struct that can be
// serialized/deserialized as a proto.
func (c *dataClient) Put(ctx context.Context, path string, value any) (*RawItem, error) {
	responses, err := c.PutBatch(ctx, &PutData{
		KeyPath: path,
		Data:    value,
	})
	if err != nil {
		return nil, err
	}
	return responses[0], nil
}

// PutBatch adds one or more Items to the Store, or replaces the Items if they
// already exist at that path.
//
// This will fail if
//   - not all the PutData requests are under the same root item path.
//   - any of the PutData requests' write conditions fails.
//   - the caller does not have permission to create Items.
//
// Additional Notes:
// All puts in the request are applied atomically - there are no partial
// successes. Data can be provided as either JSON, or as a proto encoded by a
// previously agreed upon schema, or by some combination of the two.
func (c *dataClient) PutBatch(ctx context.Context, batch ...*PutData) ([]*RawItem, error) {
	putItems, originalItems, err := mapPutRequest(batch)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Put(ctx, connect.NewRequest(&pbdata.PutRequest{
		StoreId: uint64(c.storeID),
		Puts:    putItems,
	}))
	if err != nil {
		return nil, err
	}

	return mapPutResponses(resp.Msg.GetResults(), originalItems)
}

// shared between transactional and non-transactional put.
func mapPutRequest(batchRequest []*PutData) ([]*pbdata.PutItem, []*parsedData[*structpb.Struct], error) {
	// Build the put items
	putItems := make([]*pbdata.PutItem, len(batchRequest))
	originalItems := make([]*parsedData[*structpb.Struct], len(batchRequest))
	for i, v := range batchRequest {

		item := &pbdata.Item{
			KeyPath: v.KeyPath,
		}

		jsonData, protoData, err := dataToProto(v.Data)
		if err != nil {
			return nil, nil, err
		}

		originalItems[i] = jsonData

		item.Json = jsonData.getJSONParsed()
		item.Proto = protoData

		putItems[i] = &pbdata.PutItem{
			Item: item,
		}
	}
	return putItems, originalItems, nil
}

// shared between transactional and non-transactional put.
func mapPutResponses(results []*pbdata.PutResult, originalItems []*parsedData[*structpb.Struct]) ([]*RawItem, error) {
	if results == nil {
		return nil, nil
	}
	// map the results back
	responses := make([]*RawItem, len(results))
	for idx, result := range results {
		item := &RawItem{
			JSONData: originalItems[idx].getJSONData(),
		}

		item = setProtoMetadata(result.GetMetadata(), item)
		item, err := setKeyPath(result.GetKeyPath(), item)
		if err != nil {
			return nil, err
		}
		responses[idx] = item
	}
	return responses, nil
}
