// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: db/service.proto

package dbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	db "github.com/StatelyCloud/go-sdk/pb/db"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// DatabaseServiceName is the fully-qualified name of the DatabaseService service.
	DatabaseServiceName = "stately.db.DatabaseService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DatabaseServicePutProcedure is the fully-qualified name of the DatabaseService's Put RPC.
	DatabaseServicePutProcedure = "/stately.db.DatabaseService/Put"
	// DatabaseServiceGetProcedure is the fully-qualified name of the DatabaseService's Get RPC.
	DatabaseServiceGetProcedure = "/stately.db.DatabaseService/Get"
	// DatabaseServiceDeleteProcedure is the fully-qualified name of the DatabaseService's Delete RPC.
	DatabaseServiceDeleteProcedure = "/stately.db.DatabaseService/Delete"
	// DatabaseServiceBeginListProcedure is the fully-qualified name of the DatabaseService's BeginList
	// RPC.
	DatabaseServiceBeginListProcedure = "/stately.db.DatabaseService/BeginList"
	// DatabaseServiceContinueListProcedure is the fully-qualified name of the DatabaseService's
	// ContinueList RPC.
	DatabaseServiceContinueListProcedure = "/stately.db.DatabaseService/ContinueList"
	// DatabaseServiceBeginScanProcedure is the fully-qualified name of the DatabaseService's BeginScan
	// RPC.
	DatabaseServiceBeginScanProcedure = "/stately.db.DatabaseService/BeginScan"
	// DatabaseServiceContinueScanProcedure is the fully-qualified name of the DatabaseService's
	// ContinueScan RPC.
	DatabaseServiceContinueScanProcedure = "/stately.db.DatabaseService/ContinueScan"
	// DatabaseServiceSyncListProcedure is the fully-qualified name of the DatabaseService's SyncList
	// RPC.
	DatabaseServiceSyncListProcedure = "/stately.db.DatabaseService/SyncList"
	// DatabaseServiceTransactionProcedure is the fully-qualified name of the DatabaseService's
	// Transaction RPC.
	DatabaseServiceTransactionProcedure = "/stately.db.DatabaseService/Transaction"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	databaseServiceServiceDescriptor            = db.File_db_service_proto.Services().ByName("DatabaseService")
	databaseServicePutMethodDescriptor          = databaseServiceServiceDescriptor.Methods().ByName("Put")
	databaseServiceGetMethodDescriptor          = databaseServiceServiceDescriptor.Methods().ByName("Get")
	databaseServiceDeleteMethodDescriptor       = databaseServiceServiceDescriptor.Methods().ByName("Delete")
	databaseServiceBeginListMethodDescriptor    = databaseServiceServiceDescriptor.Methods().ByName("BeginList")
	databaseServiceContinueListMethodDescriptor = databaseServiceServiceDescriptor.Methods().ByName("ContinueList")
	databaseServiceBeginScanMethodDescriptor    = databaseServiceServiceDescriptor.Methods().ByName("BeginScan")
	databaseServiceContinueScanMethodDescriptor = databaseServiceServiceDescriptor.Methods().ByName("ContinueScan")
	databaseServiceSyncListMethodDescriptor     = databaseServiceServiceDescriptor.Methods().ByName("SyncList")
	databaseServiceTransactionMethodDescriptor  = databaseServiceServiceDescriptor.Methods().ByName("Transaction")
)

