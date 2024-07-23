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
	"sync/atomic"
	"time"
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

type authBundle struct {
	accessToken string
	expires     time.Time
	refreshAt   time.Time
	err         error
}

func (a *authBundle) isExpired() bool {
	return time.Now().After(a.expires)
}

func (a *authBundle) shouldRefresh() bool {
	return time.Now().After(a.refreshAt)
}

func (a *authBundle) isValid() bool {
	return a != nil && a.accessToken != "" && !a.isExpired() && a.err == nil
}

type authTokenProvider struct {
	ctx context.Context

	clientID     string
	clientSecret string

	domain                  string
	audience                string
	grantType               string
	initialRetryBackoffTime time.Duration

	authResult     atomic.Pointer[authBundle]
	pendingRefresh atomic.Pointer[chan struct{}]
}

// Options is a struct of options to be passed to NewAuthTokenProvider.
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

// NewAuthTokenProvider creates a new AuthTokenProvider with the given context and options.
// If options is set to nil then the default options will be used.
//
// The supplied app context will be passed when performing background operations such as refreshing
// the access token, which ensures that no operation outlives the lifetime of the application.
//
// By default the AuthTokenProvider will fetch the client ID and client secret from the environment variables
// `STATELY_CLIENT_ID` and `STATELY_CLIENT_SECRET`, however these can be explicitly overridden by passing
// credentials in the options. If no credentials are found, NewAuthTokenProvider will return an error.
//
//nolint:revive // can't use client.AuthTokenProvider here because of circular dependency
func NewAuthTokenProvider(
	appCtx context.Context,
	clientID, clientSecret string,
	options *Options,
) *authTokenProvider {
	options = applyDefaults(options)
	return &authTokenProvider{
		ctx:                     appCtx,
		clientID:                clientID,
		clientSecret:            clientSecret,
		domain:                  options.Domain,
		audience:                options.Audience,
		grantType:               defaultGrantType,
		initialRetryBackoffTime: options.InitialRetryBackoffTime,
		authResult:              atomic.Pointer[authBundle]{},
		pendingRefresh:          atomic.Pointer[chan struct{}]{},
	}
}

func (p *authTokenProvider) GetAccessToken(ctx context.Context) (string, error) {
	currentAuth := p.authResult.Load()

	if currentAuth.isValid() {
		if currentAuth.shouldRefresh() {
			// If we're due for a refresh then trigger it in the background.
			// err is always nil if background == true so we can throw it away.
			bgCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_, _ = p.refreshAccessToken(bgCtx, true)
		}
		return currentAuth.accessToken, nil
	}

	// If the current auth is not valid then we need to do a blocking refresh.
	return p.refreshAccessToken(ctx, false)
}

func (p *authTokenProvider) InvalidateAccessToken() {
	p.authResult.Store(nil)
}

// refreshAccessToken makes a request to the auth0 server to refresh the access token if one is not already
// pending or returns the result of the pending refresh if one is already in progress.
// If `background` is set to true then the function will return with a nil error and empty token (""), otherwise
// it will wait for the result of the network request.
// If `background` is false and the pending network request fails, then the calling goroutine will return the network error
// while all other waiting goroutines will return a generic error and empty token.
func (p *authTokenProvider) refreshAccessToken(ctx context.Context, background bool) (string, error) {
	// Set up a channel for when we actually execute the request.
	// It's possible that this will be unused if CompareAndSwap fails.
	waiter := make(chan struct{})

	// Try to win the "right" to refresh by swapping our channel into the
	// pendingRefresh slot. If the swap succeeds then we are responsible for
	// making the network request.
	winner := p.pendingRefresh.CompareAndSwap(nil, &waiter)
	if winner {
		fn := func() error {
			defer close(waiter) // signals waiters to unblock
			defer p.pendingRefresh.Store(nil)
			initialTokens := p.authResult.Load()
			tokens, err := p.refreshAccessTokenReq(ctx)
			if err == nil {
				p.authResult.CompareAndSwap(initialTokens, tokens)
			} else if !initialTokens.isValid() {
				p.authResult.CompareAndSwap(initialTokens, &authBundle{err: err})
			}
			return err
		}
		if background {
			//nolint:errcheck // run in the background, we don't care about the error
			// since we'll return immediately.
			go fn()
		} else {
			// run in the current goroutine
			if err := fn(); err != nil {
				return "", err
			}
		}
	}

	// If this is a background refresh we can return here, regardless of whether
	// we won the right to refresh.
	if background {
		return "", nil
	}

	if !winner {
		// If we didn't do the refresh ourselves (another goroutine did it) then we
		// need to wait for that other goroutine to finish.
		currentWaiter := p.pendingRefresh.Load()
		// It's possible the other goroutine has already completed by the time we
		// get here.
		if currentWaiter != nil {
			select {
			case <-*currentWaiter: // reading from a closed channel always returns immediately
				// success
			case <-ctx.Done():
				// return early if the context is cancelled
				return "", ctx.Err()
			}
		}
	}

	currentAuth := p.authResult.Load()
	if !currentAuth.isValid() {
		if currentAuth.err != nil {
			return "", currentAuth.err
		}
		return "", errors.New("refresh request failed - unable to find valid token")
	}
	return currentAuth.accessToken, nil
}

// responseToErr converts a non-200 HTTP response into an error
// by reading the text from the response. If the response is a 200
// then the func returns nil.
func responseToErr(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	defer resp.Body.Close()
	txt, reqErr := io.ReadAll(resp.Body)
	if reqErr != nil {
		return fmt.Errorf(
			"Auth0 returned %d. Failed to decode response with error: %s",
			resp.StatusCode,
			reqErr.Error(),
		)
	}

	return fmt.Errorf("Auth0 returned %d. Response body: %s", resp.StatusCode, txt)
}

// refreshAccessTokenReq makes the network request for a new token and returns the token,
// along with a time when the token should be refreshed or an error if the request failed.
func (p *authTokenProvider) refreshAccessTokenReq(ctx context.Context) (*authBundle, error) {
	params := &authRequest{
		ClientID:     p.clientID,
		ClientSecret: p.clientSecret,
		Audience:     p.audience,
		GrantType:    p.grantType,
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		p.domain+"/oauth/token",
		bytes.NewBuffer(jsonParams),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	var resp *http.Response
	var reqErr error
	for attempt := 0; attempt < 3; attempt++ {
		resp, reqErr = http.DefaultClient.Do(req)
		// no point in retrying if the context is cancelled
		if errors.Is(reqErr, context.Canceled) || errors.Is(reqErr, context.DeadlineExceeded) {
			break
		}

		if reqErr == nil {
			// The error may be encoded in the response status
			reqErr = responseToErr(resp)
		}

		if reqErr != nil {
			err := sleepWithContext(ctx, backoff(attempt, p.initialRetryBackoffTime))
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
	authResp := &authResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return nil, err
	}

	// Calculate a random multiplier between 0.3 and 0.8 to to apply to the expiry
	// so that we refresh in the background ahead of expiration, but avoid
	// multiple processes hammering the service at the same time.
	jitter := (rand.Float64() * 0.5) + 0.3
	expires := time.Now().Add(time.Second * time.Duration(authResp.ExpiresInSecs))
	refreshAt := time.Now().Add(time.Second * time.Duration(float64(authResp.ExpiresInSecs)*jitter))
	// return the token
	return &authBundle{
		accessToken: authResp.AccessToken,
		expires:     expires,
		refreshAt:   refreshAt,
	}, err
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
