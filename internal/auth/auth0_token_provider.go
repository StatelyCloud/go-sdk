package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"slices"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

type authRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type authResponse struct {
	AccessToken   string `json:"access_token"`
	ExpiresInSecs uint64 `json:"expires_in"`
}

const (
	// default vals.
	defaultDomain                  = "https://oauth.stately.cloud"
	defaultAudience                = "api.stately.cloud"
	defaultGrantType               = "client_credentials"
	defaultInitialRetryBackoffTime = 200 * time.Millisecond
)

// Options is a struct of options to be passed to InitServerAuth.
// This can be omitted to use the default options or can be passed explicitly with
// any overrides.
type Options struct {
	// Domain is the domain to query for auth tokens.
	// Defaults to https://oauth.stately.cloud
	Domain string
	// Audience is the audience that the provider will request tokens for.
	// Defaults to api.stately.cloud
	Audience string
	// InitialRetryBackoffTime is the initial time to wait before retrying a failed auth request.
	// This will be exponentially backed off with jitter on successive failures.
	// This must be more than 0 otherwise the default will be used.
	// Defaults to 200 milliseconds
	InitialRetryBackoffTime time.Duration
}

// applyDefaults iterates through the given options struct and applied default values
// where required.
func applyDefaults(options *Options) *Options {
	if options == nil {
		options = &Options{}
	}

	// Domain
	if options.Domain == "" {
		options.Domain = defaultDomain
	}

	// Audience
	if options.Audience == "" {
		options.Audience = defaultAudience
	}

	// InitialRetryBackoffTime
	if options.InitialRetryBackoffTime <= 0 {
		options.InitialRetryBackoffTime = defaultInitialRetryBackoffTime
	}

	return options
}

// getTokenState is a wrapper for the atomic state of the access token.
type getTokenState struct {
	accessToken   string
	expiresAtSecs int64
}

// InitServerAuth creates a new AuthTokenProvider with the given context and options.
// If options is set to nil then the default options will be used.
//
// The supplied app context will be passed when performing background operations such as refreshing
// the access token, which ensures that no operation outlives the lifetime of the application.
//
// By default the AuthTokenProvider will fetch the client ID and client secret from the environment variables
// `STATELY_CLIENT_ID` and `STATELY_CLIENT_SECRET`, however these can be explicitly overridden by passing
// credentials in the options. If no credentials are found, InitServerAuth will return an error.
func InitServerAuth(
	appCtx context.Context,
	clientID, clientSecret string,
	options *Options,
) func(ctx context.Context, force bool) (string, error) {
	options = applyDefaults(options)
	var state atomic.Value
	state.Store(getTokenState{
		accessToken:   "",
		expiresAtSecs: 0,
	})
	validAccessToken := func() (string, bool) {
		currentState := state.Load().(getTokenState)
		if currentState.accessToken != "" && time.Now().Unix() < currentState.expiresAtSecs {
			return currentState.accessToken, true
		}
		return "", false
	}

	var refreshGroup singleflight.Group
	var refresh func(ctx context.Context) (string, error)
	refreshImpl := func(ctx context.Context) (string, error) {
		// refreshImpl is guaranteed to only be running once at a time
		// so we don't need to worry about currentState being modified during
		// execution.

		currentState := state.Load().(getTokenState)

		// if the context has been killed then return
		if err := appCtx.Err(); err != nil {
			return "", err
		}

		resp, err := makeAuth0Req(ctx, clientID, clientSecret, options)
		if err != nil {
			// if there is an error making a network request them propagate it to the caller
			// and leave the state intact so that if the auth isn't expired it can continue to be used
			return "", err
		}

		// read the response data
		newExpiresInSecs := int64(resp.ExpiresInSecs)
		newExpiresAtSecs := time.Now().Unix() + newExpiresInSecs
		newAccessToken := resp.AccessToken

		// if the new token has expiry time greater than the current token then update the state
		if newExpiresAtSecs > currentState.expiresAtSecs {
			state.Store(getTokenState{
				accessToken:   newAccessToken,
				expiresAtSecs: newExpiresAtSecs,
			})
		} else {
			// otherwise overwrite the new values with the old ones
			// since they will last longer
			newExpiresAtSecs = currentState.expiresAtSecs
			newAccessToken = currentState.accessToken
		}

		// schedule the next refresh using currentExpiresAtSecs
		// if the token was not updated we still use the old expiry time
		refreshIn := time.Until(time.Unix(newExpiresAtSecs, 0))
		jitter := (rand.Float64() * 0.05) + 0.9
		timer := time.NewTimer(time.Duration(float64(refreshIn) * jitter))

		go func() {
			select {
			case <-appCtx.Done():
				if !timer.Stop() {
					<-timer.C
				}
				return
			case <-timer.C:
				_, _ = refresh(appCtx)
			}
		}()

		return newAccessToken, nil
	}

	refresh = func(ctx context.Context) (string, error) {
		res, err, _ := refreshGroup.Do("auth0", func() (any, error) {
			return refreshImpl(ctx)
		})
		return res.(string), err
	}

	getToken := func(ctx context.Context, force bool) (string, error) {
		if force {
			state.Store(getTokenState{
				accessToken:   "",
				expiresAtSecs: 0,
			})
		} else if token, ok := validAccessToken(); ok {
			return token, nil
		}
		return refresh(ctx)
	}

	// kick off the first refresh immediately
	go func() { _, _ = getToken(appCtx, false) }()
	return getToken
}

