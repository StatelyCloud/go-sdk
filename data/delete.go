package data

import (
	"context"

	"connectrpc.com/connect"

	pbdata "github.com/StatelyCloud/go-sdk/pb/data"
)

// DeleteRequest one or more items to from the Group.
type DeleteRequest struct {
	ItemPaths []string
}

// NewDeleteRequest is a convenience method to construct DeleteRequest with single or more items vs vanilla golang:
// DeleteRequest{ItemPaths: []KeyPath{}{"/message-1"}}.
func NewDeleteRequest(itemPaths ...string) DeleteRequest {
	return DeleteRequest{
		ItemPaths: itemPaths,
	}
}

// Delete is a convenience method for removing a single Item from the Store by its full key path.
// See DeleteBatch for more information.
func (c *dataClient) Delete(ctx context.Context, itemPath string) error {
	return c.DeleteBatch(ctx, DeleteRequest{
		ItemPaths: []string{itemPath},
	})
}

// DeleteBatch removes one or more Items from the Store by their full key paths. This
// will fail if any Item does not exist, if not all of the DeleteItem requests
// are under the same root item path, or if the caller does not have permission
// to delete Items. Tombstones will be left for deleted items for some
// predetermined time (TBD tombstone behavior). All deletes in the request are
// applied atomically - there are no partial successes.
func (c *dataClient) DeleteBatch(ctx context.Context, request DeleteRequest) error {
	_, err := c.client.Delete(ctx, connect.NewRequest(&pbdata.DeleteRequest{
		StoreId: uint64(c.storeID),
		Deletes: mapDeleteRequest(request.ItemPaths),
	}))
	return err
}

func mapDeleteResponse(results []*pbdata.DeleteResult) []string {
	deleteResponses := make([]string, len(results))
	for idx, result := range results {
		deleteResponses[idx] = result.KeyPath
	}
	return deleteResponses
}

func mapDeleteRequest(keyPaths []string) []*pbdata.DeleteItem {
	deleteItems := make([]*pbdata.DeleteItem, len(keyPaths))
	for i, v := range keyPaths {
		deleteItems[i] = &pbdata.DeleteItem{KeyPath: v}
	}
	return deleteItems
}
