package auth

import (
	"context"
	"errors"
	"math/rand/v2"
	"net/http"
	"sync/atomic"
	"time"

	"connectrpc.com/connect"
	"github.com/planetscale/vtprotobuf/codec/grpc"
	"golang.org/x/net/http2"
	"golang.org/x/sync/singleflight"

	"github.com/StatelyCloud/go-sdk/pb/auth"
	"github.com/StatelyCloud/go-sdk/pb/auth/authconnect"
	"github.com/StatelyCloud/go-sdk/sdkerror"
)

const (
	maxRetries = 10
)

type accessKeyAuth struct {
	client           authconnect.AuthServiceClient
	accessKey        string
	state            atomic.Value
	retryBackoffTime time.Duration
	refreshGroup     singleflight.Group
	appCtx           context.Context
}

func (a *accessKeyAuth) validAccessToken() (string, bool) {
	currentState := a.state.Load().(getTokenState)
	if currentState.accessToken != "" && time.Now().Unix() < currentState.expiresAtSecs {
		return currentState.accessToken, true
	}
	return "", false
}

// refresh makes sure only one call to doRefresh is in flight at once, but all
// callers will get the same answer.
func (a *accessKeyAuth) refresh(ctx context.Context) (string, error) {
	res, err, _ := a.refreshGroup.Do("accessKey", func() (any, error) {
		// this is guaranteed to only be running once at a time
		// so we don't need to worry about currentState being modified during
		// execution.

		currentState := a.state.Load().(getTokenState)

		// if the app lifecycle context has been killed then return
		if err := a.appCtx.Err(); err != nil {
			return "", err
		}

		resp, err := a.fetchNewAuthToken(ctx)
		if err != nil {
			// if there is an error making a network request them propagate it to the caller
			// and leave the state intact so that if the auth isn't expired it can continue to be used
			return "", err
		}

		// read the response data
		newExpiresInSecs := int64(resp.ExpiresInS)
		newExpiresAtSecs := time.Now().Unix() + newExpiresInSecs
		newAuthToken := resp.AuthToken

		// if the new token has expiry time greater than the current token then update the state
		if newExpiresAtSecs > currentState.expiresAtSecs {
			a.state.Store(getTokenState{
				accessToken:   newAuthToken,
				expiresAtSecs: newExpiresAtSecs,
			})
		} else {
			// otherwise overwrite the new values with the old ones
			// since they will last longer
			newExpiresAtSecs = currentState.expiresAtSecs
			newAuthToken = currentState.accessToken
		}

		// schedule the next refresh using currentExpiresAtSecs
		// if the token was not updated we still use the old expiry time
		refreshIn := time.Until(time.Unix(newExpiresAtSecs, 0))
		jitter := (rand.Float64() * 0.05) + 0.9
		go func() {
			err := sleepWithContext(a.appCtx, time.Duration(float64(refreshIn)*jitter))
			if err != nil {
				return
			}
			//nolint:errcheck // we don't have a way to communicate errors here
			a.refresh(a.appCtx)
		}()

		return newAuthToken, nil
	})
	return res.(string), err
}

// GetToken returns the current access token. If force is true, it will
// invalidate the token (so no other request can use it either) before fetching
// a new one.
func (a *accessKeyAuth) GetToken(ctx context.Context, force bool) (string, error) {
	if force {
		a.state.Store(getTokenState{
			accessToken:   "",
			expiresAtSecs: 0,
		})
	} else if token, ok := a.validAccessToken(); ok {
		return token, nil
	}
	return a.refresh(ctx)
}

// AccessKeyAuth creates a new AuthTokenProvider that fetches auth tokens using an access key.
// If options is set to nil then the default options will be used.
//
// The supplied app context will be passed when performing background operations such as refreshing
// the access token, which ensures that no operation outlives the lifetime of the application.
//
// By default the AuthTokenProvider will fetch the access key from the environment variable
// `STATELY_ACCESS_KEY` however this can be explicitly overridden by passing
// credentials in the options. If no credentials are found, this will return an error.
func AccessKeyAuth(
	appCtx context.Context,
	accessKey string,
	endpoint string,
	transport *http2.Transport,
	retryBackoffTime time.Duration,
) func(ctx context.Context, force bool) (string, error) {
	a := &accessKeyAuth{
		client:           createAuthServiceClient(endpoint, transport),
		accessKey:        accessKey,
		retryBackoffTime: retryBackoffTime,
		appCtx:           appCtx,
	}
	a.state.Store(getTokenState{
		accessToken:   "",
		expiresAtSecs: 0,
	})

	// kick off the first refresh immediately
	//nolint:errcheck // we don't have a way to communicate errors here
	go a.GetToken(appCtx, false)

	return a.GetToken
}

// createAuthServiceClient creates a new connect client to talk to the auth service.
func createAuthServiceClient(endpoint string, transport *http2.Transport) authconnect.AuthServiceClient {
	httpClient := &http.Client{
		Transport: transport,
	}

	client := authconnect.NewAuthServiceClient(
		httpClient,
		endpoint,
		connect.WithCodec(grpc.Codec{}), // enable vtprotobuf codec
		connect.WithInterceptors(sdkerror.ConnectErrorInterceptor()),
	)
	return client
}

// makes the request to GetAuthToken. This will retry internally if the request fails in a transient way.
func (a *accessKeyAuth) fetchNewAuthToken(ctx context.Context) (*auth.GetAuthTokenResponse, error) {
	for attempt := 0; ; attempt++ {
		resp, err := a.client.GetAuthToken(ctx, connect.NewRequest(&auth.GetAuthTokenRequest{
			Identity: &auth.GetAuthTokenRequest_AccessKey{
				AccessKey: a.accessKey,
			},
		}))
		if err != nil {
			serr := &sdkerror.Error{}
			if errors.As(err, &serr) {
				retriable := serr.Code != connect.CodeUnauthenticated &&
					serr.Code != connect.CodePermissionDenied &&
					serr.Code != connect.CodeNotFound &&
					serr.Code != connect.CodeInvalidArgument
				if retriable && attempt < maxRetries {
					err := sleepWithContext(ctx, backoff(attempt, a.retryBackoffTime))
					if err != nil {
						return nil, err
					}
					continue
				}
			}
			return nil, err
		}
		return resp.Msg, nil
	}
}
