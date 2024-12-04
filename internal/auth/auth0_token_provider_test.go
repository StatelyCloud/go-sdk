package auth_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/StatelyCloud/go-sdk/internal/auth"
)

func TestGetToken(t *testing.T) {
	t.Parallel()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read the body into a map of interfaces
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reqData := map[string]string{}
		err = json.Unmarshal(body, &reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// check that the audience was set correctly
		assert.Equal(t, "test-aud", reqData["audience"])
		assert.Equal(t, "client-id", reqData["client_id"])
		assert.Equal(t, "client-secret", reqData["client_secret"])

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "test-token", "expires_in": 1000})
	}))

	getToken := auth.InitServerAuth(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Audience: "test-aud",
			Domain:   svr.URL,
		},
	)

	token, err := getToken(context.TODO(), false)
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// now close the server and call GetToken again. it should work without a network request because the value is cached
	svr.Close()
	token, err = getToken(context.TODO(), false)
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)
}

func TestConcurrentRefresh(t *testing.T) {
	t.Parallel()
	count := atomic.Uint64{}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// make the handler take a long time so that the requests back up
		time.Sleep(time.Millisecond * 100)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).
			Encode(map[string]any{"access_token": strconv.Itoa(int(count.Load())), "expires_in": 500000})

		count.Add(1)
	}))

	getToken := auth.InitServerAuth(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
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

// Test that the background refresh is scheduled as expected.
func TestBackgroundRefreshScheduler(t *testing.T) {
	t.Parallel()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "test-token", "expires_in": 1})

		wg.Done()
	}))
	defer svr.Close()

	getToken := auth.InitServerAuth(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	token, err := getToken(context.TODO(), false)
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// we won't be able to continue past this point until the automatic refresh happens
	wg.Wait()
}

func TestRefreshContextCancelled(t *testing.T) {
	t.Parallel()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).
			Encode(map[string]any{"access_token": "test-token", "expires_in": 0})
	}))
	defer svr.Close()

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	getToken := auth.InitServerAuth(
		ctx,
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	token, err := getToken(ctx, false)
	assert.Equal(t, "", token)
	assert.Error(t, err)
}

// This test mimics a situation where auth0 returns a transient error
// and the client retries the request.
func TestTransientNetworkError(t *testing.T) {
	t.Parallel()
	count := atomic.Int32{}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		defer count.Add(1)
		time.Sleep(time.Millisecond * 100)
		if count.Load() == 0 {

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("500 - Something bad happened!"))

		} else {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).
				Encode(map[string]any{"access_token": "test-token", "expires_in": 1000})

		}
	}))
	defer svr.Close()

	getToken := auth.InitServerAuth(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	token, err := getToken(context.TODO(), false)
	assert.Equal(t, "test-token", token)
	require.NoError(t, err)
	assert.Equal(t, int32(2), count.Load())
}

// Test that we don't retry forever if there is a non-transient network error.
func TestPermanentNetworkError(t *testing.T) {
	t.Parallel()

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - Something bad happened!"))
	}))
	defer svr.Close()

	getToken := auth.InitServerAuth(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain:                  svr.URL,
			InitialRetryBackoffTime: time.Millisecond * 1,
		},
	)

	token, err := getToken(context.TODO(), false)
	assert.Empty(t, token)
	require.Equal(t, "Auth0 returned 500. Response body: 500 - Something bad happened!", err.Error())
}

func TestForceOverridesExpiry(t *testing.T) {
	t.Parallel()
	count := atomic.Uint64{}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// set a token with huge expiry so that we can be sure it will be cached for the duration of the test
		_ = json.NewEncoder(w).
			Encode(map[string]any{"access_token": strconv.Itoa(int(count.Load())), "expires_in": 500000})
		count.Add(1)
	}))

	getToken := auth.InitServerAuth(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	// get the token once
	token, err := getToken(context.TODO(), false)
	require.NoError(t, err)
	assert.Equal(t, "0", token)

	// get the token again. it should be the same
	token, err = getToken(context.TODO(), false)
	require.NoError(t, err)
	assert.Equal(t, "0", token)

	// now get the token with force=true. it should be different
	token, err = getToken(context.TODO(), true)
	require.NoError(t, err)
	assert.Equal(t, "1", token)
}

// Test that running a getToken(force=true) call blocks other getToken(force=false) calls.
func TestForceIsBlocking(t *testing.T) {
	t.Parallel()
	count := atomic.Uint64{}
	delaySecs := atomic.Uint64{}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		time.Sleep(time.Second * time.Duration(delaySecs.Load()))
		_ = json.NewEncoder(w).
			Encode(map[string]any{"access_token": strconv.Itoa(int(count.Load())), "expires_in": 500000})
		count.Add(1)
	}))

	getToken := auth.InitServerAuth(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	// get the token once so the cache is populated
	token, err := getToken(context.TODO(), false)
	require.NoError(t, err)
	assert.Equal(t, "0", token)

	// now set the delay to 2 seconds and get the token with force=true
	// in the background.
	delaySecs.Store(2)
	go func() { _, _ = getToken(context.TODO(), true) }()

	// wait for a little bit so we can be sure the goroutine is running
	time.Sleep(time.Second * 1)

	// now getting without force should block until the new token is returned
	token, err = getToken(context.TODO(), false)
	require.NoError(t, err)
	assert.Equal(t, "1", token)
}

func TestNonRetryableErrorCodes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		statusCode  int
		expectRetry bool
	}{
		{
			name:        "401 - no retry",
			statusCode:  http.StatusUnauthorized,
			expectRetry: false,
		},
		{
			name:        "403 - no retry",
			statusCode:  http.StatusForbidden,
			expectRetry: false,
		},
		{
			name:        "404 - no retry",
			statusCode:  http.StatusNotFound,
			expectRetry: false,
		},
		{
			name:        "429 - retry",
			statusCode:  http.StatusTooManyRequests,
			expectRetry: true,
		},
		{
			name:        "500 - retry",
			statusCode:  http.StatusInternalServerError,
			expectRetry: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			count := atomic.Uint64{}
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				count.Add(1)
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte("done"))
			}))
			defer svr.Close()

			getToken := auth.InitServerAuth(
				context.TODO(),
				"client-id",
				"client-secret",
				&auth.Options{
					Domain:                  svr.URL,
					InitialRetryBackoffTime: time.Millisecond * 1,
				},
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
