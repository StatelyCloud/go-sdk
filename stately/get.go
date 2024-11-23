package stately

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

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

func (c *client) GetBatch(ctx context.Context, itemPaths ...string) ([]Item, error) {
	response, err := c.client.Get(ctx, connect.NewRequest(&db.GetRequest{
		StoreId:         uint64(c.storeID),
		SchemaVersionId: uint32(c.schemaVersionID),
		Gets:            mapToItemKey(itemPaths),
		AllowStale:      c.allowStale,
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
