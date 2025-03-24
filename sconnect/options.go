package sconnect

import (
	"connectrpc.com/connect"
	"github.com/planetscale/vtprotobuf/codec/grpc"

	"github.com/StatelyCloud/go-sdk/sdkerror"
)

// ConnectClientOptions is a set of standard options for our Connect clients.
var ConnectClientOptions = []connect.ClientOption{
	// enable vtprotobuf codec
	connect.WithCodec(grpc.Codec{}),
	// convert Connect errors to SDK errors
	connect.WithInterceptors(sdkerror.ConnectErrorInterceptor()),
	// Gzip large requests
	connect.WithSendGzip(),
	// By default Connect compresses everything, which is unnecessary for small messages
	connect.WithCompressMinBytes(1024),
}

// LocalConnectClientOptions are a set of options to apply when connecting to a local-host server.
// This eliminates compression.
var LocalConnectClientOptions = []connect.ClientOption{
	// enable vtprotobuf codec
	connect.WithCodec(grpc.Codec{}),
	// convert Connect errors to SDK errors
	connect.WithInterceptors(sdkerror.ConnectErrorInterceptor()),
}
