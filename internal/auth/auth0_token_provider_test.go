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

	"github.com/StatelyCloud/go-sdk/internal/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetToken(t *testing.T) {
	t.Parallel()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read the body into a map of interfaces
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		reqData := map[string]string{}
		err = json.Unmarshal(body, &reqData)
		require.NoError(t, err)

		// check that the audience was set correctly
		assert.Equal(t, "test-aud", reqData["audience"])
		assert.Equal(t, "client-id", reqData["client_id"])
		assert.Equal(t, "client-secret", reqData["client_secret"])

		// send a response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(map[string]interface{}{"access_token": "test-token", "expires_in": 1000})
		require.NoError(t, err)
	}))

	// call GetAccessToken()
	p, err := auth.NewAuthTokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Audience: "test-aud",
			Domain:   svr.URL,
		},
	)
	require.NoError(t, err)

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
			Encode(map[string]interface{}{"access_token": strconv.Itoa(int(count.Load())), "expires_in": 500000})
		require.NoError(t, err)
		count.Add(1)
	}))

	p, err := auth.NewAuthTokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)
	require.NoError(t, err)

	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			token, err := p.GetAccessToken(context.TODO())
			require.NoError(t, err)
			assert.Equal(t, "0", token)
		}()
	}

	wg.Wait()
}

func TestRefreshExpiryScheduler(t *testing.T) {
	t.Parallel()
	returnVal := atomic.Value{}
	returnVal.Store("test-token")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// send a response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).
			Encode(map[string]interface{}{"access_token": returnVal.Load(), "expires_in": 0})
		require.NoError(t, err)
	}))
	defer svr.Close()

	// call GetAccessToken()
	p, err := auth.NewAuthTokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)
	require.NoError(t, err)

	token, err := p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)

	// now update the return value and call GetAccessToken again
	// we should refresh the new value from the server
	returnVal.Store("test-token2")
	// wait 20ms to give some buffer for th 0ms refresh
	time.Sleep(time.Millisecond * 20)
	token, err = p.GetAccessToken(context.TODO())
	require.NoError(t, err)
	assert.Equal(t, "test-token2", token)
}

func TestRefreshContextCancelled(t *testing.T) {
	t.Parallel()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// send a response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).
			Encode(map[string]interface{}{"access_token": "test-token", "expires_in": 0})
		require.NoError(t, err)
	}))
	defer svr.Close()

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	// call GetAccessToken()
	p, err := auth.NewAuthTokenProvider(
		ctx,
		"client-id",
		"client-secret",
		&auth.Options{
			Domain: svr.URL,
		},
	)
	require.NoError(t, err)

	token, err := p.GetAccessToken(ctx)
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
}
