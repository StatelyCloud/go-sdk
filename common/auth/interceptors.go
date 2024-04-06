package auth

import (
	"net/http"
)

// WrapTransportWithAuthTokenMiddleware adds an HTTP middleware that will
// automatically retrieve valid access tokens and attach them to outgoing
// requests.
func WrapTransportWithAuthTokenMiddleware(tokenProvider TokenProvider, next http.RoundTripper) http.RoundTripper {
	return &httpAuthMiddleware{
		tokenProvider: tokenProvider,
		next:          next,
	}
}

type httpAuthMiddleware struct {
	tokenProvider TokenProvider
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
		token, err = m.tokenProvider.RefreshAccessToken(req.Context(), true)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = m.next.RoundTrip(req)
	}

	return resp, err
}
