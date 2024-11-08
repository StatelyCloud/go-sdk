// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth/service.proto

package authconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	auth "github.com/StatelyCloud/go-sdk/pb/auth"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "stately.auth.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceGetAccessTokenProcedure is the fully-qualified name of the AuthService's
	// GetAccessToken RPC.
	AuthServiceGetAccessTokenProcedure = "/stately.auth.AuthService/GetAccessToken"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	authServiceServiceDescriptor              = auth.File_auth_service_proto.Services().ByName("AuthService")
	authServiceGetAccessTokenMethodDescriptor = authServiceServiceDescriptor.Methods().ByName("GetAccessToken")
)

// AuthServiceClient is a client for the stately.auth.AuthService service.
type AuthServiceClient interface {
	// GetAccessToken returns a short-lived access token from some proof of
	// identity. This operation will fail if the identity cannot be verified.
	GetAccessToken(context.Context, *connect.Request[auth.GetAccessTokenRequest]) (*connect.Response[auth.GetAccessTokenResponse], error)
}

// NewAuthServiceClient constructs a client for the stately.auth.AuthService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		getAccessToken: connect.NewClient[auth.GetAccessTokenRequest, auth.GetAccessTokenResponse](
			httpClient,
			baseURL+AuthServiceGetAccessTokenProcedure,
			connect.WithSchema(authServiceGetAccessTokenMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	getAccessToken *connect.Client[auth.GetAccessTokenRequest, auth.GetAccessTokenResponse]
}

// GetAccessToken calls stately.auth.AuthService.GetAccessToken.
func (c *authServiceClient) GetAccessToken(ctx context.Context, req *connect.Request[auth.GetAccessTokenRequest]) (*connect.Response[auth.GetAccessTokenResponse], error) {
	return c.getAccessToken.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the stately.auth.AuthService service.
type AuthServiceHandler interface {
	// GetAccessToken returns a short-lived access token from some proof of
	// identity. This operation will fail if the identity cannot be verified.
	GetAccessToken(context.Context, *connect.Request[auth.GetAccessTokenRequest]) (*connect.Response[auth.GetAccessTokenResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceGetAccessTokenHandler := connect.NewUnaryHandler(
		AuthServiceGetAccessTokenProcedure,
		svc.GetAccessToken,
		connect.WithSchema(authServiceGetAccessTokenMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/stately.auth.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceGetAccessTokenProcedure:
			authServiceGetAccessTokenHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) GetAccessToken(context.Context, *connect.Request[auth.GetAccessTokenRequest]) (*connect.Response[auth.GetAccessTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("stately.auth.AuthService.GetAccessToken is not implemented"))
}
