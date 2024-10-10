package stately

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// Put adds one Item to the Store, or replaces the Item if it
// already exists at that path.
//
// This will fail if:
//   - The caller does not have permission to create Items.
func (c *client) Put(ctx context.Context, item Item) (Item, error) {
	responses, err := c.PutBatch(ctx, item)
	if err != nil {
		return nil, err
	}
	return responses[0], nil
}

// PutBatch adds up to 50 Items to the Store, or replaces the Items if they
// already exist at that path.
//
// This will fail if
//   - the caller does not have permission to create Items.
//
// Additional Notes:
// All puts in the request are applied atomically - there are no partial
// successes.
func (c *client) PutBatch(ctx context.Context, batch ...Item) ([]Item, error) {
	putItems, err := mapPutRequest(batch)
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

	return batch, mapPutResponses(resp.Msg.GetItems(), batch)
}

// Shared between transactional and non-transactional put.
func mapPutRequest(batchRequest []Item) ([]*db.PutItem, error) {
	// Build the put items
	putItems := make([]*db.PutItem, len(batchRequest))
	for i, v := range batchRequest {

		item, err := v.MarshalStately()
		if err != nil {
			return nil, err
		}
		putItems[i] = &db.PutItem{
			Item: item,
		}
	}
	return putItems, nil
}

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
