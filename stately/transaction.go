package stately

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync/atomic"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// NewTransaction starts a transaction and then hands transaction to your
// handler to preform all the logic necessary. If the handler returns an
// error, the transaction will be aborted. If the handler returns nil, the
// transaction will be committed.
func (c *client) NewTransaction(ctx context.Context, handler TransactionHandler) (*TransactionResults, error) {
	// Create a new transaction stream
	txn := &transaction{stream: c.client.Transaction(ctx), itemMapper: c.itemMapper}
	defer func(txn *transaction) {
		// After we're done with the entire txn, close out the stream.
		// Do we want to bubble up errors closing out the stream?
		_ = txn.close()
	}(txn)

	// begin the transaction and hand it to the handler
	if err := txn.begin(c.storeID); err != nil {
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
	stream     *connect.BidiStreamForClient[db.TransactionRequest, db.TransactionResponse]
	done       atomic.Bool
	id         atomic.Uint32
	itemMapper ItemTypeMapper

	// putRequests are used to map responses to requests
	putRequests []Item
}

func (t *transaction) begin(id StoreID) error {
	return t.safeSend(t.newTXNReq(&db.TransactionRequest_Begin{
		Begin: &db.TransactionBegin{StoreId: uint64(id)},
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

// abort allows many threads to try to abort the transaction, but only 1 will
// issue the command once and only once.
func (t *transaction) abort() error {
	if !t.done.CompareAndSwap(false, true) {
		return nil // aborting an already closed transaction is a no-op
	}
	req := t.newTXNReq(&db.TransactionRequest_Abort{
		Abort: &emptypb.Empty{},
	})

	err := t.safeSend(req)
	if err != nil {
		return err
	}
	_, err = receiveExpected(t, req.MessageId, (*db.TransactionResponse).GetFinished)
	return err
}

func (t *transaction) commit() (*TransactionResults, error) {
	if !t.done.CompareAndSwap(false, true) {
		return nil, connect.NewError(connect.CodeInternal, errors.New("this transaction was already closed"))
	}

	req := t.newTXNReq(&db.TransactionRequest_Commit{
		Commit: &emptypb.Empty{},
	})
	err := t.safeSend(req)
	if err != nil {
		return nil, err
	}

	resp, err := receiveExpected(t, req.MessageId, (*db.TransactionResponse).GetFinished)
	if err != nil {
		return nil, err
	}
	err = mapPutResponses(resp.GetPutResults(), t.putRequests)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &TransactionResults{
		PutResponse:    t.putRequests,
		DeleteResponse: mapDeleteResponse(resp.GetDeleteResults()),
		Committed:      resp.GetCommitted(),
	}, nil
}

// It is possible to call Send on a closed stream, this usually means an error
// was returned by the server and is stashed in the "Receive" method so
// all sends should check for this.
func (t *transaction) safeSend(req *db.TransactionRequest) error {
	err := t.stream.Send(req)
	if errors.Is(err, io.EOF) { // EOF is the error returned when the stream is closed.
		_, err := t.stream.Receive()
		if err != nil {
			return err
		}
	}
	return err // otherwise return the original.
}

// receiveExpected will either return the expected type or an error.
func receiveExpected[PT *T, T any](
	txn *transaction,
	msgID uint32,
	getter func(response *db.TransactionResponse) PT,
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
