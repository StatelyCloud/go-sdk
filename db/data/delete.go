package data

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/common/types"
	pb "github.com/StatelyCloud/go-sdk/pb/data"
)

// DeleteRequest one or more items to from the Group.
type DeleteRequest struct {
	ItemPaths []string
	// (option) Atomic indicates that all deletes must succeed or none will (i.e. that they
	// are applied in a transaction), and that other operations will be serialized
	// ahead or behind this operation.
	Atomic IsAtomic
}

// DeleteResponse indicates if a specific KeyPath failed to delete along with an error indicating why.
type DeleteResponse struct {
	KeyPath string
	Error   error
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
func (s *store) Delete(ctx context.Context, itemPath string) error {
	deletes, err := s.DeleteBatch(ctx, DeleteRequest{
		ItemPaths: []string{itemPath},
	})
	if err != nil {
		return err
	}
	return deletes[0].Error
}

// DeleteBatch removes one or more Items from the Store by their full key paths. This
// will fail if any Item does not exist, if not all of the DeleteItem requests
// are under the same root item path, or if the caller does not have permission
// to delete Items. Tombstones will be left for deleted items for some
// predetermined time (TBD tombstone behavior). All deletes in the request are
// applied atomically - there are no partial successes.
func (s *store) DeleteBatch(ctx context.Context, request DeleteRequest) ([]*DeleteResponse, error) {
	resp, err := s.client.Delete(ctx, connect.NewRequest(&pb.DeleteRequest{
		StoreId: uint64(s.storeID),
		Deletes: mapDeleteRequest(request.ItemPaths),
		Atomic:  bool(request.Atomic),
	}))
	if err != nil {
		return nil, err
	}

	return mapDeleteResponse(resp.Msg.GetResults()), nil
}

func mapDeleteResponse(results []*pb.DeleteResult) []*DeleteResponse {
	deleteResponses := make([]*DeleteResponse, len(results))
	for idx, result := range results {
		deleteResponses[idx] = &DeleteResponse{
			KeyPath: result.GetKeyPath(),
			Error:   types.MapProtoError(result.GetError()),
		}
	}
	return deleteResponses
}

func mapDeleteRequest(keyPaths []string) []*pb.DeleteItem {
	deleteItems := make([]*pb.DeleteItem, len(keyPaths))
	for i, v := range keyPaths {
		deleteItems[i] = &pb.DeleteItem{KeyPath: v}
	}
	return deleteItems
}