// DatabaseServiceClient is a client for the stately.db.DatabaseService service.
type DatabaseServiceClient interface {
	// Put adds one or more Items to the Store, or replaces the Items if they
	// already exist. This will fail if the caller does not have permission to
	// create or update Items, if there is no schema registered for the provided
	// item type, or if an item is invalid. All puts are applied atomically;
	// either all will fail or all will succeed. If an item's schema specifies an
	// `initialValue` for one or more properties used in its key paths, and the
	// item is new, you should not provide those values - the database will choose
	// them for you, and Data must be provided as either serialized binary
	// protobuf or JSON.
	Put(context.Context, *connect.Request[db.PutRequest]) (*connect.Response[db.PutResponse], error)
	// Get retrieves one or more Items by their key paths. This will return any of
	// the Items that exist. It will fail if the caller does not have permission
	// to read Items. Use the List APIs if you want to retrieve multiple items but
	// don't already know the full key paths of the items you want to get.
	Get(context.Context, *connect.Request[db.GetRequest]) (*connect.Response[db.GetResponse], error)
	// Delete removes one or more Items from the Store by their key paths. This
	// will fail if the caller does not have permission to delete Items.
	// Tombstones will be saved for deleted items for some time, so
	// that SyncList can return information about deleted items. Deletes are
	// always applied atomically; all will fail or all will succeed.
	Delete(context.Context, *connect.Request[db.DeleteRequest]) (*connect.Response[db.DeleteResponse], error)
	// BeginList retrieves Items that start with a specified key path prefix. The
	// key path prefix must minimally contain a Group Key (a single key segment
	// with a namespace and an ID). BeginList will return an empty result set if
	// there are no items matching that key prefix. This API returns a token that
	// you can pass to ContinueList to expand the result set, or to SyncList to
	// get updates within the result set. This can fail if the caller does not
	// have permission to read Items.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	BeginList(context.Context, *connect.Request[db.BeginListRequest]) (*connect.ServerStreamForClient[db.ListResponse], error)
	// ContinueList takes the token from a BeginList call and returns more results
	// based on the original query parameters and pagination options. It has very
	// few options of its own because it is a continuation of a previous list
	// operation. It will return a new token which can be used for another
	// ContinueList call, and so on. The token is the same one used by SyncList -
	// each time you call either ContinueList or SyncList, you should pass the
	// latest version of the token, and then use the new token from the result in
	// subsequent calls. You may interleave ContinueList and SyncList calls
	// however you like, but it does not make sense to make both calls in
	// parallel. Calls to ContinueList are tied to the authorization of the
	// original BeginList call, so if the original BeginList call was allowed,
	// ContinueList with its token should also be allowed.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	ContinueList(context.Context, *connect.Request[db.ContinueListRequest]) (*connect.ServerStreamForClient[db.ListResponse], error)
	// BeginScan initiates a scan request which will scan over the entire store
	// and apply the provided filters. This API returns a token that you can pass
	// to ContinueScan to paginate through the result set. This can fail if the
	// caller does not have permission to read Items.
	// WARNING: THIS API CAN BE EXTREMELY EXPENSIVE FOR STORES WITH A LARGE NUMBER
	// OF ITEMS.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	BeginScan(context.Context, *connect.Request[db.BeginScanRequest]) (*connect.ServerStreamForClient[db.ListResponse], error)
	// ContinueScan takes the token from a BeginScan call and returns more results
	// based on the original request parameters and pagination options. It has
	// very few options of its own because it is a continuation of a previous list
	// operation. It will return a new token which can be used for another
	// ContinueScan call, and so on. Calls to ContinueScan are tied to the
	// authorization of the original BeginScan call, so if the original BeginScan
	// call was allowed, ContinueScan with its token should also be allowed.
	// WARNING: THIS API CAN BE EXTREMELY EXPENSIVE FOR STORES WITH A LARGE NUMBER OF ITEMS.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	ContinueScan(context.Context, *connect.Request[db.ContinueScanRequest]) (*connect.ServerStreamForClient[db.ListResponse], error)
	// SyncList returns all changes to Items within the result set of a previous
	// List operation. For all Items within the result set that were modified, it
	// returns the full Item at in its current state. It also returns a list of
	// Item key paths that were deleted since the last SyncList, which you should
	// reconcile with your view of items returned from previous
	// BeginList/ContinueList calls. Using this API, you can start with an initial
	// set of items from BeginList, and then stay up to date on any changes via
	// repeated SyncList requests over time. The token is the same one used by
	// ContinueList - each time you call either ContinueList or SyncList, you
	// should pass the latest version of the token, and then use the new token
	// from the result in subsequent calls. Note that if the result set has
	// already been expanded to the end (in the direction of the original
	// BeginList request), SyncList will return newly created Items. You may
	// interleave ContinueList and SyncList calls however you like, but it does
	// not make sense to make both calls in parallel. Calls to SyncList are tied
	// to the authorization of the original BeginList call, so if the original
	// BeginList call was allowed, SyncList with its token should also be allowed.
	SyncList(context.Context, *connect.Request[db.SyncListRequest]) (*connect.ServerStreamForClient[db.SyncListResponse], error)
	// Transaction performs a transaction, within which you can issue writes
	// (Put/Delete) and reads (Get/List) in any order, followed by a commit
	// message. Reads are guaranteed to reflect the state as of when the
	// transaction started, and writes are committed atomically. This method may
	// fail if another transaction commits before this one finishes - in that
	// case, you should retry your transaction.
	Transaction(context.Context) *connect.BidiStreamForClient[db.TransactionRequest, db.TransactionResponse]
}

