package stately

import (
	"context"
	"strconv"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
	"github.com/StatelyCloud/go-sdk/sdkerror"
)

// WithPutOptions wraps an item and adds options for use when putting the item.
// This may be used in place of an Item in Put or PutBatch.
type WithPutOptions struct {
	Item
	// MustNotExist is a condition that indicates this item must not already exist
	// at any of its key paths. If there is already an item at one of those paths,
	// the Put operation will fail with a ConditionalCheckFailed error. Note that
	// if the item has an `initialValue` field in its key, that initial value will
	// automatically be chosen not to conflict with existing items, so this
	// condition only applies to key paths that do not contain the `initialValue`
	// field.
	MustNotExist bool
}

func (c *client) Put(ctx context.Context, item Item) (Item, error) {
	responses, err := c.PutBatch(ctx, item)
	if err != nil {
		return nil, err
	}
	return responses[0], nil
}

func (c *client) PutBatch(ctx context.Context, batch ...Item) ([]Item, error) {
	items, putItems, err := mapPutRequestWithOptions(batch)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Put(ctx, connect.NewRequest(&db.PutRequest{
		StoreId:         uint64(c.storeID),
		SchemaVersionId: uint32(c.schemaVersionID),
		Puts:            putItems,
	}))
	if err != nil {
		return nil, err
	}

	return items, mapPutResponses(resp.Msg.GetItems(), items)
}

// mapPutRequestWithOptions maps a list of items or WithOptions to a list of
// Items and a list of PutItem inputs.
// Shared between transactional and non-transactional put.
func mapPutRequestWithOptions(itemsOrOptions []Item) (items []Item, putItems []*db.PutItem, err error) {
	items = make([]Item, len(itemsOrOptions))
	putItems = make([]*db.PutItem, len(itemsOrOptions))
	for i, v := range itemsOrOptions {
		var withOptions WithPutOptions
		if po, ok := v.(WithPutOptions); ok {
			withOptions = po
		} else {
			withOptions = WithPutOptions{
				Item: v,
			}
		}
		if withOptions.Item == nil {
			return nil, nil, &sdkerror.Error{
				Code:        connect.CodeInvalidArgument,
				StatelyCode: "ItemIsRequired",
				Message:     "items[" + strconv.Itoa(i) + "] is nil",
			}
		}
		item, err := withOptions.Item.MarshalStately()
		if err != nil {
			return nil, nil, err
		}
		items[i] = withOptions.Item
		putItems[i] = &db.PutItem{
			Item:         item,
			MustNotExist: withOptions.MustNotExist,
		}
	}
	return items, putItems, err
}

// Replace the list of items in place with the results of the put operation.
// This is needed because UnmarshalStately needs an item of the correct type to
// unmarshal into.
// shared between transactional and non-transactional put.
func mapPutResponses(results []*db.Item, original []Item) error {
	if results == nil {
		return nil
	}
	// map the results back
	for idx, result := range results {
		if err := original[idx].UnmarshalStately(result); err != nil {
			return err
		}
	}
	return nil
}
