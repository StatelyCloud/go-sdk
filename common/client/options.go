package client

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strings"

	"github.com/StatelyCloud/go-sdk/common/auth"
	"golang.org/x/net/http2"
)

// Options is a set of common options for a stately API client.
type Options struct {
	// Endpoint is the Stately API endpoint.
	// Defaults to https://api.stately.cloud
	Endpoint          string
	AuthTokenProvider auth.TokenProvider
}

// ApplyDefaults applies the default values to the options.
func (o *Options) ApplyDefaults(appCtx context.Context) (*Options, error) {
	var err error
	if o == nil {
		o = &Options{}
	}
	if o.Endpoint == "" {
		o.Endpoint = "https://api.stately.cloud"
	}
	if o.AuthTokenProvider == nil {
		o.AuthTokenProvider, err = auth.NewAuthTokenProvider(appCtx, nil)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
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
		// Install auth middleware at the HTTP client layer
		Transport: auth.WrapTransportWithAuthTokenMiddleware(o.AuthTokenProvider, http2Transport),
	}

	return httpClient
}
