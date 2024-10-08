package schema

import (
	"context"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
	"github.com/StatelyCloud/go-sdk/stately"
)

func (c *schemaClient) Put(
	ctx context.Context,
	fileDescriptor *descriptorpb.FileDescriptorProto,
	changeDescription string,
	schemaID stately.SchemaID,
	options ...*PutOptions,
) (*schemaservice.PutResponse, error) {
	// only take the first option or use the default
	// if no options are provided
	opts := &PutOptions{}
	if len(options) > 0 {
		opts = options[0]
	}

	resp, err := c.client.Put(ctx, connect.NewRequest(&schemaservice.PutRequest{
		SchemaId:                   uint64(schemaID),
		FileDescriptor:             fileDescriptor,
		ChangeDescription:          changeDescription,
		DryRun:                     opts.DryRun,
		AllowBackwardsIncompatible: opts.Force,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg, err
}

// PutOptions are optional parameters for Put.
type PutOptions struct {
	// If DryRun is true the request will be validated against the existing schema but not applied.
	DryRun bool
	// If force is true, the schema will be applied even if it is backwards incompatible.
	Force bool
}
