package management

import (
	"context"

	"google.golang.org/grpc"

	"github.com/StatelyCloud/go-sdk/common/auth"
	"github.com/StatelyCloud/go-sdk/common/net"
	"github.com/StatelyCloud/go-sdk/db/common"
	pb "github.com/StatelyCloud/go-sdk/pb/dbmanagement"
)

type clientImpl struct {
	client pb.ManagementClient
}

// Client is a Stately management client that performs DB management operations.
type Client interface {
	// DescribeStoreByID maps to Management API.
	DescribeStoreByID(ctx context.Context, storeID uint64) (common.Store, error)
	// DescribeStoreByName maps to Management API.
	DescribeStoreByName(ctx context.Context, projectID uint64, name string) (common.Store, error)
	// ListStores maps to Management API.
	ListStores(ctx context.Context, projectID uint64) ([]common.Store, error)
	// DeleteStore maps to Management API.
	DeleteStore(ctx context.Context, storeID uint64) error
	// CreateStore maps to Management API.
	CreateStore(ctx context.Context, projectID uint64, name, description string) (common.Store, error)
}

// NewClientWithAuthProvider builds a new client that authenticates with the given AuthTokenProvider.
func NewClientWithAuthProvider(
	appCtx context.Context,
	provider auth.TokenProvider,
	options ...net.StatelyConnectionOption,
) (Client, error) {
	conn, err := net.NewStatelyConnection(appCtx, provider, options...)
	if err != nil {
		return nil, err
	}
	return NewClientWithConn(conn), nil
}

// NewClientWithCreds builds a new client that authenticates with the given auth credentials.
func NewClientWithCreds(
	appCtx context.Context,
	clientID, clientSecret string,
	options ...net.StatelyConnectionOption,
) (Client, error) {
	return NewClientWithAuthProvider(appCtx, auth.NewAuth0TokenProvider(appCtx, clientID, clientSecret), options...)
}

// NewClientWithConn builds a new client with the provided connection to the stately API.
func NewClientWithConn(conn *grpc.ClientConn) Client {
	return &clientImpl{
		client: pb.NewManagementClient(conn),
	}
}

func (c *clientImpl) DescribeStoreByID(ctx context.Context, storeID uint64) (common.Store, error) {
	response, err := c.client.DescribeStore(ctx, &pb.DescribeStoreRequest{
		Identifier: &pb.DescribeStoreRequest_StoreId{
			StoreId: storeID,
		},
	})
	if err != nil {
		return common.Store{}, err
	}
	return common.Store{
		ID:          response.Store.StoreId,
		Name:        response.Store.Name,
		Description: response.Store.Description,
	}, nil
}

func (c *clientImpl) DescribeStoreByName(ctx context.Context, projectID uint64, name string) (common.Store, error) {
	response, err := c.client.DescribeStore(ctx, &pb.DescribeStoreRequest{
		Identifier: &pb.DescribeStoreRequest_StoreLookup{
			StoreLookup: &pb.StoreLookup{
				ProjectId: projectID,
				Name:      name,
			},
		},
	})
	if err != nil {
		return common.Store{}, err
	}
	return common.Store{
		ID:          response.Store.StoreId,
		Name:        response.Store.Name,
		Description: response.Store.Description,
	}, nil
}

func (c *clientImpl) ListStores(ctx context.Context, projectID uint64) ([]common.Store, error) {
	response, err := c.client.ListStores(ctx, &pb.ListStoresRequest{
		ProjectId: projectID,
	})
	if err != nil {
		return []common.Store{}, err
	}
	storeInfos := make([]common.Store, len(response.Stores))
	for i, v := range response.Stores {
		storeInfos[i] = common.Store{
			ID:          v.StoreId,
			Name:        v.Name,
			Description: v.Description,
		}
	}
	return storeInfos, nil
}

func (c *clientImpl) DeleteStore(ctx context.Context, storeID uint64) error {
	// DeleteStoreResponse is empty, so there is nothing to do with the response
	_, err := c.client.DeleteStore(ctx, &pb.DeleteStoreRequest{
		StoreId: storeID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *clientImpl) CreateStore(
	ctx context.Context,
	projectID uint64,
	name, description string,
) (common.Store, error) {
	response, err := c.client.CreateStore(ctx, &pb.CreateStoreRequest{
		ProjectId:   projectID,
		Name:        name,
		Description: description,
	})
	if err != nil {
		return common.Store{}, err
	}
	return common.Store{
		ID:          response.StoreId,
		Name:        name,
		Description: description,
	}, nil
}
