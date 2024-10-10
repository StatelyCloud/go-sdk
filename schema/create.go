package schema

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
	"github.com/StatelyCloud/go-sdk/stately"
)

func (c *schemaClient) Create(
	ctx context.Context,
	id stately.ProjectID,
	name string,
	description string,
) (stately.SchemaID, error) {
	resp, err := c.client.Create(ctx, connect.NewRequest(&schemaservice.CreateRequest{
		ProjectId:   uint64(id),
		Description: description,
		Name:        name,
	}))
	if err != nil {
		return 0, err
	}

	return stately.SchemaID(resp.Msg.SchemaId), nil
}
