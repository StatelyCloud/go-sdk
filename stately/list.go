package stately

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// ListRequest starts a list operation.
type ListRequest struct {
	// KeyPathPrefix must be at least the root component but optionally contain
	// any number of path components to narrow your list result.
	// Example: [/state-washington]/city - the first segment is the root
	// component, this would return all cities in washington. Separately, you
	// could issue a list request with just `/state-washington` which would
	// return the item at "/state-washington"
	KeyPathPrefix string
}

// ListOptions are optional parameters for List.
type ListOptions struct {
	// Limit is the maximum number of items to return. If 0, the server will
	// default to unlimited.
	Limit uint32
	// SortableProperty is the property to sort by. Default is SortByKeyPath.
	SortableProperty SortableProperty
	// SortDirection is the direction to sort by. Default is Ascending.
	SortDirection SortDirection
}

// Merge combines two ListOptions into one. "other" takes precedence over "this".
// Nils will overwrite non-nil values.
func (lo *ListOptions) Merge(other *ListOptions) *ListOptions {
	if other == nil {
		return lo
	}
	if lo == nil {
		return other
	}
	lo.Limit = other.Limit
	lo.SortableProperty = other.SortableProperty
	lo.SortDirection = other.SortDirection
	return lo
}

// ContinueOptions are optional parameters for Continue.
type ContinueOptions struct {
	// SortDirection is the direction to sort by. Default is Ascending.
	SortDirection SortDirection
}

// SortableProperty is the property to sort by.
type SortableProperty int32

const (
	// SortByKeyPath sorts by the key path.
	SortByKeyPath SortableProperty = iota
	// SortByLastModifiedVersion sorts by the last time the item was modified.
	SortByLastModifiedVersion
)

// SortDirection is the direction to sort by.
type SortDirection int32

const (
	// Ascending is the default sort direction.
	Ascending SortDirection = iota
	// Descending is the reverse sort direction.
	Descending
)

// ListToken is a stateless token that acts like an iterator on a list of
// results efficiently fetching the next window. To fetch additional results,
// use the "next" token produced by Continue.
type ListToken struct {
	// Data will never be nil
	Data []byte
	// CanContinue indicates if there are more results to fetch using ContinueList
	CanContinue bool
	// CanSync indicates that you could call SyncList with this token later to
	// get updated items. This is determined by the type of store you're listing
	// from.
	CanSync bool
}

// newToken creates a new ListToken from a proto token and a store.
// If the proto token is nil, newToken will return nil.
func newToken(token *db.ListToken) *ListToken {
	if token == nil {
		return nil
	}
	return &ListToken{
		Data:        token.GetTokenData(),
		CanContinue: token.GetCanContinue(),
		CanSync:     token.GetCanSync(),
	}
}

// stream enables us to abstract txn and non-txn streams. For txn we need to
// verify messageId but for non-txn we don't.
type stream struct {
	// receive is a function that reads the next item from the stream. It
	// should return false if the stream is done.
	receive  func() bool
	response db.ListResponder
	err      error
}

type listIterator struct {
	stream     *stream
	itemMapper ItemTypeMapper

	// tracks where we're at in the current response page
	currPos  int
	currResp []*db.Item
	currItem Item

	// holds the final token and error to be returned by Token()
	finalToken *ListToken
	finalErr   error
}

// Next reads an item of the stream, and populates Item() with the current item.
// If there are no more items OR there is an error, Next will return false and
// the error will be available via Token().
func (li *listIterator) Next() bool {
	// if we're beyond the current response page, get the next page
	if li.currPos >= len(li.currResp) {

		// if the stream is done, return any possible error from the stream
		if !li.stream.receive() {
			li.finalErr = li.stream.err
			return false
		}

		switch v := li.stream.response.(type) {
		case *db.ListPartialResult:
			li.currResp = v.GetItems()
			li.currPos = 0
		case *db.ListFinished:
			li.finalToken = newToken(v.GetToken())
		}

	}
	// If we see the final token, we should stop
	if li.finalToken != nil {
		return false
	}

	// if we have items in the current response page, return the next one
	li.currItem, li.finalErr = li.itemMapper(li.currResp[li.currPos])
	li.currPos++

	// if there's no error, continue iterating
	return li.finalErr == nil
}

// Token returns the current token OR any error that occurred during iteration.
func (li *listIterator) Token() (*ListToken, error) {
	return li.finalToken, li.finalErr
}

// Value returns the current item in the iteration.
func (li *listIterator) Value() Item {
	return li.currItem
}

// ContinueList picks back up where this token left off. If there are no more
// results, `nil` will be returned. The default sort direction is Ascending if
// you do not specify ContinueOptions.
func (c *client) ContinueList(ctx context.Context, token []byte) (ListResponse[Item], error) {
	if token == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("token is nil"))
	}

	// call continue list
	response, err := c.client.ContinueList(ctx, connect.NewRequest(&db.ContinueListRequest{
		TokenData: token,
	}))
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     newStream(response),
		itemMapper: c.itemMapper,
	}, nil
}

// BeginList loads Items that start with a specified key path, subject to
// additional filtering. The prefix must minimally contain a Group Key (an
// item type and an item ID). BeginList will return an empty result set if
// there are no items matching that key prefix. A token is returned from this
// API that you can then pass to ContinueList to expand the result set, or to
// SyncList to get updates within the result set. This can fail if the caller
// does not have permission to read Items.
func (c *client) BeginList(
	ctx context.Context,
	keyPath string,
	opts ...ListOptions,
) (ListResponse[Item], error) {
	options := &ListOptions{}
	for _, opt := range opts {
		options = options.Merge(&opt)
	}

	response, err := c.client.BeginList(ctx, connect.NewRequest(&db.BeginListRequest{
		StoreId:         uint64(c.storeID),
		SchemaVersionId: uint32(c.schemaVersionID),
		KeyPathPrefix:   keyPath,
		AllowStale:      c.allowStale,
		Limit:           options.Limit,
		SortProperty:    db.SortableProperty(options.SortableProperty),
		SortDirection:   db.SortDirection(options.SortDirection),
	}))
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     newStream(response),
		itemMapper: c.itemMapper,
	}, nil
}

func newStream(response *connect.ServerStreamForClient[db.ListResponse]) *stream {
	newStream := &stream{}

	newStream.receive = func() bool {
		canContinue := response.Receive()
		newStream.response = response.Msg().GetListResponse()
		newStream.err = response.Err()
		return canContinue
	}
	return newStream
}
