package net

import (
	"context"
	"crypto/tls"

	"github.com/StatelyCloud/go-sdk/common/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	// This import installs the vtproto grpc codec.
	_ "github.com/StatelyCloud/go-sdk/common/net/vtproto"
)

const (
	defaultEndpoint = "api.stately.cloud:443"
)

type connectionOptions struct {
	endpoint             string
	transportCredentials credentials.TransportCredentials
}

// StatelyConnectionOption is an option to be passed to NewStatelyConnection.
type StatelyConnectionOption = func(*connectionOptions)

// WithEndpoint creates an option to override the stately API endpoint.
func WithEndpoint(endpoint string) StatelyConnectionOption {
	return func(o *connectionOptions) {
		o.endpoint = endpoint
	}
}

// WithInsecure creates an options to disable TLS for the stately API connection.
func WithInsecure() StatelyConnectionOption {
	return func(o *connectionOptions) {
		o.transportCredentials = insecure.NewCredentials()
	}
}

// NewStatelyConnection creates a new grpc connection to the stately API.
func NewStatelyConnection(
	appCtx context.Context,
	authTokenProvider auth.TokenProvider,
	options ...StatelyConnectionOption,
) (*grpc.ClientConn, error) {
	// setup the default connection options
	connOpts := &connectionOptions{
		endpoint: defaultEndpoint,
		transportCredentials: credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS13,
		}),
	}

	// run any overrides
	for _, opt := range options {
		opt(connOpts)
	}

	// create the client
	return grpc.DialContext(
		appCtx,
		connOpts.endpoint,
		grpc.WithTransportCredentials(connOpts.transportCredentials),
		auth.WithOAuthRefreshUnaryInterceptor(authTokenProvider),
	)
}