// NewDatabaseServiceClient constructs a client for the stately.db.DatabaseService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDatabaseServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DatabaseServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &databaseServiceClient{
		put: connect.NewClient[db.PutRequest, db.PutResponse](
			httpClient,
			baseURL+DatabaseServicePutProcedure,
			connect.WithSchema(databaseServicePutMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[db.GetRequest, db.GetResponse](
			httpClient,
			baseURL+DatabaseServiceGetProcedure,
			connect.WithSchema(databaseServiceGetMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[db.DeleteRequest, db.DeleteResponse](
			httpClient,
			baseURL+DatabaseServiceDeleteProcedure,
			connect.WithSchema(databaseServiceDeleteMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		beginList: connect.NewClient[db.BeginListRequest, db.ListResponse](
			httpClient,
			baseURL+DatabaseServiceBeginListProcedure,
			connect.WithSchema(databaseServiceBeginListMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		continueList: connect.NewClient[db.ContinueListRequest, db.ListResponse](
			httpClient,
			baseURL+DatabaseServiceContinueListProcedure,
			connect.WithSchema(databaseServiceContinueListMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		beginScan: connect.NewClient[db.BeginScanRequest, db.ListResponse](
			httpClient,
			baseURL+DatabaseServiceBeginScanProcedure,
			connect.WithSchema(databaseServiceBeginScanMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		continueScan: connect.NewClient[db.ContinueScanRequest, db.ListResponse](
			httpClient,
			baseURL+DatabaseServiceContinueScanProcedure,
			connect.WithSchema(databaseServiceContinueScanMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		syncList: connect.NewClient[db.SyncListRequest, db.SyncListResponse](
			httpClient,
			baseURL+DatabaseServiceSyncListProcedure,
			connect.WithSchema(databaseServiceSyncListMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		transaction: connect.NewClient[db.TransactionRequest, db.TransactionResponse](
			httpClient,
			baseURL+DatabaseServiceTransactionProcedure,
			connect.WithSchema(databaseServiceTransactionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// databaseServiceClient implements DatabaseServiceClient.
type databaseServiceClient struct {
	put          *connect.Client[db.PutRequest, db.PutResponse]
	get          *connect.Client[db.GetRequest, db.GetResponse]
	delete       *connect.Client[db.DeleteRequest, db.DeleteResponse]
	beginList    *connect.Client[db.BeginListRequest, db.ListResponse]
	continueList *connect.Client[db.ContinueListRequest, db.ListResponse]
	beginScan    *connect.Client[db.BeginScanRequest, db.ListResponse]
	continueScan *connect.Client[db.ContinueScanRequest, db.ListResponse]
	syncList     *connect.Client[db.SyncListRequest, db.SyncListResponse]
	transaction  *connect.Client[db.TransactionRequest, db.TransactionResponse]
}

// Put calls stately.db.DatabaseService.Put.
func (c *databaseServiceClient) Put(ctx context.Context, req *connect.Request[db.PutRequest]) (*connect.Response[db.PutResponse], error) {
	return c.put.CallUnary(ctx, req)
}

// Get calls stately.db.DatabaseService.Get.
func (c *databaseServiceClient) Get(ctx context.Context, req *connect.Request[db.GetRequest]) (*connect.Response[db.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Delete calls stately.db.DatabaseService.Delete.
func (c *databaseServiceClient) Delete(ctx context.Context, req *connect.Request[db.DeleteRequest]) (*connect.Response[db.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// BeginList calls stately.db.DatabaseService.BeginList.
func (c *databaseServiceClient) BeginList(ctx context.Context, req *connect.Request[db.BeginListRequest]) (*connect.ServerStreamForClient[db.ListResponse], error) {
	return c.beginList.CallServerStream(ctx, req)
}

// ContinueList calls stately.db.DatabaseService.ContinueList.
func (c *databaseServiceClient) ContinueList(ctx context.Context, req *connect.Request[db.ContinueListRequest]) (*connect.ServerStreamForClient[db.ListResponse], error) {
	return c.continueList.CallServerStream(ctx, req)
}

// BeginScan calls stately.db.DatabaseService.BeginScan.
func (c *databaseServiceClient) BeginScan(ctx context.Context, req *connect.Request[db.BeginScanRequest]) (*connect.ServerStreamForClient[db.ListResponse], error) {
	return c.beginScan.CallServerStream(ctx, req)
}

// ContinueScan calls stately.db.DatabaseService.ContinueScan.
func (c *databaseServiceClient) ContinueScan(ctx context.Context, req *connect.Request[db.ContinueScanRequest]) (*connect.ServerStreamForClient[db.ListResponse], error) {
	return c.continueScan.CallServerStream(ctx, req)
}

// SyncList calls stately.db.DatabaseService.SyncList.
func (c *databaseServiceClient) SyncList(ctx context.Context, req *connect.Request[db.SyncListRequest]) (*connect.ServerStreamForClient[db.SyncListResponse], error) {
	return c.syncList.CallServerStream(ctx, req)
}

// Transaction calls stately.db.DatabaseService.Transaction.
func (c *databaseServiceClient) Transaction(ctx context.Context) *connect.BidiStreamForClient[db.TransactionRequest, db.TransactionResponse] {
	return c.transaction.CallBidiStream(ctx)
}

// DatabaseServiceHandler is an implementation of the stately.db.DatabaseService service.
type DatabaseServiceHandler interface {
	// Put adds one or more Items to the Store, or replaces the Items if they
	// already exist. This will fail if the caller does not have permission to
	// create or update Items, if there is no schema registered for the provided
	// item type, or if an item is invalid. All puts are applied atomically;
	// either all will fail or all will succeed. If an item's schema specifies an
	// `initialValue` for one or more properties used in its key paths, and the
	// item is new, you should not provide those values - the database will choose
	// them for you, and Data must be provided as either serialized binary
	// protobuf or JSON.
	Put(context.Context, *connect.Request[db.PutRequest]) (*connect.Response[db.PutResponse], error)
	// Get retrieves one or more Items by their key paths. This will return any of
	// the Items that exist. It will fail if the caller does not have permission
	// to read Items. Use the List APIs if you want to retrieve multiple items but
	// don't already know the full key paths of the items you want to get.
	Get(context.Context, *connect.Request[db.GetRequest]) (*connect.Response[db.GetResponse], error)
	// Delete removes one or more Items from the Store by their key paths. This
	// will fail if the caller does not have permission to delete Items.
	// Tombstones will be saved for deleted items for some time, so
	// that SyncList can return information about deleted items. Deletes are
	// always applied atomically; all will fail or all will succeed.
	Delete(context.Context, *connect.Request[db.DeleteRequest]) (*connect.Response[db.DeleteResponse], error)
	// BeginList retrieves Items that start with a specified key path prefix. The
	// key path prefix must minimally contain a Group Key (a single key segment
	// with a namespace and an ID). BeginList will return an empty result set if
	// there are no items matching that key prefix. This API returns a token that
	// you can pass to ContinueList to expand the result set, or to SyncList to
	// get updates within the result set. This can fail if the caller does not
	// have permission to read Items.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	BeginList(context.Context, *connect.Request[db.BeginListRequest], *connect.ServerStream[db.ListResponse]) error
	// ContinueList takes the token from a BeginList call and returns more results
	// based on the original query parameters and pagination options. It has very
	// few options of its own because it is a continuation of a previous list
	// operation. It will return a new token which can be used for another
	// ContinueList call, and so on. The token is the same one used by SyncList -
	// each time you call either ContinueList or SyncList, you should pass the
	// latest version of the token, and then use the new token from the result in
	// subsequent calls. You may interleave ContinueList and SyncList calls
	// however you like, but it does not make sense to make both calls in
	// parallel. Calls to ContinueList are tied to the authorization of the
	// original BeginList call, so if the original BeginList call was allowed,
	// ContinueList with its token should also be allowed.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	ContinueList(context.Context, *connect.Request[db.ContinueListRequest], *connect.ServerStream[db.ListResponse]) error
	// BeginScan initiates a scan request which will scan over the entire store
	// and apply the provided filters. This API returns a token that you can pass
	// to ContinueScan to paginate through the result set. This can fail if the
	// caller does not have permission to read Items.
	// WARNING: THIS API CAN BE EXTREMELY EXPENSIVE FOR STORES WITH A LARGE NUMBER
	// OF ITEMS.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	BeginScan(context.Context, *connect.Request[db.BeginScanRequest], *connect.ServerStream[db.ListResponse]) error
	// ContinueScan takes the token from a BeginScan call and returns more results
	// based on the original request parameters and pagination options. It has
	// very few options of its own because it is a continuation of a previous list
	// operation. It will return a new token which can be used for another
	// ContinueScan call, and so on. Calls to ContinueScan are tied to the
	// authorization of the original BeginScan call, so if the original BeginScan
	// call was allowed, ContinueScan with its token should also be allowed.
	// WARNING: THIS API CAN BE EXTREMELY EXPENSIVE FOR STORES WITH A LARGE NUMBER OF ITEMS.
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	ContinueScan(context.Context, *connect.Request[db.ContinueScanRequest], *connect.ServerStream[db.ListResponse]) error
	// SyncList returns all changes to Items within the result set of a previous
	// List operation. For all Items within the result set that were modified, it
	// returns the full Item at in its current state. It also returns a list of
	// Item key paths that were deleted since the last SyncList, which you should
	// reconcile with your view of items returned from previous
	// BeginList/ContinueList calls. Using this API, you can start with an initial
	// set of items from BeginList, and then stay up to date on any changes via
	// repeated SyncList requests over time. The token is the same one used by
	// ContinueList - each time you call either ContinueList or SyncList, you
	// should pass the latest version of the token, and then use the new token
	// from the result in subsequent calls. Note that if the result set has
	// already been expanded to the end (in the direction of the original
	// BeginList request), SyncList will return newly created Items. You may
	// interleave ContinueList and SyncList calls however you like, but it does
	// not make sense to make both calls in parallel. Calls to SyncList are tied
	// to the authorization of the original BeginList call, so if the original
	// BeginList call was allowed, SyncList with its token should also be allowed.
	SyncList(context.Context, *connect.Request[db.SyncListRequest], *connect.ServerStream[db.SyncListResponse]) error
	// Transaction performs a transaction, within which you can issue writes
	// (Put/Delete) and reads (Get/List) in any order, followed by a commit
	// message. Reads are guaranteed to reflect the state as of when the
	// transaction started, and writes are committed atomically. This method may
	// fail if another transaction commits before this one finishes - in that
	// case, you should retry your transaction.
	Transaction(context.Context, *connect.BidiStream[db.TransactionRequest, db.TransactionResponse]) error
}

// NewDatabaseServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDatabaseServiceHandler(svc DatabaseServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	databaseServicePutHandler := connect.NewUnaryHandler(
		DatabaseServicePutProcedure,
		svc.Put,
		connect.WithSchema(databaseServicePutMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceGetHandler := connect.NewUnaryHandler(
		DatabaseServiceGetProcedure,
		svc.Get,
		connect.WithSchema(databaseServiceGetMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceDeleteHandler := connect.NewUnaryHandler(
		DatabaseServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(databaseServiceDeleteMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceBeginListHandler := connect.NewServerStreamHandler(
		DatabaseServiceBeginListProcedure,
		svc.BeginList,
		connect.WithSchema(databaseServiceBeginListMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceContinueListHandler := connect.NewServerStreamHandler(
		DatabaseServiceContinueListProcedure,
		svc.ContinueList,
		connect.WithSchema(databaseServiceContinueListMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceBeginScanHandler := connect.NewServerStreamHandler(
		DatabaseServiceBeginScanProcedure,
		svc.BeginScan,
		connect.WithSchema(databaseServiceBeginScanMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceContinueScanHandler := connect.NewServerStreamHandler(
		DatabaseServiceContinueScanProcedure,
		svc.ContinueScan,
		connect.WithSchema(databaseServiceContinueScanMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceSyncListHandler := connect.NewServerStreamHandler(
		DatabaseServiceSyncListProcedure,
		svc.SyncList,
		connect.WithSchema(databaseServiceSyncListMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	databaseServiceTransactionHandler := connect.NewBidiStreamHandler(
		DatabaseServiceTransactionProcedure,
		svc.Transaction,
		connect.WithSchema(databaseServiceTransactionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/stately.db.DatabaseService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DatabaseServicePutProcedure:
			databaseServicePutHandler.ServeHTTP(w, r)
		case DatabaseServiceGetProcedure:
			databaseServiceGetHandler.ServeHTTP(w, r)
		case DatabaseServiceDeleteProcedure:
			databaseServiceDeleteHandler.ServeHTTP(w, r)
		case DatabaseServiceBeginListProcedure:
			databaseServiceBeginListHandler.ServeHTTP(w, r)
		case DatabaseServiceContinueListProcedure:
			databaseServiceContinueListHandler.ServeHTTP(w, r)
		case DatabaseServiceBeginScanProcedure:
			databaseServiceBeginScanHandler.ServeHTTP(w, r)
		case DatabaseServiceContinueScanProcedure:
			databaseServiceContinueScanHandler.ServeHTTP(w, r)
		case DatabaseServiceSyncListProcedure:
			databaseServiceSyncListHandler.ServeHTTP(w, r)
		case DatabaseServiceTransactionProcedure:
			databaseServiceTransactionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDatabaseServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDatabaseServiceHandler struct{}

func (UnimplementedDatabaseServiceHandler) Put(context.Context, *connect.Request[db.PutRequest]) (*connect.Response[db.PutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.Put is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) Get(context.Context, *connect.Request[db.GetRequest]) (*connect.Response[db.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.Get is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) Delete(context.Context, *connect.Request[db.DeleteRequest]) (*connect.Response[db.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.Delete is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) BeginList(context.Context, *connect.Request[db.BeginListRequest], *connect.ServerStream[db.ListResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.BeginList is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) ContinueList(context.Context, *connect.Request[db.ContinueListRequest], *connect.ServerStream[db.ListResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.ContinueList is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) BeginScan(context.Context, *connect.Request[db.BeginScanRequest], *connect.ServerStream[db.ListResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.BeginScan is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) ContinueScan(context.Context, *connect.Request[db.ContinueScanRequest], *connect.ServerStream[db.ListResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.ContinueScan is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) SyncList(context.Context, *connect.Request[db.SyncListRequest], *connect.ServerStream[db.SyncListResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.SyncList is not implemented"))
}

func (UnimplementedDatabaseServiceHandler) Transaction(context.Context, *connect.BidiStream[db.TransactionRequest, db.TransactionResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("stately.db.DatabaseService.Transaction is not implemented"))
}
