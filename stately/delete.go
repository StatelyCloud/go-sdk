package stately

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

func (c *client) Delete(ctx context.Context, itemPaths ...string) error {
	_, err := c.client.Delete(ctx, connect.NewRequest(&db.DeleteRequest{
		StoreId:         uint64(c.storeID),
		Deletes:         mapDeleteRequest(itemPaths),
		SchemaVersionId: uint32(c.schemaVersionID),
	}))
	return err
}

func mapDeleteResponse(results []*db.DeleteResult) []string {
	deleteResponses := make([]string, len(results))
	for idx, result := range results {
		deleteResponses[idx] = result.KeyPath
	}
	return deleteResponses
}

func mapDeleteRequest(keyPaths []string) []*db.DeleteItem {
	deleteItems := make([]*db.DeleteItem, len(keyPaths))
	for i, v := range keyPaths {
		deleteItems[i] = &db.DeleteItem{KeyPath: v}
	}
	return deleteItems
}
