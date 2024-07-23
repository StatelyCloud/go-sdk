package schema

import (
	"context"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/StatelyCloud/go-sdk/pb/schemaservice"
)

func (c *schemaClient) Validate(
	ctx context.Context,
	fileDescriptor *descriptorpb.FileDescriptorProto,
) (*schemaservice.ValidateResponse, error) {
	resp, err := c.client.Validate(ctx, connect.NewRequest(&schemaservice.ValidateRequest{
		FileDescriptor: fileDescriptor,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}
