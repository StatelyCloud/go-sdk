package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	timeComm "github.com/StatelyCloud/go-sdk/common/time"
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

// TokenProvider is the interface which must be passed into the WithOAuthRefreshUnaryInterceptor
// so that it can authenticate outgoing requests.
// This is a threadsafe interface.
type TokenProvider interface {
	// GetAccessToken returns an access token or an error.
	// If there is no current access token then the provider will attempt to refresh
	// and get a new access token.
	// An error is returned if there is no access token and the refresh operation fails
	GetAccessToken(ctx context.Context) (string, error)
	// RefreshAccessToken refreshes the current access token.
	// If there is already an existing access token that token will be returned unless
	// force=true is passed.
	// If the refresh network operation fails then an error will be returned
	// On successful refresh the internal state of the provider will be updated with the new token
	// and successive calls to GetAccessToken() will return the new token.
	RefreshAccessToken(ctx context.Context, force bool) (string, error)
}

const (
	// stagingAuthDomain = "https://oauth-dev.stately.cloud"
	defaultDomain    = "https://oauth.stately.cloud"
	defaultAudience  = "api.stately.cloud"
	defaultGrantType = "client_credentials"
)

type auth0TokenProvider struct {
	appCtx       context.Context
	clientID     string
	clientSecret string
	domain       string
	audience     string
	grantType    string
	accessToken  string
	mutex        *sync.RWMutex
}

// Auth0TokenProviderOpt is an option to be passed to NewAuth0TokenProvider.
type Auth0TokenProviderOpt = func(*auth0TokenProvider)

// WithDomain creates an option to override the auth domain that the
// Auth0TokenProvider fetches auth from.
func WithDomain(domain string) Auth0TokenProviderOpt {
	return func(a *auth0TokenProvider) {
		a.domain = domain
	}
}

// WithAudience creates an option to override the audience that the
// Auth0TokenProvider requests.
func WithAudience(audience string) Auth0TokenProviderOpt {
	return func(a *auth0TokenProvider) {
		a.audience = audience
	}
}

// NewAuth0TokenProvider creates a new AuthTokenProvider which vends token
// from auth0 using the given client credentials.
func NewAuth0TokenProvider(
	appCtx context.Context,
	clientID, clientSecret string,
	options ...Auth0TokenProviderOpt,
) TokenProvider {
	p := &auth0TokenProvider{
		appCtx:       appCtx,
		clientID:     clientID,
		clientSecret: clientSecret,
		domain:       defaultDomain,
		audience:     defaultAudience,
		grantType:    defaultGrantType,
		accessToken:  "",
		mutex:        &sync.RWMutex{},
	}

	for _, opt := range options {
		opt(p)
	}

	// refresh access token as soon as we create this thing so the first request is faster
	go func() {
		_, _ = p.RefreshAccessToken(appCtx, false)
	}()
	return p
}

func (p *auth0TokenProvider) GetAccessToken(ctx context.Context) (string, error) {
	// return the current token
	p.mutex.RLock()

	if p.accessToken != "" {
		defer p.mutex.RUnlock()
		return p.accessToken, nil
	}

	// if theres no access token then we need to do a refetch
	// so unlock the rlock before we try to take a full lock in RefreshAccessToken()
	p.mutex.RUnlock()
	return p.RefreshAccessToken(ctx, false)
}

func (p *auth0TokenProvider) RefreshAccessToken(ctx context.Context, force bool) (string, error) {
	// take a full lock. there is only one caller in this function at once so its
	// totally safe to update the state
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// if someone beat us to the lock and already did a refresh
	// then simply return the value. We know that happened because accessToken is not empty
	//
	// if we are forcing an update then dont worry about what is
	// in p.accessToken
	if !force && p.accessToken != "" {
		return p.accessToken, nil
	}

	// otherwise fetch the value and store it
	newToken, err := p.refreshAccessTokenImpl(ctx)
	p.accessToken = newToken
	return newToken, err
}

func (p *auth0TokenProvider) refreshAccessTokenImpl(ctx context.Context) (string, error) {
	// build the request
	params := &authRequest{
		ClientID:     p.clientID,
		ClientSecret: p.clientSecret,
		Audience:     p.audience,
		GrantType:    p.grantType,
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	// make the request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		p.domain+"/oauth/token",
		bytes.NewBuffer(jsonParams),
	)
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	} else if resp.StatusCode != http.StatusOK {
		txt, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("Auth0 returned %d. Failed to decode response with error: %e", resp.StatusCode, err)
		}
		return "", fmt.Errorf("Auth0 returned %d. Response body: %s", resp.StatusCode, txt)
	}

	// decode the response
	defer resp.Body.Close()
	authResp := &authResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return "", err
	}

	// setup a task to refresh token slightly before it expires
	// TODO - make this configurable
	go func() {
		// refresh auth between 2 and 5 sec before its required
		jitter, err := timeComm.Jitter(time.Second*2, time.Second*5)
		if err != nil {
			// if the jitter generator fails just use 5sec
			jitter = time.Second * 5
		}
		time.Sleep(
			(time.Duration(authResp.ExpiresInSecs) * time.Second) - jitter,
		)
		_, _ = p.RefreshAccessToken(p.appCtx, true)
	}()

	// return the token
	return authResp.AccessToken, err
}
