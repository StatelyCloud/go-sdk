package stately

import (
	"context"

	"github.com/StatelyCloud/go-sdk/pb/db/dbconnect"
	"github.com/StatelyCloud/go-sdk/sconnect"
)

type client struct {
	client          dbconnect.DatabaseServiceClient
	storeID         StoreID
	schemaVersionID SchemaVersionID
	schemaID        SchemaID
	itemMapper      ItemTypeMapper

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

	// GetBatch retrieves multiple items by their full key paths. This will return
	// any of the Items that exist. Use BeginList if you want to retrieve multiple
	// items but don't already know the full key paths of the items you want to
	// get. You can get items of different types in a single getBatch but you will
	// need to use a type switch to determine what item type each item is.
	//
	// Example:
	//  items, err := client.GetBatch(ctx, "/movies-123", "/movies-456")
	GetBatch(ctx context.Context, itemPaths ...string) ([]Item, error)

	// Get retrieves one Item by its full key path, or nil if no item exists at
	// that path.
	//
	// Example:
	//  item, err := client.Get(ctx, "/movies-123")
	Get(ctx context.Context, itemPath string) (Item, error)

	// PutBatch adds multiple Items to the Store, or replaces Items if they
	// already exist at that path. Each item may optionally be wrapped in an
	// WithPutOptions to specify additional per-item options. All puts in the
	// request are applied atomically - there are no partial successes.
	//
	// This will fail if:
	//   - Any Item conflicts with an existing Item at the same path and its
	//     MustNotExist option is set, or the item's ID will be chosen with an
	//     `initialValue` and one of its other key paths conflicts with an existing
	//     item.
	//
	// Additional Notes: Example:
	//  items, err := client.PutBatch(ctx, item, item2)
	//  items, err := client.PutBatch(ctx, item, stately.WithPutOptions{Item: item2, MustNotExist:true})
	PutBatch(ctx context.Context, items ...Item) ([]Item, error)

	// Put adds one Item to the Store, or replaces the Item if it already exists
	// at that path. The item may optionally be wrapped in an WithPutOptions
	// to specify additional per-item options.
	//
	// This call will fail if:
	//   - The Item conflicts with an existing Item at the same path and the
	//     MustNotExist option is set, or the item's ID will be chosen with an
	//     `initialValue` and one of its other key paths conflicts with an existing
	//     item.
	//
	// Example:
	//  item, err := client.Put(ctx, item)
	//  item, err := client.Put(ctx, stately.WithPutOptions{Item: item, MustNotExist:true})
	Put(ctx context.Context, item Item) (Item, error)

	// Delete removes one or more items from the Store by their full key paths.
	// Delete succeeds even if there isn't an item at that key path. Tombstones
	// will be saved for deleted items for some time, so that SyncList can return
	// information about deleted items. Deletes are always applied atomically; all
	// will fail or all will succeed.
	//
	// Example:
	//  err := client.Delete(ctx, "/movies-123", "/movies-456")
	Delete(ctx context.Context, itemPaths ...string) error

	// BeginList retrieves Items that start with a specified keyPathPrefix from a
	// single Group. Because it can only list items from a single Group, the key
	// path prefix must at least start with a full Group Key (a single key segment
	// with a namespace and an ID, e.g. `/user-1234`).
	//
	// BeginList will return an empty result set if there are no items matching
	// that key prefix. This API returns a token that you can pass to ContinueList
	// to expand the result set, or to SyncList to get updates within the result
	// set.
	//
	// The options parameter is optional and can be used to limit the number of
	// results in a page, change the sort order, etc. If you provide multiple
	// options objects, the last one will take precedence.
	//
	// You can list items of different types in a single BeginList, and you can
	// use a type switch to handle different item types.
	//
	// Example:
	//   iter, err := client.BeginList(ctx, "/movies-movieID", stately.ListOptions{Limit: 10})
	//   if err != nil { return err }
	//   for iter.Next() {
	//     item := iter.Value()
	//     // do something with item
	//   }
	//   token, err := iter.Token() // Save this for ContinueList/SyncList
	BeginList(ctx context.Context, keyPath string, opts ...ListOptions) (ListResponse[Item], error)

	// ContinueList takes the token from a BeginList call and returns the next
	// "page" of results based on the original query parameters and pagination
	// options. It doesn't have options because it is a continuation of a previous
	// list operation. It will return a new token which can be used for another
	// ContinueList call, and so on. The token is the same one used by SyncList -
	// each time you call either ContinueList or SyncList, you should pass the
	// latest version of the token, and then use the new token from the result in
	// subsequent calls. You may interleave ContinueList and SyncList calls
	// however you like, but it does not make sense to make both calls in
	// parallel. Calls to ContinueList are tied to the authorization of the
	// original BeginList call, so if the original BeginList call was allowed,
	// ContinueList with its token should also be allowed.
	//
	// You can list items of different types in a single ContinueList, and you can
	// use a type switch to handle different item types.
	//
	// Example:
	//   iter, err := client.ContinueList(ctx, token.Data)
	//   if err != nil { return err }
	//   for iter.Next() {
	//     item := iter.Value()
	//     // do something with item
	//   }
	//   token, err := iter.Token() // Save this for ContinueList/SyncList
	ContinueList(ctx context.Context, token []byte) (ListResponse[Item], error)

	// BeginScan retrieves all Items from the Store. This API returns a token that
	// you can pass to ContinueScan to expand the result set. This can fail if the
	// caller does not have permission to read Items.
	//
	// The options parameter is optional and can be used to limit the number of
	// results in a page, return only items of a certain type, or implement
	// parallel segmented scans. If you provide multiple options objects, the last
	// one will take precedence.
	//
	// You can list items of different types in a single BeginScan, and you can
	// use a type switch to handle different item types.
	//
	// WARNING: THIS API CAN BE EXPENSIVE FOR STORES WITH A LARGE NUMBER OF ITEMS.
	//
	// Example:
	//   iter, err := client.BeginScan(ctx, stately.ScanOptions{Limit: 10, ItemTypes: []string{"Movie"}})
	//   if err != nil { return err }
	//   for iter.Next() {
	//     item := iter.Value()
	//     // do something with item
	//   }
	//   token, err := iter.Token() // Save this for ContinueScan
	BeginScan(ctx context.Context, opts ...ScanOptions) (ListResponse[Item], error)

	// ContinueScan takes the token from a BeginScan call and returns the next
	// "page" of results based on the original scan parameters and pagination
	// options. It will return a new token which can be used for another
	// ContinueScan call, and so on. Each time you call ContinueScan, you should
	// pass the latest version of the token, and then use the new token from the
	// result in subsequent calls. Calls to ContinueScan are tied to the
	// authorization of the original BeginScan call, so if the original BeginScan
	// call was allowed, ContinueScan with its token should also be allowed.
	//
	// You can list items of different types in a single ContinueScan, and you can
	// use a type switch to handle different item types.
	//
	// WARNING: THIS API CAN BE EXPENSIVE FOR STORES WITH A LARGE NUMBER OF ITEMS.
	//
	// Example:
	//   iter, err := client.ContinueScan(ctx, token.Data)
	//   if err != nil { return err }
	//   for iter.Next() {
	//     item := iter.Value()
	//     // do something with item
	//   }
	//   token, err := iter.Token() // Save this for ContinueScan
	ContinueScan(ctx context.Context, token []byte) (ListResponse[Item], error)

	// NewTransaction allows you to issue reads and writes in any order, and all
	// writes will either succeed or all will fail when the transaction finishes.
	// You pass it a function with a single parameter, the transaction handler,
	// which lets you perform operations within the transaction.
	//
	// Reads are guaranteed to reflect the state as of when the transaction
	// started. A transaction may fail if another transaction commits before this
	// one finishes - in that case, you should retry your transaction.
	//
	// If any error is returned from the handler function, the transaction is
	// aborted and none of the changes made in it will be applied. If the handler
	// returns without error, the transaction is automatically committed.
	//
	// If any of the operations in the handler function fails (e.g. a request is
	// invalid) you may not find out until the *next* operation, or once the block
	// finishes, due to some technicalities about how requests are handled.
	//
	// When the transaction is committed, the result property will contain the
	// full version of any items that were put in the transaction, and the
	// committed property will be True. If the transaction was aborted, the
	// committed property will be False.
	//
	// Example:
	//  results, err := client.NewTransaction(ctx, func(txn stately.Transaction) error {
	//    item, err := txn.Get(ctx, "/movies-123")
	//    if err != nil { return err }
	//    item.Title = "New Title"
	//    _, err = txn.Put(ctx, item)
	//    return err
	//  })
	//  if err != nil { return err }
	//  for _, item := range results.PutResponse {
	//    // do something with the updated item
	//  }
	NewTransaction(ctx context.Context, handler TransactionHandler) (*TransactionResults, error)

	// SyncList returns all changes to Items within the result set of a previous
	// List operation. For all Items within the result set that were modified, it
	// returns the full Item at in its current state. If the result set has
	// already been expanded to the end (in the direction of the original
	// BeginList request), SyncList will return newly created Items as well. It
	// also returns a list of Item key paths that were deleted since the last
	// SyncList, which you should reconcile with your view of items returned from
	// previous BeginList/ContinueList calls. Using this API, you can start with
	// an initial set of items from beginList, and then stay up to date on any
	// changes via repeated SyncList requests over time.
	//
	// The token is the same one used by ContinueList - each time you call either
	// ContinueList or SyncList, you should pass the latest version of the token,
	// and then use the new token from the result in subsequent calls. You may
	// interleave ContinueList and SyncList calls however you like, but it does
	// not make sense to make both calls in parallel. Calls to SyncList are tied
	// to the authorization of the original BeginList call, so if the original
	// beginList call was allowed, SyncList with its token should also be allowed.
	//
	// Each result will be one of the following types:
	//     - *stately.Changed: An item that was changed or added since the last
	//       SyncList call.
	//     - *stately.Deleted: The key path of an item that was deleted since
	//       the last SyncList call.
	//     - *stately.UpdateOutsideOfWindow: An item that was updated but
	//       is not within the current result set. You can treat this like
	//       stately.Deleted, but the item hasn't actually been deleted, it's
	//       just not part of your view of the list anymore.
	//     - *stately.Reset: A reset signal that indicates any previously cached
	//       view of the result set is no longer valid. You should throw away
	//       any locally cached data. This will always be followed by a series
	//       of *stately.Changed messages that make up a new view of the result set.
	//
	// Example:
	//
	//  iter, err := client.SyncList(ctx, token.Data)
	//  for iter.Next() {
	//    switch v := iter.Value().(type) {
	//    case *stately.Changed:
	//      // do something with the changed item: v.Item
	//    case *stately.Deleted:
	//      // do something with removed key path: v.KeyPath
	//    case *stately.UpdateOutsideOfWindow:
	//      // do something with the out of window update: v.KeyPath
	//    case *stately.Reset:
	//      // reset the sync operation
	//    }
	//  }
	//  err, token := iter.Token() // Save this for ContinueList/SyncList
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
	// Get retrieves one Item by its full key path, or nil if no item exists at
	// that path. Use BeginList if you want to retrieve multiple items but don't
	// already know the full key paths of the items you want to get. Use GetBatch
	// if you want to retrieve multiple items by their full key paths.
	//
	// Example:
	//  item, err := txn.Get("/movies-123")
	Get(item string) (Item, error)

	// GetBatch retrieves multiple items by their full key paths. This will return
	// any of the Items that exist. Use BeginList if you want to retrieve multiple
	// items but don't already know the full key paths of the items you want to
	// get. Use Get if you want to retrieve a single item. You can get items of
	// different types in a single GetBatch - you will need to use a type switch
	// to determine what item type each item is.
	//
	// Example:
	//  items, err := client.GetBatch("/movies-123", "/movies-456")
	GetBatch(itemKeys ...string) ([]Item, error)

	// Put adds one Item to the Store, or replaces the Item if it already exists
	// at that path. The item may optionally be wrapped in an WithPutOptions
	// to specify additional per-item options.

	// The metadata (create time/version + modified time/version) for each put
	// item is returned only at the end of the transaction. The returned
	// GeneratedID represents the ID the item *will* have upon a successful
	// commit. You can use this ID to build other Items in the same commit.
	//
	// This call will cause the transaction to fail if:
	//   - The Item conflicts with an existing Item at the same path and the
	//     MustNotExist option is set, or the item's ID will be chosen with an
	//     `initialValue` and one of its other key paths conflicts with an existing
	//     item.
	//
	// Example:
	//  genID, err := txn.Put(item)
	//  genID, err := txn.Put(stately.WithPutOptions{Item: item, MustNotExist:true})
	Put(item Item) (GeneratedID, error)

	// PutBatch adds multiple Items to the Store, or replaces the Items if they
	// already exist at that path. Each item may optionally be wrapped in an
	// WithPutOptions to specify additional per-item options.
	//
	// The metadata (create time/version + modified time/version) for each put
	// item is returned only at the end of the transaction. The returned
	// GeneratedID represents the ID the item *will* have upon a successful
	// commit. You can use this ID to build other Items in the same commit.
	//
	// This will cause the transaction to fail if:
	//   - Any Item conflicts with an existing Item at the same path and its
	//     MustNotExist option is set, or the item's ID will be chosen with an
	//     `initialValue` and one of its other key paths conflicts with an existing
	//     item.
	//
	// Example:
	//  genIDs, err := txn.PutBatch(item, item2)
	//  genIDs, err := txn.PutBatch(item, stately.WithPutOptions{Item: item2, MustNotExist:true})
	PutBatch(items ...Item) ([]GeneratedID, error)

	// Delete removes multiple Items from the Store by their key paths. Delete
	// succeeds even if there isn't an item at that key path.
	//
	// Example:
	//  err := txn.Delete("/movies-123", "/movies-456")
	Delete(itemKeys ...string) error

	// BeginList retrieves Items that start with a specified keyPathPrefix from a
	// single Group. Because it can only list items from a single Group, the key
	// path prefix must at least start with a full Group Key (a single key segment
	// with a namespace and an ID, e.g. `/user-1234`).
	//
	// BeginList will return an empty result set if there are no items matching
	// that key prefix. This API returns a token that you can pass to ContinueList
	// to expand the result set.
	//
	// The options parameter is optional and can be used to limit the number of
	// results in a page, change the sort order, etc. If you provide multiple
	// options objects, the last one will take precedence.
	//
	// Example:
	//   iter, err := txn.BeginList("/movies", stately.ListOptions{Limit: 10})
	//   if err != nil { return err }
	//   for iter.Next() {
	//     item := iter.Value()
	//     // do something with item
	//   }
	//   token, err := iter.Token() // Save this for ContinueList/SyncList
	BeginList(prefix string, options ...ListOptions) (ListResponse[Item], error)

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
	//
	// Example:
	//   iter, err := txn.ContinueList(token.Data)
	//   if err != nil { return err }
	//   for iter.Next() {
	//     item := iter.Value()
	//     // do something with item
	//   }
	//   token, err := iter.Token() // Save this for ContinueList/SyncList
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

// NewClient creates a new client with the given store + schema version with
// options.
// Deprecated: This function is deprecated and will be removed in a future release.
func NewClient(
	appCtx context.Context,
	storeID uint64,
	schemaVersionID SchemaVersionID,
	schemaID uint64,
	itemTypeMapper ItemTypeMapper,
	options ...*Options,
) Client {
	return InternalBindClientToTypes(
		appCtx,
		storeID,
		itemTypeMapper,
		InternalClientOptions{
			SchemaID:        SchemaID(schemaID),
			SchemaVersionID: schemaVersionID,
		},
		options...,
	)
}

// InternalClientOptions are options that will be passed from generated code to
// the internal client constructor InternalBindClientToTypes. This is a struct
// to make it so new fields can be added later without breaking backwards
// compatibility every time, the way adding parameters to a function would.
type InternalClientOptions struct {
	// SchemaID is the schema ID that this client was generated for. All its types
	// are specific to a particular version of this schema.
	SchemaID

	// SchemaVersionID is the schema version that this client was generated for. All its types are specific to this version.
	SchemaVersionID
}

// InternalBindClientToTypes should not be called directly by users - instead,
// you should use the NewClient function in your generated schema code, which
// passes the correct type information to this function.
func InternalBindClientToTypes(
	appCtx context.Context,
	storeID uint64,
	itemTypeMapper ItemTypeMapper,
	internalOptions InternalClientOptions,
	options ...*Options,
) Client {
	if itemTypeMapper == nil {
		panic("ItemTypeMapper is required when creating a client")
	}
	if internalOptions.SchemaVersionID == 0 {
		panic("SchemaVersionID is required when creating a client")
	}
	opts := &Options{}
	for _, o := range options {
		opts = opts.Merge(o)
	}
	opts = opts.ApplyDefaults(appCtx)

	clientOpts := sconnect.ConnectClientOptions
	if isLocalEndpoint(opts.Endpoint) {
		clientOpts = sconnect.LocalConnectClientOptions
	}

	return &client{
		client: dbconnect.NewDatabaseServiceClient(
			opts.HTTPClient(),
			opts.Endpoint,
			clientOpts...,
		),
		storeID:         StoreID(storeID),
		itemMapper:      itemTypeMapper,
		schemaVersionID: internalOptions.SchemaVersionID,
		schemaID:        internalOptions.SchemaID,
	}
}

// WithAllowStale returns a new client with the given allowStale value set. Any
// read APIs called with this copy of the client will allow
// eventually-consistent (i.e. possibly stale) reads, which cost less. By
// default, clients do not allow stale reads.
func (c *client) WithAllowStale(allowStale bool) Client {
	newClient := *c
	newClient.allowStale = allowStale
	return &newClient
}
