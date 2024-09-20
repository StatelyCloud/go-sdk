package stately

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/http2"

	"github.com/StatelyCloud/go-sdk/internal/auth"
	"github.com/StatelyCloud/go-sdk/pb/db"
	"github.com/StatelyCloud/go-sdk/pb/schemaservice/schemaserviceconnect"
)

const (
	clientIDEnvVar     = "STATELY_CLIENT_ID"
	clientSecretEnvVar = "STATELY_CLIENT_SECRET"
)

// Options is a set of common options for a stately API client.
// You can either construct a single Options struct and pass it to the client,
// or you can use the ApplyDefaults method to fill in the defaults.
// See: NewClient and ApplyDefaults for more information.
type Options struct {
	// ClientID is your Stately Client ID. If not set, this is loaded from the
	// STATELY_CLIENT_ID environment variable when using the default
	// AuthTokenProvider.
	ClientID string

	// ClientSecret is your Stately Client ID. If not set, this is loaded from the
	// STATELY_CLIENT_SECRET environment variable when using the default
	// AuthTokenProvider.
	ClientSecret string

	// AuthTokenProvider handles fetching auth tokens for requests. It is
	// defaulted to an appropriate implementation for most services.
	AuthTokenProvider

	// Either Region or Endpoint should be set, but not both. Doing so will
	// result in an error. An empty Region and Endpoint will default to
	// https://api.stately.cloud
	//
	// Region is the Stately environment to use. This is used to determine
	// the correct endpoint to use for SDK calls.
	//
	Region string

	// Either Endpoint or Region should be set, but not both. Doing so will
	// result in an error. An empty Region and Endpoint will default to
	// https://api.stately.cloud
	//
	// Endpoint is the full URL to the Stately API endpoint to use.
	Endpoint string

	// JSONResponseFormat is a flag to indicate that the item in the response
	// should be in JSON format. This can be used in conjunction with the
	// JSONItemMapper to unmarshal the item to/from a JSON representation.
	// This would be helpful if you are using the Stately client to send/receive
	// items that only accept JSON. However, this is not a typical use case as
	// you should use the item types generated by the Stately cli code generator.
	// Defaults to false.
	JSONResponseFormat bool
}

// unauthenticatedMethods do not require auth and therefore will be skipped
// when needing to add an auth token.
var unauthenticatedMethods = []string{
	// Our schema validate api is auth free!
	schemaserviceconnect.SchemaServiceValidateProcedure,
}

// AuthTokenProvider is the interface which the client uses to authenticate
// outgoing requests. The client will be wired with a default implementation
// that is suitable for most use cases. This is a thread-safe interface.
type AuthTokenProvider interface {
	// GetAccessToken returns an access token or an error. If there is no current
	// access token then the provider will block and attempt to refresh and get a
	// new access token. An error is returned if there is no access token and the
	// refresh operation fails.
	GetAccessToken(ctx context.Context) (string, error)

	// InvalidateAccessToken marks the current token, if any, as invalid. This
	// should be called if the service responds with a status that indicates the
	// token is no longer valid. This should cause the next call to GetAccessToken
	// to fetch a fresh token.
	InvalidateAccessToken()
}

// ItemTypeMapper is a function that maps a db.Item to your SDK generated types.
// We will generate this type mapper when using Stately's code generation to
// handle the unmarshalling for you.
type ItemTypeMapper func(item *db.Item) (Item, error)

// ApplyDefaults applies the default values (listed in Options) to the options.
func (o *Options) ApplyDefaults(appCtx context.Context) (*Options, error) {
	if o == nil {
		o = &Options{}
	}

	// If both endpoint and environment are provided, return an error.
	if o.Endpoint != "" && o.Region != "" && o.Endpoint != RegionToEndpoint(o.Region) {
		return nil, fmt.Errorf("both Endpoint: %q and Region: %q are set in options. "+
			"Please provide one or the other, "+
			"or neither to default to https://api.stately.cloud", o.Endpoint, o.Region)
	} else if o.Endpoint == "" {
		// If there's no endpoint specified, use the region.
		o.Endpoint = RegionToEndpoint(o.Region)
	}

	if o.AuthTokenProvider == nil {
		clientID := o.ClientID
		if clientID == "" {
			clientID = os.Getenv(clientIDEnvVar)
		}
		if clientID == "" {
			return nil, fmt.Errorf("unable to read client ID from %s env var", clientIDEnvVar)
		}

		clientSecret := o.ClientSecret
		if clientSecret == "" {
			clientSecret = os.Getenv(clientSecretEnvVar)
		}
		if clientSecret == "" {
			return nil, fmt.Errorf("unable to read client secret from %s env var", clientSecretEnvVar)
		}
		o.AuthTokenProvider = auth.NewAuthTokenProvider(appCtx, clientID, clientSecret, nil)
	}
	return o, nil
}

// Merge merges non-default settings from o2 into o.
func (o *Options) Merge(o2 *Options) *Options {
	if o2 == nil {
		return o
	}
	if o2.ClientID != "" {
		o.ClientID = o2.ClientID
	}
	if o2.ClientSecret != "" {
		o.ClientSecret = o2.ClientSecret
	}
	if o2.AuthTokenProvider != nil {
		o.AuthTokenProvider = o2.AuthTokenProvider
	}
	if o2.Endpoint != "" {
		o.Endpoint = o2.Endpoint
	}
	if o2.Region != "" {
		o.Region = o2.Region
	}
	if o2.JSONResponseFormat {
		o.JSONResponseFormat = o2.JSONResponseFormat
	}
	return o
}

// HTTPClient builds an HTTP/2 client for the given options.
func (o *Options) HTTPClient() *http.Client {
	// We want to use HTTP/2 (it's required for bidi streams anyway)
	http2Transport := &http2.Transport{
		// TODO: Timeout configs
		// TODO: Cert pinning
	}

	// This allows talking to localhost h2c servers
	if strings.HasPrefix(o.Endpoint, "http://") {
		http2Transport.AllowHTTP = true
		http2Transport.DialTLSContext = func(_ context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		}
	}

	httpClient := &http.Client{
		Transport: http2Transport,
	}

	// Install the necessary middleware:
	if o.AuthTokenProvider != nil {
		httpClient = wrapRoundTripper(httpClient, &httpAuthMiddleware{o.AuthTokenProvider})
	}

	if o.JSONResponseFormat {
		httpClient = wrapRoundTripperFunc(httpClient,
			func(req *http.Request, next http.RoundTripper) (*http.Response, error) {
				req.Header.Set("sc-rf", "application/json")
				return next.RoundTrip(req)
			})
	}

	return httpClient
}

type httpAuthMiddleware struct {
	tokenProvider AuthTokenProvider
}

func (m *httpAuthMiddleware) RoundTripper(req *http.Request, next http.RoundTripper) (*http.Response, error) {
	// filter out apis that don't require auth
	for _, path := range unauthenticatedMethods {
		if req.URL.Path == path {
			return next.RoundTrip(req)
		}
	}

	token, err := m.tokenProvider.GetAccessToken(req.Context())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := next.RoundTrip(req)

	// If the RPC failed due to auth, attempt to refresh the access token and retry once.
	if resp != nil && resp.StatusCode == http.StatusUnauthorized {
		m.tokenProvider.InvalidateAccessToken()
		token, err = m.tokenProvider.GetAccessToken(req.Context())
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = next.RoundTrip(req)
	}

	return resp, err
}
