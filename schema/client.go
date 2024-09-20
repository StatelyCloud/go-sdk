package schema

import (
	"context"

	"connectrpc.com/connect"
	"github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
	"github.com/StatelyCloud/go-sdk/pb/schemaservice/schemaserviceconnect"
	"github.com/StatelyCloud/go-sdk/sdkerror"
	"github.com/StatelyCloud/go-sdk/stately"
)

type schemaClient struct {
	client  schemaserviceconnect.SchemaServiceClient
	storeID stately.StoreID
}

// Client is a stately schema client that interacts with the given store.
type Client interface {
	// Get retrieves the fully self-contained Schema for the corresponding Store
	// ID. There is only one Schema per Store so the result of this call will
	// contain the most up-to-date representation of the Items in the Store. It
	// will fail if the caller does not have permission the Store.
	// If the store does not have a schema yet, this will return nil, nil.
	Get(ctx context.Context) (*schemaservice.SchemaModel, error)

	// Put adds a Schema to the StatelyDB Schema Store or replaces the Schema if
	// it already exists. If the caller attempts to put a Schema for a Store that
	// does not exist the request will fail. If the caller does not have
	// permissions to access the Store the request will fail. If the Schema is not
	// valid the request will fail. If a Schema already exists for the Store then
	// the update will only be accepted if the new Schema is backwards compatible
	// with the existing Schema.
	Put(
		ctx context.Context,
		fileDescriptor *descriptorpb.FileDescriptorProto,
		changeDescription string,
		options ...*PutOptions,
	) (bool, *schemaservice.ValidateResponse, error)

	Validate(
		ctx context.Context,
		fileDescriptor *descriptorpb.FileDescriptorProto,
	) (*schemaservice.ValidateResponse, error)

	// ListAuditLog retrieves the audit log for the Schema associated with the provided storeId.
	// The audit log consists of a list of audit log entries that represent each change to the Schema including
	// its creation. The list is ordered by the time of the change with the most recent change first.
	// If there is no Schema for the provided Store ID, an empty list will be returned.
	ListAuditLog(ctx context.Context, options ...*ListAuditLogOptions) ([]*schemaservice.SchemaAuditLogEntry, error)
}

// NewClient creates a new client with the given store and options.
func NewClient(appCtx context.Context, storeID uint64, options ...*stately.Options) (Client, error) {
	opts := &stately.Options{}
	for _, o := range options {
		opts = opts.Merge(o)
	}
	opts, err := opts.ApplyDefaults(appCtx)
	if err != nil {
		return nil, err
	}

	return &schemaClient{
		client: schemaserviceconnect.NewSchemaServiceClient(
			opts.HTTPClient(),
			opts.Endpoint,
			connect.WithCodec(grpc.Codec{}), // enable vtprotobuf codec
			connect.WithInterceptors(sdkerror.ConnectErrorInterceptor()),
		),
		storeID: stately.StoreID(storeID),
	}, nil
}
