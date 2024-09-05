package stately

import (
	"net/http"
)

// roundTripperFunc is a function that also implements the roundTripper.type interface.
type roundTripperFunc func(req *http.Request, next http.RoundTripper) (*http.Response, error)

// roundTripper is an interface that can be used to wrap an http.RoundTripper when
// a pure function is not sufficient.
type roundTripper interface {
	RoundTripper(req *http.Request, next http.RoundTripper) (*http.Response, error)
}

// roundTripWrapper implements the http.RoundTripper interface, but calls a
// roundTripperFunc which is a nicer interface to with for injecting middleware.
type roundTripWrapper struct {
	next http.RoundTripper
	fn   roundTripperFunc
}

// RoundTrip implements the http.RoundTripper interface.
func (r *roundTripWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.fn(req, r.next)
}

// wrapRoundTripper wraps an http.Client with a roundTripper.
func wrapRoundTripper(client *http.Client, f roundTripper) *http.Client {
	client.Transport = &roundTripWrapper{
		next: client.Transport,
		fn:   f.RoundTripper,
	}
	return client
}

// wrapRoundTripperFunc wraps an http.Client with a roundTripperFunc.
func wrapRoundTripperFunc(client *http.Client, f roundTripperFunc) *http.Client {
	client.Transport = &roundTripWrapper{
		next: client.Transport,
		fn:   f,
	}
	return client
}
