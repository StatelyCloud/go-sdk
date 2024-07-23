package schema

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
)

func (c *schemaClient) Get(ctx context.Context) (*schemaservice.SchemaModel, error) {
	resp, err := c.client.Get(ctx, connect.NewRequest(&schemaservice.GetRequest{
		StoreId: uint64(c.storeID),
	}))
	if err != nil {
		return nil, err
	}

	return resp.Msg.Schema, nil
}
