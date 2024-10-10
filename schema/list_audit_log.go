package schema

import (
	"context"

	"connectrpc.com/connect"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
	"github.com/StatelyCloud/go-sdk/stately"
)

func (c *schemaClient) ListAuditLog(
	ctx context.Context,
	schemaID stately.SchemaID,
	options ...*ListAuditLogOptions,
) ([]*schemaservice.SchemaAuditLogEntry, error) {
	// only take the first option or use the default
	// if no options are provided
	opts := &ListAuditLogOptions{}
	if len(options) > 0 {
		opts = options[0]
	}
	resp, err := c.client.ListAuditLog(ctx, connect.NewRequest(&schemaservice.ListAuditLogRequest{
		SchemaId: uint64(schemaID),
		Limit:    opts.Limit,
	}))
	if err != nil {
		return nil, err
	}

	return resp.Msg.Entries, nil
}

// ListAuditLogOptions are optional parameters for ListAuditLog.
type ListAuditLogOptions struct {
	// Limit is the maximum number of entries to return.
	Limit uint32
}
