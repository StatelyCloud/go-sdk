package data

import (
	"context"

	"connectrpc.com/connect"

	pbdata "github.com/StatelyCloud/go-sdk/pb/data"
)

// GetRequest will fetch all specified keys and return any found.
type GetRequest struct {
	// ItemPaths is an array of the full path(s) to the item(s).
	ItemPaths []string
	// (option) AllowStale  indicates that you're okay with getting a slightly stale item,
	// that is, if you had just changed an item and then call GetItem, you might
	// get the old version of the item. This can result in improved performance,
	// availability, and cost.
	AllowStale AllowStale
}

// GetOptions is a struct that can be passed to Get. Default handling is to not allow stale items.
type GetOptions struct {
	// (option) AllowStale  indicates that you're okay with getting a slightly stale item,
	// that is, if you had just changed an item and then call GetItem, you might
	// get the old version of the item. This can result in improved performance,
	// availability, and cost.
	AllowStale AllowStale
}

// NewGetRequest is a convenience method in golang to allow simple construction of GetRequests via
// NewGetRequest("/message-1") vs GetRequest{ ItemPaths: []KeyPath{}{"/message-1"} }\
// You can then set any optional fields such as "AllowStale".
func NewGetRequest(itemPaths ...string) GetRequest {
	return GetRequest{
		ItemPaths: itemPaths,
	}
}

// Get is a convenience method for retrieving a single Item by its full key path.
// See GetBatch for more information.
func (c *dataClient) Get(ctx context.Context, itemPath string, option ...GetOptions) (*RawItem, error) {
	opts := GetOptions{}
	if len(option) > 0 {
		opts.AllowStale = option[0].AllowStale
	}

	items, err := c.GetBatch(ctx, GetRequest{
		ItemPaths:  []string{itemPath},
		AllowStale: opts.AllowStale,
	})
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], nil
}

// GetBatch retrieves one or more Items by their full key paths. This will return any
// of the Items that exist. It will fail if not all of the GetItem requests are
// under the same root item path, or if the caller does not have permission to
// read Items. Use Query if you want to retrieve multiple items but don't
// already know the full key paths of the items you want to get.
func (c *dataClient) GetBatch(ctx context.Context, request GetRequest) ([]*RawItem, error) {
	response, err := c.client.Get(ctx, connect.NewRequest(&pbdata.GetRequest{
		StoreId:    uint64(c.storeID),
		Gets:       mapToItemKey(request.ItemPaths),
		AllowStale: bool(request.AllowStale),
	}))
	if err != nil {
		return nil, err
	}

	items := make([]*RawItem, len(response.Msg.GetItems()))
	for i, v := range response.Msg.Items {
		items[i], err = protoToItem(v)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func mapToItemKey(keys []string) []*pbdata.GetItem {
	getItems := make([]*pbdata.GetItem, len(keys))
	for i, v := range keys {
		getItems[i] = &pbdata.GetItem{KeyPath: v}
	}
	return getItems
}
