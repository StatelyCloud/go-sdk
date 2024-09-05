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
		err = json.NewEncoder(w).Encode(map[string]any{"access_token": "test-token", "expires_in": 1000})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	p := auth.NewAuthTokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Audience: "test-aud",
			Domain:   svr.URL,
		},
	)

	token, err := p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// now close the server and call GetToken again. it should work without a network request because the value is cached
	svr.Close()
	token, err = p.GetAccessToken(context.TODO())
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
		err := json.NewEncoder(w).
			Encode(map[string]any{"access_token": strconv.Itoa(int(count.Load())), "expires_in": 500000})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		count.Add(1)
	}))

	p := auth.NewAuthTokenProvider(
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
			token, err := p.GetAccessToken(context.TODO())
			assert.NoError(t, err)
			assert.Equal(t, "0", token)
		}()
	}

	wg.Wait()
}

func TestGetExpiredAuth(t *testing.T) {
	t.Parallel()
	returnVal := atomic.Value{}
	returnVal.Store("test-token")

	// Mock out a token that expires in 1 second
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// send a response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).
			Encode(map[string]any{"access_token": returnVal.Load(), "expires_in": 1})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	defer svr.Close()

	p := auth.NewAuthTokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	token, err := p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// Let the token expire
	time.Sleep(1 * time.Second)

	// now update the return value and call GetAccessToken again
	// we will trigger a blocking refresh because the current token is expired
	returnVal.Store("test-token2")
	// triggers the background refresh
	token, err = p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	// We expect the new token to be returned because the old one had expired
	assert.Equal(t, "test-token2", token)
}

// NOTE: This test has 2 sleeps in it. I can see a world where it is flaky because of that.
// If you see this test failing in the CI then maybe just comment it out.
func TestBackgroundRefresh(t *testing.T) {
	t.Parallel()
	returnVal := atomic.Value{}
	returnVal.Store("test-token")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).
			Encode(map[string]any{"access_token": returnVal.Load(), "expires_in": 1})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	defer svr.Close()

	p := auth.NewAuthTokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	token, err := p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// now update the return value and call GetAccessToken again
	// we will trigger a blocking refresh because the current token is expired
	returnVal.Store("test-token2")
	// sleep 850 millis which is guaranteed to put us in the background refresh window
	time.Sleep(time.Millisecond * 850)
	// triggers the background refresh
	token, err = p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// sleep another 200 millis to allow for the request to complete
	time.Sleep(time.Millisecond * 200)
	token, err = p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	assert.Equal(t, "test-token2", token)
}

func TestRefreshContextCancelled(t *testing.T) {
	t.Parallel()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).
			Encode(map[string]any{"access_token": "test-token", "expires_in": 0})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	defer svr.Close()

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	p := auth.NewAuthTokenProvider(
		ctx,
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)

	token, err := p.GetAccessToken(ctx)
	assert.Equal(t, "", token)
	assert.Error(t, err)
}

// This test mimics a situation where auth0 is down and returning 500 errors.
// The test asserts that calls to GetAccessToken() will return an error and an empty token.
func TestNetworkError(t *testing.T) {
	t.Parallel()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(time.Millisecond * 100)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("500 - Something bad happened!"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	defer svr.Close()

	p := auth.NewAuthTokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain:                  svr.URL,
			InitialRetryBackoffTime: time.Microsecond * 1,
		},
	)

	token, err := p.GetAccessToken(context.TODO())
	assert.Equal(t, "", token)
	assert.Error(t, err)
}