// makes the auth request to auth0. This will retry internally if the request fails in a transient way.
func makeAuth0Req(ctx context.Context, clientID, clientSecret string, options *Options) (*authResponse, error) {
	params := &authRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Audience:     options.Audience,
		GrantType:    defaultGrantType,
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	var reqErr error
	var retryable bool
	for attempt := 0; attempt < 10; attempt++ {
		// need to create the req inside the loop to avoid
		// errors caused by reading the same body twice
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			options.Domain+"/oauth/token",
			bytes.NewBuffer(jsonParams),
		)
		if err != nil {
			return nil, err
		}
		req.Header.Add("content-type", "application/json")
		//nolint:bodyclose // we are closing the body. the linter just can't work it out
		resp, retryable, reqErr = parseErrorResponse(http.DefaultClient.Do(req))

		if reqErr != nil {
			if !retryable {
				return nil, reqErr
			}
			err := sleepWithContext(ctx, backoff(attempt, options.InitialRetryBackoffTime))
			if err != nil {
				return nil, err
			}
			continue
		}
		break
	}

	// read the error the retry loop finished with
	if reqErr != nil {
		return nil, reqErr
	}

	// decode the response
	defer resp.Body.Close()
	authResp := &authResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return nil, err
	}
	return authResp, nil
}

var nonRetryableErrors = []int{
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusNotFound,
	http.StatusMethodNotAllowed,
	http.StatusNotAcceptable,
	http.StatusProxyAuthRequired,
	http.StatusBadRequest,
}

// parseErrorResponse takes the output from a http.Do() request and
// converts a non-200 HTTP response into an error by reading the text from the response body.
// If Do() returns an error then that error is returned direcly.
// If the response is a 200 then the response is returned with a nil error.
// If an error is returned then a boolean is returned to indicate if the error is retryable.
func parseErrorResponse(resp *http.Response, err error) (*http.Response, bool, error) {
	// if we got an error then return that error directly
	if err != nil {
		// don't bother retrying context errors
		retryable := !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded)
		return nil, retryable, err
	}

	// if the response is a 200 then pass through the response
	if resp.StatusCode == http.StatusOK {
		return resp, false, nil
	}

	// after this point we know we have an error and we are reading the body here so we have to close it
	defer resp.Body.Close()

	// otherwise we actually need to read the body
	retryable := !slices.Contains(nonRetryableErrors, resp.StatusCode)

	txt, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return resp, retryable, fmt.Errorf(
			"Auth0 returned %d. Failed to decode response with error: %s",
			resp.StatusCode,
			readErr.Error(),
		)
	}

	return resp, retryable, fmt.Errorf("Auth0 returned %d. Response body: %s", resp.StatusCode, txt)
}

// backoff returns a duration to wait before retrying a request. `attempt` is
// the current attempt number, starting from 0 (e.g. the first attempt is 0,
// then 1, then 2...).
func backoff(attempt int, baseBackoff time.Duration) time.Duration {
	// Double the base backoff time per attempt, starting with 1
	exp := 1 << attempt // 2^attempt
	// Add a full jitter to the backoff time, from no wait to 100% of the exponential backoff.
	// See https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
	jitter := rand.Float64()
	return time.Duration(float64(exp)*jitter) * baseBackoff
}

// TODO: When adding jitter to scheduled work, we do not select the jitter on
// each host randomly. Instead, we use a consistent method that produces the
// same number every time on the same host. This way, if there is a service
// being overloaded, or a race condition, it happens the same way in a pattern.
// We humans are good at identifying patterns, and we're more likely to
// determine the root cause.
// https://aws.amazon.com/builders-library/timeouts-retries-and-backoff-with-jitter/
