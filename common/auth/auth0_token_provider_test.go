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

	"github.com/StatelyCloud/go-sdk/common/auth"
	"github.com/StatelyCloud/stately/gocommon/stypes"
)

func TestGetToken(t *testing.T) {
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
	p := auth.NewAuth0TokenProvider(
		context.TODO(),
		"client-id",
		"client-secret",
		auth.WithAudience("test-aud"), auth.WithDomain(svr.URL),
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
	count := atomic.Uint64{}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make the handler take a long time so that the requests back up
		time.Sleep(time.Millisecond * 100)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).
			Encode(map[string]interface{}{"access_token": strconv.Itoa(int(count.Load())), "expires_in": 500000})
		require.NoError(t, err)
		count.Add(1)
	}))

	p := auth.NewAuth0TokenProvider(context.TODO(), "", "", auth.WithDomain(svr.URL))
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
	returnVal := stypes.AtomicValueOf("test-token")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// send a response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).
			Encode(map[string]interface{}{"access_token": returnVal.Load(), "expires_in": 0})
		require.NoError(t, err)
	}))
	defer svr.Close()

	// call GetAccessToken()
	p := auth.NewAuth0TokenProvider(context.TODO(), "", "", auth.WithDomain(svr.URL))
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
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	p := auth.NewAuth0TokenProvider(ctx, "", "", auth.WithDomain(svr.URL))
	token, err := p.GetAccessToken(ctx)
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
}
