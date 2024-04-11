package auth

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"
)

var logger = log.New(log.Writer(), "[stately-sdk][auth_token_provider] ", log.LstdFlags)

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
	defaultDomain    = "https://oauth.stately.cloud"
	defaultAudience  = "api.stately.cloud"
	defaultGrantType = "client_credentials"
)

type authTokenProvider struct {
	ctx context.Context

	clientID     string
	clientSecret string

	domain      string
	audience    string
	grantType   string
	accessToken string
	mutex       *sync.RWMutex
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
) (*authTokenProvider, error) {
	options = applyDefaults(options)
	p := &authTokenProvider{
		ctx:          appCtx,
		clientID:     clientID,
		clientSecret: clientSecret,
		domain:       options.Domain,
		audience:     options.Audience,
		grantType:    defaultGrantType,
		accessToken:  "",
		mutex:        &sync.RWMutex{},
	}

	// refresh access token as soon as we create this thing so the first request is faster
	go func() {
		_, err := p.RefreshAccessToken(p.ctx, false)
		if err != nil {
			logger.Printf("Error performing initial refresh: %e", err)
		}
	}()
	return p, nil
}

func (p *authTokenProvider) GetAccessToken(ctx context.Context) (string, error) {
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

func (p *authTokenProvider) RefreshAccessToken(ctx context.Context, force bool) (string, error) {
	// take a full lock. there is only one caller in this function at once so its
	// totally safe to update the state
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// if someone beat us to the lock and already did a refresh
	// then simply return the value. We know that happened because accessToken is not empty
	//
	// if we are forcing an update then don't worry about what is
	// in p.accessToken
	if !force && p.accessToken != "" {
		return p.accessToken, nil
	}

	// otherwise fetch the value and store it
	newToken, err := p.refreshAccessTokenImpl(ctx)
	p.accessToken = newToken
	return newToken, err
}

// TODO - just use auth0 go SDK LoginWithClientCredentials or RefreshToken so we don't have
// to manually make the network request here:
// https://github.com/auth0/go-auth0/blob/main/authentication/oauth.go#L136-L181
func (p *authTokenProvider) refreshAccessTokenImpl(ctx context.Context) (string, error) {
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
	// TODO - this needs to stop when the context is cancelled
	// TODO - this probably shouldn't refresh in the background automatically - instead, if we detect it's time to refresh *during* a request, we should refresh in the background while still using the old token
	go func() {
		// refresh auth between 2 and 5 sec before its required
		jitter, err := jitter(time.Second*2, time.Second*5)
		if err != nil {
			// if the jitter generator fails just use 5sec
			logger.Printf("Error generating jitter: %e. Using default value: 5 seconds", err)
			jitter = time.Second * 5
		}
		time.Sleep(
			(time.Duration(authResp.ExpiresInSecs) * time.Second) - jitter,
		)
		_, err = p.RefreshAccessToken(p.ctx, true)
		if err == nil {
			logger.Printf("Error performing scheduled refresh: %e", err)
		}
	}()

	// return the token
	return authResp.AccessToken, err
}

// jitter creates a cryptographically random duration between the given bounds.
func jitter(min, max time.Duration) (time.Duration, error) {
	jitterNanos, err := rand.Int(rand.Reader, big.NewInt(max.Nanoseconds()-min.Nanoseconds()))
	if err != nil {
		return 0, err
	}
	return min + time.Duration(jitterNanos.Int64()), nil
}
