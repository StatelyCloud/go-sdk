package data

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/StatelyCloud/go-sdk/common/identifiers"
	pb "github.com/StatelyCloud/go-sdk/pb/data"
)

// NewTransaction starts a transaction and then hands transaction to your handler to preform all the logic necessary.
// If the handler returns an error, the transaction will be aborted. If the handler returns nil, the transaction will be committed.
func (s *store) NewTransaction(ctx context.Context, handler TransactionHandler) (*TransactionResults, error) {
	// Create a new transaction stream
	txn := &transaction{stream: s.client.Transaction(ctx)}
	defer func(txn *transaction) {
		// After we're done with the entire txn, close out the stream.
		// Do we want to bubble up errors closing out the stream?
		_ = txn.close()
	}(txn)

	// begin the transaction and hand it to the handler
	if err := txn.begin(s.storeID); err != nil {
		return nil, err
	}

	// hand the transaction stream to the handler and await for any errors.
	// in later versions, we may want to wrap this in a retryable handler
	if err := handler(txn); err != nil {
		// if there were errors, attempt to abort the transaction
		// Do we want to bubble up errors when aborting the txn?
		_ = txn.abort()
		return nil, err
	}

	// else commit the transaction and report the results back to the caller
	return txn.commit()
}

// Transaction represents a single transaction.
type transaction struct {
	stream *connect.BidiStreamForClient[pb.TransactionRequest, pb.TransactionResponse]
	done   atomic.Bool
	id     atomic.Uint32

	// putRequests and appendRequests are used to map responses to requests
	putRequests    []*parsedData[*structpb.Struct]
	appendRequests []*parsedData[*structpb.Struct]
}

func (t *transaction) begin(id identifiers.StoreID) error {
	return t.stream.Send(t.newTXNReq(&pb.TransactionRequest_Begin{
		Begin: &pb.TransactionBegin{StoreId: uint64(id)},
	}))
}

// close closes the stream for reading and writing.
func (t *transaction) close() error {
	if err := t.stream.CloseRequest(); err != nil {
		return err
	}
	if err := t.stream.CloseResponse(); err != nil {
		return err
	}
	return nil
}

// abort allows many threads to try to abort the transaction, but only 1 will issue the command once and only once.
func (t *transaction) abort() error {
	if !t.done.CompareAndSwap(false, true) {
		return nil // aborting an already closed transaction is a no-op
	}
	req := t.newTXNReq(&pb.TransactionRequest_Abort{
		Abort: &emptypb.Empty{},
	})

	err := t.stream.Send(req)
	if err != nil {
		return err
	}
	_, err = receiveExpected(t, req.MessageId, (*pb.TransactionResponse).GetFinished)
	return err
}

func (t *transaction) commit() (*TransactionResults, error) {
	if !t.done.CompareAndSwap(false, true) {
		return nil, connect.NewError(connect.CodeInternal, errors.New("this transaction was already closed"))
	}

	req := t.newTXNReq(&pb.TransactionRequest_Commit{
		Commit: &emptypb.Empty{},
	})
	err := t.stream.Send(req)
	if err != nil {
		return nil, err
	}

	resp, err := receiveExpected(t, req.MessageId, (*pb.TransactionResponse).GetFinished)
	if err != nil {
		return nil, err
	}

	return &TransactionResults{
		PutResponse:    mapPutResponses(resp.GetPutResults(), t.putRequests),
		AppendResponse: mapAppendResponses(resp.GetAppendResults(), t.appendRequests),
		DeleteResponse: mapDeleteResponse(resp.GetDeleteResults()),
		Committed:      resp.GetCommitted(),
	}, nil
}

// receiveExpected will either return the expected type or an error.
func receiveExpected[PT *T, T any](
	txn *transaction,
	msgID uint32,
	getter func(response *pb.TransactionResponse) PT,
) (PT, error) {
	resp, err := txn.stream.Receive()
	done := err != nil || resp.GetFinished() != nil
	if done {
		txn.done.Store(true)
	}
	if err != nil {
		return nil, err
	}

	if resp.MessageId != msgID {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf(
			"did not receive message... expected %T with ID: %d found, %T with ID %d",
			*new(PT), msgID, resp.GetResult(), resp.MessageId,
		))
	}

	v := getter(resp)
	if v == nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf(
			"did not receive expected type... expected %T found, %T", *new(PT), resp.GetResult(),
		))
	}
	return v, nil
}
