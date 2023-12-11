package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcMetadata "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	// This import installs the vtproto grpc codec.
	_ "github.com/StatelyCloud/stately/gocommon/sgrpc/vtproto"
)

// WithOAuthRefreshUnaryInterceptor returns a dial option that will automatically authenticate
// using the given token provider and attach bearer tokens to outgoing requests.
// Requests that fail with Unauthenticated will have their token refreshed and then be retried
// once before propagating the error.
func WithOAuthRefreshUnaryInterceptor(tokenProvider TokenProvider) grpc.DialOption {
	return grpc.WithUnaryInterceptor(
		func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			token, err := tokenProvider.GetAccessToken(ctx)
			if err != nil {
				return err
			}
			ctxWithAuth := grpcMetadata.AppendToOutgoingContext(
				ctx,
				"authorization",
				"Bearer "+token,
			)

			// Attempt RPC.

			err = invoker(ctxWithAuth, method, req, reply, cc, opts...)

			// If the RPC failed due to auth, attempt to refresh the access token and retry once.
			replyStatus, _ := status.FromError(err)
			if replyStatus.Code() == codes.Unauthenticated {
				token, err_ := tokenProvider.RefreshAccessToken(ctx, true)
				if err_ != nil {
					return err_
				}
				ctxWithAuth = grpcMetadata.AppendToOutgoingContext(
					ctx,
					"authorization",
					"Bearer "+token,
				)
				err = invoker(ctxWithAuth, method, req, reply, cc, opts...)
			}

			return err
		},
	)
}

// TODO - add a stream interceptor
