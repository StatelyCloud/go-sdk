package schema

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
	"github.com/StatelyCloud/go-sdk/stately"
)

func (c *schemaClient) Bind(
	ctx context.Context,
	schemaID stately.SchemaID,
	storeID stately.StoreID,
	force bool,
) error {
	_, err := c.client.Bind(ctx, connect.NewRequest(&schemaservice.BindRequest{
		SchemaId: uint64(schemaID),
		StoreId:  uint64(storeID),
		Force:    force,
	}))
	return err
}
