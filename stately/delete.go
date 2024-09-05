package stately

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// Delete removes one or more Items from the Store by their full key
// paths. This will fail if any Item does not exist, if not all the
// DeleteItem requests are under the same root item path, or if the caller does
// not have permission to delete Items. Tombstones will be left for deleted
// items for some predetermined time (TBD tombstone behavior). All deletes in
// the request are applied atomically - there are no partial successes.
func (c *client) Delete(ctx context.Context, itemPaths ...string) error {
	_, err := c.client.Delete(ctx, connect.NewRequest(&db.DeleteRequest{
		StoreId: uint64(c.storeID),
		Deletes: mapDeleteRequest(itemPaths),
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
