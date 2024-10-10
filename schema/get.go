package schema

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
	"github.com/StatelyCloud/go-sdk/stately"
)

func (c *schemaClient) Get(
	ctx context.Context,
	id stately.SchemaID,
	versionID stately.SchemaVersionID,
) (*schemaservice.SchemaModel, error) {
	resp, err := c.client.Get(ctx, connect.NewRequest(&schemaservice.GetRequest{
		SchemaId:        uint64(id),
		SchemaVersionId: uint32(versionID),
	}))
	if err != nil {
		return nil, err
	}

	return resp.Msg.Schema, nil
}
