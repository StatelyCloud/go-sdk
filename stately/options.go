package stately

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/http2"

	"github.com/StatelyCloud/go-sdk/internal/auth"
	"github.com/StatelyCloud/go-sdk/pb/db"
)

// Options is a set of common options for a stately API client.
// You can either construct a single Options struct and pass it to the client,
// or you can use the ApplyDefaults method to fill in the defaults.
// See: NewClient and ApplyDefaults for more information.
type Options struct {
	// AccessKey is your Stately Access Key. If this is not set, it will be loaded
	// from the STATELY_ACCESS_KEY environment variable.
	AccessKey string

	// NoAuth is a flag to indicate that the client should not attempt to get an
	// auth token. This is used when talking to the Stately BYOC Data Plane on
	// localhost.
	NoAuth bool

	// AuthTokenProvider handles fetching auth tokens for requests. It is
	// defaulted to an appropriate implementation for most services.
	AuthTokenProvider

	// Either Region or Endpoint should be set, but not both. Doing so will
	// result in an error. An empty Region and Endpoint will default to
	// https://api.stately.cloud
	//
	// Region is the cloud region to use. This is used to determine
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
	// you should use the item types generated by the Stately CLI code generator.
	// Defaults to false.
	JSONResponseFormat bool

	// (internal) NoAdmin is an internal flag used by Stately employees.
	// This flag has no effect for non-Stately Employees.
	NoAdmin bool

	// Cached transport to our service endpoint.
	transport *http2.Transport
}

// AuthTokenProvider is the functional interface which the client uses to
// authenticate outgoing requests. The client will be wired with a default
// implementation that is suitable for most use cases. This is a thread-safe
// interface. The `force` parameter causes current token to be invalidated and a
// new one to be synchronously fetched. This will block other incoming requests
// until the new token is fetched.
type AuthTokenProvider func(ctx context.Context, force bool) (string, error)

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

	o.transport = createTransport(o.Endpoint)

	if o.AuthTokenProvider == nil && !o.NoAuth {
		accessKey := o.AccessKey
		if accessKey == "" {
			accessKey = os.Getenv("STATELY_ACCESS_KEY")
		}
		if accessKey != "" {
			if o.transport == nil {
				o.transport = createTransport(o.Endpoint)
			}
			o.AuthTokenProvider = auth.AccessKeyAuth(
				appCtx,
				accessKey,
				o.Endpoint,
				o.transport,
				200*time.Millisecond,
			)
		}
		if o.AuthTokenProvider == nil {
			return nil, fmt.Errorf(
				"unable to find an access key in the STATELY_ACCESS_KEY environment variable. " +
					"Either pass your access key in the options when creating a client or set this environment variable. " +
					"Alternatively, set NoAuth to true in the options if you are using the Stately BYOC Data Plane on localhost",
			)
		}
	}
	return o, nil
}

// Merge merges non-default settings from o2 into o.
func (o *Options) Merge(o2 *Options) *Options {
	if o2 == nil {
		return o
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
	if o2.AccessKey != "" {
		o.AccessKey = o2.AccessKey
	}
	o.NoAuth = o2.NoAuth
	o.NoAdmin = o2.NoAdmin
	o.transport = nil
	return o
}

func isLocalEndpoint(endpoint string) bool {
	return strings.HasPrefix(endpoint, "http://localhost") ||
		strings.HasPrefix(endpoint, "http://0.0.0.0")
}

func createTransport(endpoint string) *http2.Transport {
	// We want to use HTTP/2 (it's required for bidi streams anyway)
	http2Transport := &http2.Transport{
		// TODO: Timeout configs
		// TODO: Cert pinning
	}

	// This allows talking to localhost h2c servers
	if strings.HasPrefix(endpoint, "http://") {
		// Disable compression locally
		http2Transport.DisableCompression = isLocalEndpoint(endpoint)
		http2Transport.AllowHTTP = true
		http2Transport.DialTLSContext = func(_ context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		}
	}
	return http2Transport
}

// HTTPClient builds an HTTP/2 client for the given options.
func (o *Options) HTTPClient() *http.Client {
	if o.transport == nil {
		o.transport = createTransport(o.Endpoint)
	}

	httpClient := &http.Client{
		Transport: o.transport,
	}

	// Install the necessary middleware:
	if o.AuthTokenProvider != nil {
		// TODO: fix this to take the unauthenticated methods as an arg
		httpClient = wrapRoundTripper(
			httpClient,
			&httpAuthMiddleware{o.AuthTokenProvider, []string{"/stately.schemaservice.SchemaService/Validate"}},
		)
	}

	if o.JSONResponseFormat {
		httpClient = wrapRoundTripperFunc(httpClient,
			func(req *http.Request, next http.RoundTripper) (*http.Response, error) {
				req.Header.Set("sc-rf", "application/json")
				return next.RoundTrip(req)
			})
	}

	if o.NoAdmin {
		httpClient = wrapRoundTripperFunc(httpClient,
			func(req *http.Request, next http.RoundTripper) (*http.Response, error) {
				req.Header.Set("X-Stately-NoAdmin", "1")
				return next.RoundTrip(req)
			})
	}

	return httpClient
}

type httpAuthMiddleware struct {
	getToken AuthTokenProvider
	// unauthenticatedMethods do not require auth and therefore will be skipped
	// when needing to add an auth token.
	unauthenticatedMethods []string
}

// connect.PayloadCloser implements a Rewind method that allows the body to be reused.
// https://github.com/connectrpc/connect-go/blob/74a6754f29185b85fefa2915bf8fb680a36ca8f0/duplex_http_call.go#L456
type rewindable interface {
	Rewind() bool
}

func (m *httpAuthMiddleware) RoundTripper(req *http.Request, next http.RoundTripper) (*http.Response, error) {
	// filter out apis that don't require auth
	for _, path := range m.unauthenticatedMethods {
		if req.URL.Path == path {
			return next.RoundTrip(req)
		}
	}

	token, err := m.getToken(req.Context(), false /*force*/)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Make the original request
	resp, err := next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// If the RPC failed due to auth, force refresh the access token and retry once.
	if resp != nil && resp.StatusCode == http.StatusUnauthorized {
		token, err = m.getToken(req.Context(), true /*force*/)
		if err != nil {
			return nil, err
		}
		// We need to reset the body or we'll get "request declared a Content-Length
		// of 57 but only wrote 0 bytes" on requests with a body.
		if rewinder, ok := req.Body.(rewindable); ok {
			if rewinder.Rewind() {
				// Redrive the request with the new token
				req.Header.Set("Authorization", "Bearer "+token)
				resp, err = next.RoundTrip(req)
			}

			// If we can't rewind the body, just return the original error. Sad.
		}
	}

	return resp, err
}
