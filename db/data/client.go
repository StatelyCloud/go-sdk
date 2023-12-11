package data

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/StatelyCloud/go-sdk/common/auth"
	"github.com/StatelyCloud/go-sdk/common/identifiers"
	"github.com/StatelyCloud/go-sdk/common/net"
	pb "github.com/StatelyCloud/go-sdk/pb/data"
)

type clientImpl struct {
	client  pb.DataClient
	storeID identifiers.StoreID
}

// Client is a stately data client that interacts with the given store.
type Client interface {
	// Delete removes one or more Items from the Store by their full key paths. This
	// will fail if any Item does not exist, if not all of the DeleteItem requests
	// are under the same root item path, or if the caller does not have permission
	// to delete Items. Tombstones will be left for deleted items for some
	// predetermined time (TBD tombstone behavior). All deletes in the request are
	// applied atomically - there are no partial successes.
	// TODO: We should move away from having positional arguments in favor of
	// request objects, since at some point we'll have a lot more options than just
	// a list of key paths.
	Delete(ctx context.Context, itemPaths ...KeyPath) error

	// Get retrieves one or more Items by their full key paths. This will return any
	// of the Items that exist. It will fail if not all of the GetItem requests are
	// under the same root item path, or if the caller does not have permission to
	// read Items. Use Query if you want to retrieve multiple items but don't
	// already know the full key paths of the items you want to get.
	Get(ctx context.Context, itemPaths ...KeyPath) ([]*Item, error)

	// Put adds one or more Items to the Store, or replaces the Items if they
	// already exist at that path. This will fail if not all of the PutItem requests
	// are under the same root item path, if any of the PutItem requests' write
	// conditions fails, or if the caller does not have permission to create Items.
	// All puts in the request are applied atomically - there are no partial
	// successes. Data can be provided as either JSON, or as a proto encoded by a
	// previously agreed upon schema, or by some combination of the two.
	Put(ctx context.Context, items ...*Item) error

	// Query loads Items that start with a specified key path, subject to additional
	// filtering. The prefix must minimally contain two segments, an item type and
	// an item ID. Query will return an empty result set if there are no items
	// matching that key prefix. It is paginated, so you may have to call it
	// repeatedly, passing a pagination token each time. This can also fail if the
	// caller does not have permission to read Items.
	Query(ctx context.Context, keyPathPrefix KeyPath) ([]*Item, error)
}

// NewClientWithAuthProvider builds a new client that authenticates with the given AuthTokenProvider.
func NewClientWithAuthProvider(
	appCtx context.Context,
	provider auth.TokenProvider,
	storeID identifiers.StoreID,
	options ...net.StatelyConnectionOption,
) (Client, error) {
	conn, err := net.NewStatelyConnection(appCtx, provider, options...)
	if err != nil {
		return nil, err
	}
	return NewClientWithConn(conn, storeID), nil
}

// NewClientWithCreds builds a new client that authenticates with the given auth credentials.
func NewClientWithCreds(
	appCtx context.Context,
	clientID, clientSecret string,
	storeID identifiers.StoreID,
	options ...net.StatelyConnectionOption,
) (Client, error) {
	return NewClientWithAuthProvider(
		appCtx,
		auth.NewAuth0TokenProvider(appCtx, clientID, clientSecret),
		storeID,
		options...)
}

// NewClientWithConn builds a new client with the provided connection to the stately API.
func NewClientWithConn(conn *grpc.ClientConn, storeID identifiers.StoreID) Client {
	return &clientImpl{
		client:  pb.NewDataClient(conn),
		storeID: storeID,
	}
}

func (c *clientImpl) Delete(ctx context.Context, itemPaths ...KeyPath) error {
	deleteItems := make([]*pb.DeleteItem, len(itemPaths))
	for i, v := range itemPaths {
		deleteItems[i] = &pb.DeleteItem{KeyPath: v.String()}
	}

	_, err := c.client.Delete(ctx, &pb.DeleteRequest{
		StoreId: uint64(c.storeID),
		Deletes: deleteItems,
	})
	return err
}

func (c *clientImpl) Get(ctx context.Context, itemPaths ...KeyPath) ([]*Item, error) {
	getItems := make([]*pb.GetItem, len(itemPaths))
	for i, v := range itemPaths {
		getItems[i] = &pb.GetItem{KeyPath: v.String()}
	}

	response, err := c.client.Get(ctx, &pb.GetRequest{
		StoreId: uint64(c.storeID),
		Gets:    getItems,
	})
	if err != nil {
		return nil, err
	}

	items := make([]*Item, len(response.Results))
	for i, v := range response.Results {
		items[i], err = protoItemToItem(v.GetItem())
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (c *clientImpl) Put(ctx context.Context, items ...*Item) error {
	putItems := make([]*pb.PutItem, len(items))
	for i, v := range items {
		jsonStruct := structpb.Struct{}
		err := json.Unmarshal([]byte(v.JSON), &jsonStruct)
		if err != nil {
			return err
		}

		putItems[i] = &pb.PutItem{
			Item: &pb.Item{
				KeyPath: v.KeyPath.String(),
				Json:    &jsonStruct,
			},
		}
	}

	_, err := c.client.Put(ctx, &pb.PutRequest{
		StoreId: uint64(c.storeID),
		Puts:    putItems,
	})

	return err
}

func (c *clientImpl) Query(ctx context.Context, keyPathPrefix KeyPath) ([]*Item, error) {
	response, err := c.client.Query(ctx, &pb.QueryRequest{
		StoreId:       uint64(c.storeID),
		KeyPathPrefix: keyPathPrefix.String(),
	})
	if err != nil {
		return nil, err
	}

	items := make([]*Item, len(response.Results))
	for i, v := range response.Results {
		items[i], err = protoItemToItem(v.GetItem())
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func protoItemToItem(protoItem *pb.Item) (*Item, error) {
	item := &Item{}
	item.KeyPath = KeyPath(protoItem.GetKeyPath())
	jsonStr, err := protoItem.GetJson().MarshalJSON()
	if err != nil {
		return nil, err
	}
	item.JSON = string(jsonStr)
	return item, nil
}
