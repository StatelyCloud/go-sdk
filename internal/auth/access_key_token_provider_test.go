package auth_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/StatelyCloud/go-sdk/internal/auth"
	pbauth "github.com/StatelyCloud/go-sdk/pb/auth"
	"github.com/StatelyCloud/go-sdk/pb/auth/authconnect"
)

type testAuthServer struct {
	tokenFunc  func() (string, error)
	expiresInS int
}

func (s *testAuthServer) GetAuthToken(
	context.Context,
	*connect.Request[pbauth.GetAuthTokenRequest],
) (*connect.Response[pbauth.GetAuthTokenResponse], error) {
	token := "test-token"
	if s.tokenFunc != nil {
		var err error
		token, err = s.tokenFunc()
		if err != nil {
			return nil, err
		}
	}
	expires := 1000
	if s.expiresInS != 0 {
		expires = s.expiresInS
	}

	return connect.NewResponse(&pbauth.GetAuthTokenResponse{
		AuthToken:  token,
		ExpiresInS: uint64(expires),
	}), nil
}

func makeTestServer(testServer *testAuthServer) *httptest.Server {
	mux := http.NewServeMux()
	path, handler := authconnect.NewAuthServiceHandler(testServer)
	mux.Handle(path, handler)
	svr := httptest.NewUnstartedServer(h2c.NewHandler(mux, &http2.Server{}))
	svr.EnableHTTP2 = true
	svr.Start()
	return svr
}

func makeTransport() *http2.Transport {
	return &http2.Transport{
		AllowHTTP: true,
		DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
			dialer := &net.Dialer{}
			return dialer.DialContext(ctx, network, addr)
		},
	}
}

func TestAccessKeyGetToken(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	svr := makeTestServer(&testAuthServer{})
	defer svr.Close()

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		200*time.Millisecond,
	)

	token, err := getToken(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// now close the server and call GetToken again. it should work without a network request because the value is cached
	svr.Close()
	token, err = getToken(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)
}

func TestAccessKeyConcurrentRefresh(t *testing.T) {
	t.Parallel()
	count := atomic.Uint64{}

	ctx := t.Context()
	testServer := &testAuthServer{
		tokenFunc: func() (string, error) {
			token := strconv.Itoa(int(count.Load()))
			count.Add(1)
			return token, nil
		},
	}
	svr := makeTestServer(testServer)

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		200*time.Millisecond,
	)

	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			token, err := getToken(context.TODO(), false)
			assert.NoError(t, err)
			assert.Equal(t, "0", token)
		}()
	}

	wg.Wait()
}

func TestAccessKeyBackgroundRefresh(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	testServer := &testAuthServer{
		tokenFunc: func() (string, error) {
			wg.Done()
			return "test-token", nil
		},
		expiresInS: 1,
	}
	svr := makeTestServer(testServer)
	defer svr.Close()

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		200*time.Millisecond,
	)

	token, err := getToken(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// we won't be able to continue past this point until the automatic refresh happens
	wg.Wait()
}

func TestAccessKeyRefreshContextCancelled(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.TODO())
	t.Cleanup(cancel)
	testServer := &testAuthServer{
		expiresInS: 0,
	}
	svr := makeTestServer(testServer)
	defer svr.Close()
	cancel()

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		200*time.Millisecond,
	)

	token, err := getToken(ctx, false)
	assert.Empty(t, token)
	assert.Error(t, err)
}

// This test mimics a situation where auth service returns a transient error
// and the client retries the request.
func TestAccessKeyTransientNetworkError(t *testing.T) {
	t.Parallel()
	count := atomic.Int32{}
	ctx := t.Context()
	testServer := &testAuthServer{
		tokenFunc: func() (string, error) {
			defer count.Add(1)
			time.Sleep(time.Millisecond * 100)
			if count.Load() == 0 {
				return "", connect.NewError(connect.CodeInternal, fmt.Errorf("something bad happened"))
			}
			return "test-token", nil
		},
	}
	svr := makeTestServer(testServer)
	defer svr.Close()

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		200*time.Millisecond,
	)

	token, err := getToken(context.TODO(), false)
	assert.Equal(t, "test-token", token)
	require.NoError(t, err)
	assert.Equal(t, int32(2), count.Load())
}

