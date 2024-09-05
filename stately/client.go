package stately

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/planetscale/vtprotobuf/codec/grpc"

	"github.com/StatelyCloud/go-sdk/pb/db/dbconnect"
	"github.com/StatelyCloud/go-sdk/sdkerror"
)

type client struct {
	client     dbconnect.DatabaseServiceClient
	storeID    StoreID
	itemMapper ItemTypeMapper

	// allowStale indicates whether operations like get and list can return
	// slightly stale data.
	allowStale bool
}

// Client is a stately client that interacts with the given store.
type Client interface {
	// WithAllowStale returns a new client with the given allowStale value set. By
	// default clients do not allow stale reads, but this method can used to
	// create a lightweight copy of the client where reads can show stale data.
	WithAllowStale(allowStale bool) Client

	// GetBatch retrieves one or more Items by their full key paths. This will
	// return any of the Items that exist. It will fail if the caller does not
	// have permission to read Items. Use BeginList if you want to retrieve
	// multiple items but don't already know the full key paths of the items
	// you want to get.
	GetBatch(ctx context.Context, itemPaths ...string) ([]Item, error)

	// Get retrieves one Item by their full key paths. It will fail if the
	// caller does not have permission to read Items. Use BeginList if you want
	// to retrieve multiple items but don't already know the full key paths of
	// the items you want to get. Use GetBatch if you want to retrieve multiple
	// items by their full key paths.
	Get(ctx context.Context, itemPath string) (Item, error)

	// PutBatch adds one or more Items to the Store, or replaces the Items if
	// they already exist at that path.
	// This will fail if
	//   - the caller does not have permission to write to the underlying store.
	//
	// Additional Notes:
	// All puts in the request are applied atomically - there are no partial
	// successes.
	PutBatch(ctx context.Context, batch ...Item) ([]Item, error)

	// Put adds one Item to the Store, or replaces the Item if it
	// already exists at that path.
	//
	// This will fail if:
	//   - The caller does not have permission to create Items.
	//
	Put(ctx context.Context, item Item) (Item, error)

	// Delete removes one or more Items from the Store by their full key
	// paths. This will fail if the caller does not have permission to delete
	// Items. Tombstones will be left for deleted items for some predetermined
	// time (TBD tombstone behavior). All deletes in the request are applied
	// atomically - there are no partial successes.
	Delete(ctx context.Context, itemPaths ...string) error

	// BeginList loads Items that start with a specified key path, subject to
	// additional filtering. The prefix must minimally contain a Group Key (a
	// namespace and an item ID). BeginList will return an empty result set if
	// there are no items matching that key prefix. A token is returned from this
	// API that you can then pass to ContinueList to expand the result set, or to
	// SyncList to get updates within the result set. This can fail if the caller
	// does not have permission to read Items.
	BeginList(ctx context.Context, keyPath string, opts ...ListOptions) (ListResponse[Item], error)

	// ContinueList takes the token from a BeginList call and returns the next
	// "page" of results based on the original query parameters and pagination
	// options. It will return a new token which can be used for another
	// ContinueList call, and so on. The token is the same one used by SyncList -
	// each time you call either ContinueList or SyncList, you should pass the
	// latest version of the token, and then use the new token from the result in
	// subsequent calls. You may interleave ContinueList and SyncList calls
	// however you like, but it does not make sense to make both calls in
	// parallel. Calls to ContinueList are tied to the authorization of the
	// original BeginList call, so if the original BeginList call was allowed,
	// ContinueList with its token should also be allowed.
	ContinueList(ctx context.Context, token []byte) (ListResponse[Item], error)

	// NewTransaction starts a new transaction on a stream, and calls the
	// handler with a Transaction object. The handler can then interact with
	// the transaction by calling Get, Put, Append, Delete, and List. The
	// transaction is committed when the handler returns, and the results
	// are returned.
	NewTransaction(ctx context.Context, handler TransactionHandler) (*TransactionResults, error)

	// SyncList returns an iterator for a sync operation.
	SyncList(ctx context.Context, token []byte) (ListResponse[SyncResponse], error)
}

