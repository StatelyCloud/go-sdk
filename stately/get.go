package stately

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// Get is a convenience method for retrieving a single Item by its full key path.
// See GetBatch for more information.
func (c *client) Get(ctx context.Context, itemPath string) (Item, error) {
	items, err := c.GetBatch(ctx, itemPath)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], nil
}

// GetBatch retrieves one or more Items by their full key paths. This will
// return any of the Items that exist. It will fail if not all the GetItem
// requests are under the same root item path, or if the caller does not have
// permission to read Items. Use Query if you want to retrieve multiple items
// but don't already know the full key paths of the items you want to get.
func (c *client) GetBatch(ctx context.Context, itemPaths ...string) ([]Item, error) {
	response, err := c.client.Get(ctx, connect.NewRequest(&db.GetRequest{
		StoreId:    uint64(c.storeID),
		Gets:       mapToItemKey(itemPaths),
		AllowStale: c.allowStale,
	}))
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(response.Msg.GetItems()))
	for i, v := range response.Msg.Items {
		items[i], err = c.itemMapper(v)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func mapToItemKey(keys []string) []*db.GetItem {
	getItems := make([]*db.GetItem, len(keys))
	for i, v := range keys {
		getItems[i] = &db.GetItem{KeyPath: v}
	}
	return getItems
}
