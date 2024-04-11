package dbmanagement

import (
	"context"

	"connectrpc.com/connect"
	"github.com/StatelyCloud/go-sdk/client"
	pb "github.com/StatelyCloud/go-sdk/pb/dbmanagement"
	"github.com/StatelyCloud/go-sdk/pb/dbmanagement/dbmanagementconnect"
	"github.com/planetscale/vtprotobuf/codec/grpc"
)

type clientImpl struct {
	client dbmanagementconnect.ManagementClient
}

// Client is a Stately management client that performs DB management operations.
type Client interface {
	// DeleteStore maps to Management API.
	DeleteStore(ctx context.Context, storeID client.StoreID) error
	// CreateStore maps to Management API.
	CreateStore(ctx context.Context, projectID client.ProjectID, name, description string) (*StoreInfo, error)
}

// NewClient creates a new client with the given store and options.
func NewClient(appCtx context.Context, options ...*client.Options) (Client, error) {
	opts := &client.Options{}
	for _, o := range options {
		opts = opts.Merge(o)
	}
	opts, err := opts.ApplyDefaults(appCtx)
	if err != nil {
		return nil, err
	}
	return &clientImpl{
		client: dbmanagementconnect.NewManagementClient(
			opts.HTTPClient(),
			opts.Endpoint,
			connect.WithCodec(grpc.Codec{}),
		),
	}, nil
}

func (c *clientImpl) DeleteStore(ctx context.Context, storeID client.StoreID) error {
	// DeleteStoreResponse is empty, so there is nothing to do with the response
	_, err := c.client.DeleteStore(ctx, connect.NewRequest(&pb.DeleteStoreRequest{
		StoreId: uint64(storeID),
	}))
	if err != nil {
		return err
	}
	return nil
}

func (c *clientImpl) CreateStore(
	ctx context.Context,
	projectID client.ProjectID,
	name, description string,
) (*StoreInfo, error) {
	response, err := c.client.CreateStore(ctx, connect.NewRequest(&pb.CreateStoreRequest{
		ProjectId:   uint64(projectID),
		Name:        name,
		Description: description,
	}))
	if err != nil {
		return nil, err
	}
	return &StoreInfo{
		ID:          client.StoreID(response.Msg.StoreId),
		Name:        name,
		Description: description,
	}, nil
}
