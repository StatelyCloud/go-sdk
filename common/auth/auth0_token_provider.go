package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	timeComm "github.com/StatelyCloud/go-sdk/common/time"
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

// TokenProvider is the interface which must be passed into the WithAuthToken{Unary,Stream}Interceptor
// so that it can authenticate outgoing requests.
// This is a thread-safe interface.
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
	// default vals.
	defaultDomain    = "https://oauth.stately.cloud"
	defaultAudience  = "api.stately.cloud"
	defaultGrantType = "client_credentials"

	// env vars.
	clientIDEnvVar     = "STATELY_CLIENT_ID"
	clientSecretEnvVar = "STATELY_CLIENT_SECRET"
)

type authTokenProvider struct {
	ctx         context.Context
	credentials *Credentials
	domain      string
	audience    string
	grantType   string
	accessToken string
	mutex       *sync.RWMutex
}

// Credentials is a struct containing the client ID and and secret to be used
// when requesting a token.
type Credentials struct {
	ClientID     string
	ClientSecret string
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
	// Credentials is the the client ID and secret to use when requesting auth.
	// Defaults to value of the STATELY_CLIENT_ID and STATELY_CLIENT_SECRET env vars
	Credentials *Credentials
}

func newDefaultCredentials() (*Credentials, error) {
	clientID := os.Getenv(clientIDEnvVar)
	if clientID == "" {
		return nil, fmt.Errorf("unable to read client ID from %s env var", clientIDEnvVar)
	}

	clientSecret := os.Getenv(clientSecretEnvVar)
	if clientSecret == "" {
		return nil, fmt.Errorf("unable to read client secret from %s env var", clientSecretEnvVar)
	}

	return &Credentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}, nil
}

func newDefaultOptions() (*Options, error) {
	creds, err := newDefaultCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to read default credentials from environment")
	}
	return &Options{
		Domain:      defaultDomain,
		Audience:    defaultAudience,
		Credentials: creds,
	}, nil
}

// applyDefaults iterates through the given options struct and applied default values
// where required.
func applyDefaults(options *Options) (*Options, error) {
	if options == nil {
		return newDefaultOptions()
	}

	// Domain
	if options.Domain == "" {
		options.Domain = defaultDomain
	}

	// Audience
	if options.Audience == "" {
		options.Audience = defaultAudience
	}

	// Credentials
	if options.Credentials == nil {
		creds, err := newDefaultCredentials()
		if err != nil {
			return nil, err
		}
		options.Credentials = creds
	}

	return options, nil
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
func NewAuthTokenProvider(
	appCtx context.Context,
	options *Options,
) (TokenProvider, error) {
	options, err := applyDefaults(options)
	if err != nil {
		return nil, err
	}
	p := &authTokenProvider{
		ctx:         appCtx,
		domain:      options.Domain,
		audience:    options.Audience,
		grantType:   defaultGrantType,
		credentials: options.Credentials,
		accessToken: "",
		mutex:       &sync.RWMutex{},
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
		ClientID:     p.credentials.ClientID,
		ClientSecret: p.credentials.ClientSecret,
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
		jitter, err := timeComm.Jitter(time.Second*2, time.Second*5)
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
