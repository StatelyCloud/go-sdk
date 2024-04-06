package data

import (
	"context"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/StatelyCloud/go-sdk/common/types"
	pb "github.com/StatelyCloud/go-sdk/pb/data"
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
func (s *store) Put(ctx context.Context, path string, value any) (*RawItem, error) {
	responses, err := s.PutBatch(ctx, PutBatchRequest{
		PutData: []*PutData{
			{
				KeyPath: path,
				Data:    value,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	response := responses[0]
	return response.RawItem, response.Error
}

// PutBatchRequest allows you to specify items to put either atomically or non-atomic. Atomic operations are more expensive
// in terms of io/latency/cost.
type PutBatchRequest struct {
	PutData []*PutData
	// (option) Atomic indicates that all puts must succeed or none will (i.e. that they
	// are applied in a transaction), and that other operations will be serialized
	// ahead or behind this operation.
	Atomic IsAtomic
}

// NewPutBatchRequest is a convenience method to generate a new PutBatchRequest.
func NewPutBatchRequest[T any](putRequests ...*PutData) PutBatchRequest {
	return PutBatchRequest{
		PutData: putRequests,
	}
}

// PutBatchResponse is a simple tuple of either RawItem or Error.
type PutBatchResponse struct {
	*RawItem
	Error error
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
func (s *store) PutBatch(ctx context.Context, batchRequest PutBatchRequest) ([]*PutBatchResponse, error) {
	putItems, originalItems, err := mapPutRequest(batchRequest.PutData)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Put(ctx, connect.NewRequest(&pb.PutRequest{
		StoreId: uint64(s.storeID),
		Puts:    putItems,
		Atomic:  bool(batchRequest.Atomic),
	}))
	if err != nil {
		return nil, err
	}

	responses := mapPutResponses(resp.Msg.GetResults(), originalItems)

	return responses, nil
}

// shared between transactional and non-transactional put.
func mapPutRequest(batchRequest []*PutData) ([]*pb.PutItem, []*parsedData[*structpb.Struct], error) {
	// Build the put items
	putItems := make([]*pb.PutItem, len(batchRequest))
	originalItems := make([]*parsedData[*structpb.Struct], len(batchRequest))
	for i, v := range batchRequest {

		item := &pb.Item{
			KeyPath: v.KeyPath,
		}

		jsonData, protoData, err := dataToProto(v.Data)
		if err != nil {
			return nil, nil, err
		}

		originalItems[i] = jsonData

		item.Json = jsonData.getJSONParsed()
		item.Proto = protoData

		putItems[i] = &pb.PutItem{
			Item: item,
		}
	}
	return putItems, originalItems, nil
}

// shared between transactional and non-transactional put.
func mapPutResponses(results []*pb.PutResult, originalItems []*parsedData[*structpb.Struct]) []*PutBatchResponse {
	if results == nil {
		return nil
	}
	// map the results back
	responses := make([]*PutBatchResponse, len(results))
	for idx, result := range results {

		// Handle err first
		if err := result.GetError(); err != nil {
			responses[idx] = &PutBatchResponse{Error: types.MapProtoError(err)}
			continue
		}

		item := &RawItem{
			JSONData: originalItems[idx].getJSONData(),
		}

		item = setProtoMetadata(result.GetMetadata(), item)
		item, err := setKeyPath(result.GetKeyPath(), item)
		if err != nil {
			responses[idx] = &PutBatchResponse{Error: err}
			continue
		}

		responses[idx] = &PutBatchResponse{RawItem: item}
	}
	return responses
}
