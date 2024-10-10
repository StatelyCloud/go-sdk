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
	client schemaserviceconnect.SchemaServiceClient
}

// Client models all schema management related apis.
type Client interface {
	// Get retrieves the fully self-contained Schema version for the corresponding Schema
	// ID. The call will fail if the caller does not have permission the SchemaID.
	Get(ctx context.Context, id stately.SchemaID, version stately.SchemaVersionID) (*schemaservice.SchemaModel, error)

	// Put appends a Schema version to the StatelyDB Schema. If the caller attempts
	// to put a SchemaVersion for a Schema that does not exist the request will fail.
	// If the caller does not have permissions to access the SchemaID the request will fail.
	// If the Schema is not valid the request will fail. If a previous Schema
	// already exists, then the new version will only be accepted if the
	// new Schema is backwards compatible.
	Put(
		ctx context.Context,
		fileDescriptor *descriptorpb.FileDescriptorProto,
		changeDescription string,
		schemaID stately.SchemaID,
		options ...*PutOptions,
	) (*schemaservice.PutResponse, error)

	// Validate ensures the provided file descriptor proto is a valid stately schema
	// and returns a list of validation errors. This does not persist the schema
	// nor does it require auth.
	Validate(
		ctx context.Context,
		fileDescriptor *descriptorpb.FileDescriptorProto,
	) (*schemaservice.ValidateResponse, error)

	// ListAuditLog retrieves the audit log for the Schema associated with the provided schemaID.
	// The audit log consists of a list of audit log entries that represent each change to the Schema including
	// its creation. The list is ordered by the time of the change with the most recent change first.
	// If there is no Schema for the provided Schema ID, an empty list will be returned.
	ListAuditLog(
		ctx context.Context,
		schemaID stately.SchemaID,
		options ...*ListAuditLogOptions,
	) ([]*schemaservice.SchemaAuditLogEntry, error)

	// Bind associates a Schema with a Store. The caller must have permission
	// to access the SchemaID and StoreID. If a different Schema is already
	// bound to the Store, the request will fail. The schemaID must exist
	// in the same organization as the storeID or the request will fail.
	//
	// (option) force will perform the bind operation, even if the store is
	// already bound to a schema. This is very dangerous if your store already
	// has items that are not compatible with the new schema. You must be
	// exceedingly sure the store is empty or that ALL items are compatible with
	// the new schema.
	Bind(
		ctx context.Context,
		schemaID stately.SchemaID,
		storeID stately.StoreID,
		force bool,
	) error

	// Create creates a new Schema with the given name and description.
	// The caller must have permission to create a Schema in the organization
	Create(
		ctx context.Context,
		projectID stately.ProjectID,
		name string,
		description string,
	) (stately.SchemaID, error)
}

// NewClient creates a new client with the given options.
func NewClient(appCtx context.Context, options ...*stately.Options) (Client, error) {
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
	}, nil
}
