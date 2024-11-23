package stately

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/db"
)

// ListRequest starts a list operation.
type ListRequest struct {
	// KeyPathPrefix must contain at least a full group key but can optionally
	// contain any number of path components to narrow your list result. Example:
	// [/state-washington]/city - the first segment is the group key, and this
	// would return all cities in washington. Separately, you could issue a list
	// request with just `/state-washington` which would also return the item at
	// "/state-washington"
	KeyPathPrefix string
}

// ListOptions are optional parameters for List.
type ListOptions struct {
	// Limit is the maximum number of items to return. The default is unlimited -
	// all items will be returned.
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

// ListToken is a stateless token that saves your place in a result set,
// allowing you to fetch additional results with ContinueList, or get updated
// results with SyncList.
type ListToken struct {
	// Data will never be nil. This is the token data that you pass to
	// ContinueList or SyncList.
	Data []byte
	// CanContinue indicates if there are more results to fetch using
	// ContinueList.
	CanContinue bool
	// CanSync indicates that you could call SyncList with this token later to get
	// updated items. This is determined by the type of store you're listing from.
	CanSync bool

	// SchemaVersionID is the schema version ID of the store that produced this token.
	// When making ContinueList calls, ensure your client version uses tokens that
	// match this field. Using tokens with different schema versions will result
	// in a SchemaVersionMismatch error. For SyncList calls, you only need to
	// ensure you handle Reset responses correctly.
	SchemaVersionID SchemaVersionID
}

// newToken creates a new ListToken from a proto token and a store.
// If the proto token is nil, newToken will return nil.
func newToken(token *db.ListToken) *ListToken {
	if token == nil {
		return nil
	}
	return &ListToken{
		Data:            token.GetTokenData(),
		CanContinue:     token.GetCanContinue(),
		CanSync:         token.GetCanSync(),
		SchemaVersionID: SchemaVersionID(token.GetSchemaVersionId()),
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

func (c *client) ContinueList(ctx context.Context, token []byte) (ListResponse[Item], error) {
	if token == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("token is nil"))
	}

	// call continue list
	response, err := c.client.ContinueList(ctx, connect.NewRequest(&db.ContinueListRequest{
		TokenData:       token,
		SchemaVersionId: uint32(c.schemaVersionID),
	}))
	if err != nil {
		return nil, err
	}

	return &listIterator{
		stream:     newStream(response),
		itemMapper: c.itemMapper,
	}, nil
}

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
