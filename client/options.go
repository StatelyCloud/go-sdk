package client

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
)

const (
	clientIDEnvVar     = "STATELY_CLIENT_ID"
	clientSecretEnvVar = "STATELY_CLIENT_SECRET"
)

// Options is a set of common options for a stately API client.
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

	// Endpoint is the Stately API endpoint.
	// Defaults to https://api.stately.cloud
	Endpoint string
}

// AuthTokenProvider is the interface which must be passed into the WithAuthToken{Unary,Stream}Interceptor
// so that it can authenticate outgoing requests.
// This is a thread-safe interface.
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

// ApplyDefaults applies the default values to the options.
func (o *Options) ApplyDefaults(appCtx context.Context) (*Options, error) {
	if o == nil {
		o = &Options{}
	}
	if o.Endpoint == "" {
		o.Endpoint = "https://api.stately.cloud"
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

	// Install auth middleware at the HTTP client layer if an auth token provider is set.
	var roundTripper http.RoundTripper = http2Transport
	if o.AuthTokenProvider != nil {
		roundTripper = wrapTransportWithAuthTokenMiddleware(o.AuthTokenProvider, http2Transport)
	}

	httpClient := &http.Client{
		Transport: roundTripper,
	}

	return httpClient
}

// wrapTransportWithAuthTokenMiddleware adds an HTTP middleware that will
// automatically retrieve valid access tokens and attach them to outgoing
// requests.
func wrapTransportWithAuthTokenMiddleware(tokenProvider AuthTokenProvider, next http.RoundTripper) http.RoundTripper {
	return &httpAuthMiddleware{
		tokenProvider: tokenProvider,
		next:          next,
	}
}

type httpAuthMiddleware struct {
	tokenProvider AuthTokenProvider
	next          http.RoundTripper
}

var _ http.RoundTripper = &httpAuthMiddleware{}

func (m *httpAuthMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := m.tokenProvider.GetAccessToken(req.Context())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := m.next.RoundTrip(req)

	// If the RPC failed due to auth, attempt to refresh the access token and retry once.
	if resp != nil && resp.StatusCode == http.StatusUnauthorized {
		m.tokenProvider.InvalidateAccessToken()
		token, err = m.tokenProvider.GetAccessToken(req.Context())
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = m.next.RoundTrip(req)
	}

	return resp, err
}