// TransactionResults holds all the results of a transaction after a commit.
type TransactionResults struct {
	// PutResponse contains the full result of each Put operation. This only
	// comes back with the transaction is finished message because full
	// metadata isn't available until then.
	PutResponse []Item

	// DeleteResponse contains the full result of each Delete operation. This
	// only comes back with the TransactionFinished message because full
	// metadata isn't available until then.
	DeleteResponse []string

	// Did the commit finish (the alternative is that it was aborted/rolled back)
	Committed bool
}

// Transaction represents a single transaction.
type Transaction interface {
	// Get one Item by its full key path.
	Get(item string) (Item, error)

	// GetBatch one or more Items by their full key paths.
	GetBatch(itemKeys ...string) ([]Item, error)

	// Put is a convenience method for adding a single Item to the Store, or
	// replacing the Item if it exists at that path. The returned GeneratedID
	// represents the ID the item will have upon a successful commit. This can
	// be referenced by other Items in the same commit or used to fill
	// our additional items in the same transaction safely.
	Put(item Item) (GeneratedID, error)

	// PutBatch adds one or more Items to the Store, or replaces the Items if
	// they exist at that path. The metadata (create time/version + modified
	// time/version)  for each PutItem is returned only at the end of the
	// transaction. The returned GeneratedID represents the ID the item will
	// have upon a successful commit. This can be referenced by other Items in
	// the same commit or used to fill our additional items in the same
	// transaction safely.
	PutBatch(items ...Item) ([]GeneratedID, error)

	// Delete removes one or more Items from the Store by their full key paths.
	Delete(itemKeys ...string) error

	// BeginList loads Items that start with a specified key path, subject to
	// additional filtering. The returned ListResponse can be used to iterate
	// over the stream of results for example:
	//	iter, err := txn.BeginList("/path/to/items")
	//	// handle err
	//	for iter.Next() {
	//		item := iter.Value()
	//		// do something with item
	//	}
	//	token, err := iter.Token();
	//	// handle err and token
	BeginList(prefix string, options ...ListOptions) (ListResponse[Item], error)

	// ContinueList picks back up where this token left off. As with BeginList,
	// you use the ListResponse to iterate over the stream of results.
	ContinueList(token *ListToken) (ListResponse[Item], error)
}

// TransactionHandler operates on a single transaction.
//
// The Transaction argument is passed to the handler function to allow the
// handler to interact with the transaction. This handler is not thread safe
// and should not be shared between goroutines. Additionally, do not share
// state outside the transaction handler. e.g. don't use a closure that
// captures variables from the outer scope.
//
// If you wish to cancel/abort the transaction, simply return an error from the
// handler and we'll take care of cleaning up the transaction.
type TransactionHandler func(Transaction) error

// NewClient creates a new client with the given store and options.
func NewClient(
	appCtx context.Context,
	storeID StoreID,
	itemTypeMapper ItemTypeMapper,
	options ...*Options,
) (Client, error) {
	if itemTypeMapper == nil {
		return nil, errors.New("itemTypeMapper is required when creating a client")
	}
	opts := &Options{}
	for _, o := range options {
		opts = opts.Merge(o)
	}
	opts, err := opts.ApplyDefaults(appCtx)
	if err != nil {
		return nil, err
	}

	return &client{
		client: dbconnect.NewDatabaseServiceClient(
			opts.HTTPClient(),
			opts.Endpoint,
			connect.WithCodec(grpc.Codec{}), // enable vtprotobuf codec
			connect.WithInterceptors(sdkerror.ConnectErrorInterceptor()),
		),
		storeID:    storeID,
		itemMapper: itemTypeMapper,
	}, nil
}

func (c *client) WithAllowStale(allowStale bool) Client {
	newClient := *c
	newClient.allowStale = allowStale
	return &newClient
}
