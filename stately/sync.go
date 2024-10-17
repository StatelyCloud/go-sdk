package stately

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// SyncResponse is a response from a sync operation.
type SyncResponse interface {
	// IsSyncResponse is a marker method to indicate that a type is a SyncResponse.
	IsSyncResponse()
}

// Changed is a SyncResponse that indicates that the item was changed.
type Changed struct {
	Item Item
}

// Deleted is a SyncResponse that indicates that the item was deleted.
type Deleted struct {
	KeyPath string
}

// UpdateOutsideOfWindow is a SyncResponse containing items that were updated
// but Stately cannot tell if they were in the sync window. Treat these as
// deleted in most cases. For more information see:
// https://docs.stately.cloud/api/sync
type UpdateOutsideOfWindow struct {
	KeyPath string
}

// Reset is a SyncResponse that indicates that the sync operation should be reset.
type Reset struct{}

// IsSyncResponse is a marker method to indicate that a type is a SyncResponse.
func (r *Changed) IsSyncResponse() {}

// IsSyncResponse is a marker method to indicate that a type is a SyncResponse.
func (r *Deleted) IsSyncResponse() {}

// IsSyncResponse is a marker method to indicate that a type is a SyncResponse.
func (r *UpdateOutsideOfWindow) IsSyncResponse() {}

// IsSyncResponse is a marker method to indicate that a type is a SyncResponse.
func (r *Reset) IsSyncResponse() {}

// SyncList returns an iterator for a sync operation.
// You should use the iterator to get the results of the sync operation.
// For example:
//
//	iter, err := store.SyncList(ctx, token)
//	for iter.Continue() {
//	  switch v := iter.Next().(type) {
//	  case *stately.Changed:
//	    // do something with the changed item: v.Item
//	  case *stately.Deleted:
//	    // do something with removed key path: v.KeyPath
//	  case *stately.UpdateOutsideOfWindow:
//	    // do something with the out of window update: v.KeyPath
//	  case *stately.Reset:
//	    // reset the sync operation
//	  }
//	}
//	err, token := iter.Token()
//	// handle error and token
func (c *client) SyncList(ctx context.Context, token []byte) (ListResponse[SyncResponse], error) {
	resp, err := c.client.SyncList(ctx, connect.NewRequest(&db.SyncListRequest{
		TokenData: token,
	}))
	if err != nil {
		return nil, err
	}

	return &syncIterator{stream: resp, itemMapper: c.itemMapper}, nil
}

type syncIterator struct {
	stream *connect.ServerStreamForClient[db.SyncListResponse]
	// readNext allows us to set state that we want to pull more messages off the stream
	readNext bool

	// bookkeeping for handling partial responses
	partialResponsePos int
	partialResponseLen int
	partialResponse    *db.SyncListPartialResponse

	currValue  SyncResponse
	itemMapper ItemTypeMapper

	// final state of the iterator
	finalToken *ListToken
	finalErr   error
}

func (s *syncIterator) Next() bool {
	// if we're at the end of the current response page, and there's no final
	// token, we want to read the next message.
	if s.partialResponsePos >= s.partialResponseLen && s.finalToken == nil {
		s.readNext = true
	}

	// if we're beyond the current response page, get the next page
	if s.readNext {
		s.readNext = false
		// if the stream is done, return any possible error from the stream
		if !s.stream.Receive() {
			s.finalErr = s.stream.Err()
			return false
		}

		switch v := s.stream.Msg().GetResponse().(type) {
		case *db.SyncListResponse_Reset_:
			s.currValue = &Reset{}
			// after a reset message, we want to read the next message
			s.readNext = true
			// reset the partial response state, so we want to continue
			return true
		case *db.SyncListResponse_Result:
			// reset the partial response state
			s.partialResponsePos = 0
			s.partialResponseLen = len(v.Result.GetChangedItems()) + len(v.Result.GetDeletedItems()) + len(v.Result.GetUpdatedItemKeysOutsideListWindow())
			s.partialResponse = v.Result
		case *db.SyncListResponse_Finished:
			// terminal state
			s.currValue = nil // nil so if they call iter.Next() they'll get nil and have to call iter.Token()
			s.finalToken = newToken(v.Finished.GetToken())
			return false
		}
	}

	// handle iterating through the current response
	pos := s.partialResponsePos
	s.partialResponsePos++
	changeNum := len(s.partialResponse.GetChangedItems())
	deleteNum := len(s.partialResponse.GetDeletedItems())
	updateNum := len(s.partialResponse.GetUpdatedItemKeysOutsideListWindow())

	if pos < changeNum {
		item, err := s.itemMapper(s.partialResponse.GetChangedItems()[pos])
		if err != nil {
			s.finalErr = err
			return false
		}
		s.currValue = &Changed{Item: item}
		return true
	}
	pos -= changeNum
	if pos < deleteNum {
		s.currValue = &Deleted{KeyPath: s.partialResponse.GetDeletedItems()[pos].GetKeyPath()}
		return true
	}
	pos -= deleteNum
	// we probably don't need this last statement "if" statement but it's here for clarity
	if pos < updateNum {
		s.currValue = &UpdateOutsideOfWindow{
			KeyPath: s.partialResponse.GetUpdatedItemKeysOutsideListWindow()[pos],
		}
		return true
	}

	return true
}

func (s *syncIterator) Token() (*ListToken, error) {
	return s.finalToken, s.finalErr
}

func (s *syncIterator) Value() SyncResponse {
	return s.currValue
}