// Test that we don't retry forever if there is a non-transient network error.
func TestAccessKeyPermanentNetworkError(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	testServer := &testAuthServer{
		tokenFunc: func() (string, error) {
			return "", connect.NewError(connect.CodeInternal, fmt.Errorf("something bad happened"))
		},
	}
	svr := makeTestServer(testServer)
	defer svr.Close()

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		1*time.Millisecond,
	)

	token, err := getToken(context.TODO(), false)
	assert.Empty(t, token)
	assert.ErrorContains(t, err, "(Internal/Unknown) internal: something bad happened")
}

func TestAccessKeyForceOverridesExpiry(t *testing.T) {
	t.Parallel()
	count := atomic.Uint64{}
	ctx := t.Context()
	testServer := &testAuthServer{
		tokenFunc: func() (string, error) {
			token := strconv.Itoa(int(count.Load()))
			count.Add(1)
			return token, nil
		},
		// set a token with huge expiry so that we can be sure it will be cached for the duration of the test
		expiresInS: 500000,
	}
	svr := makeTestServer(testServer)
	defer svr.Close()

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		200*time.Millisecond,
	)

	// get the token once
	token, err := getToken(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "0", token)

	// get the token again. it should be the same
	token, err = getToken(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "0", token)

	// now get the token with force=true. it should be different
	token, err = getToken(ctx, true)
	require.NoError(t, err)
	assert.Equal(t, "1", token)
}

// Test that running a getToken(force=true) call doesn't block other getToken(force=false) calls.
func TestAccessKeyForceIsBlocking(t *testing.T) {
	t.Parallel()
	count := atomic.Uint64{}
	delaySecs := atomic.Uint64{}
	ctx := t.Context()
	testServer := &testAuthServer{
		tokenFunc: func() (string, error) {
			time.Sleep(time.Second * time.Duration(delaySecs.Load()))
			token := strconv.Itoa(int(count.Load()))
			count.Add(1)
			return token, nil
		},
		// set a token with huge expiry so that we can be sure it will be cached for the duration of the test
		expiresInS: 500000,
	}
	svr := makeTestServer(testServer)
	defer svr.Close()

	getToken := auth.AccessKeyAuth(
		ctx,
		"access-key",
		svr.URL,
		makeTransport(),
		200*time.Millisecond,
	)

	// get the token once so the cache is populated
	token, err := getToken(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "0", token)

	// now set the delay to 2 seconds and get the token with force=true
	// in the background.
	delaySecs.Store(2)
	go func() { _, _ = getToken(ctx, true) }()

	// wait for a little bit so we can be sure the goroutine is running
	time.Sleep(time.Second * 1)

	// now getting without force should block until the new token is returned
	token, err = getToken(ctx, false)
	require.NoError(t, err)
	assert.Equal(t, "1", token)
}

func TestAccessKeyNonRetryableErrorCodes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		code        connect.Code
		expectRetry bool
	}{
		{
			name:        "Unauthenticated - no retry",
			code:        connect.CodeUnauthenticated,
			expectRetry: false,
		},
		{
			name:        "PermissionDenied - no retry",
			code:        connect.CodePermissionDenied,
			expectRetry: false,
		},
		{
			name:        "NotFound - no retry",
			code:        connect.CodeNotFound,
			expectRetry: false,
		},
		{
			name:        "ResourceExhausted - retry",
			code:        connect.CodeResourceExhausted,
			expectRetry: true,
		},
		{
			name:        "Internal - retry",
			code:        connect.CodeInternal,
			expectRetry: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			count := atomic.Uint64{}
			ctx := t.Context()
			testServer := &testAuthServer{
				tokenFunc: func() (string, error) {
					count.Add(1)
					return "", connect.NewError(tt.code, fmt.Errorf("something bad happened"))
				},
			}
			svr := makeTestServer(testServer)
			defer svr.Close()

			getToken := auth.AccessKeyAuth(
				ctx,
				"access-key",
				svr.URL,
				makeTransport(),
				1*time.Millisecond,
			)

			_, err := getToken(context.TODO(), false)
			require.Error(t, err)
			if tt.expectRetry {
				assert.Greater(t, count.Load(), uint64(1))
			} else {
				assert.Equal(t, uint64(1), count.Load())
			}
		})
	}
}
