package bitbucket

import (
	"context"
	"fmt"
	"github.com/davidji99/simpleresty"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/bitbucket"
	"golang.org/x/oauth2/clientcredentials"
)

// Option is a functional option for configuring the API client.
type Option func(*Client) error

func HTTP(http *simpleresty.Client) Option {
	return func(c *Client) error {
		c.http = http
		return nil
	}
}

// UserAgent allows overriding of the default User Agent.
func UserAgent(userAgent string) Option {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

// BaseURL allows overriding of the default base API URL.
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		// Validate that there is no trailing slashes before setting the custom baseURL
		if baseURL[len(baseURL)-1:] == "/" {
			return fmt.Errorf("custom base URL cannot contain a trailing slash")
		}

		c.baseURL = baseURL
		return nil
	}
}

// CustomHTTPHeaders sets additional HTTPHeaders
func CustomHTTPHeaders(headers map[string]string) Option {
	return func(c *Client) error {
		c.customHTTPHeaders = headers
		return nil
	}
}

// OAuthClientCredentials uses the Client Credentials Grant oauth2 flow to authenticate to Bitbucket.
func OAuthClientCredentials(clientID, clientSecret string) Option {
	return func(c *Client) error {
		ctx := context.Background()
		conf := &clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     bitbucket.Endpoint.TokenURL,
		}

		token, err := conf.Token(ctx)
		if err != nil {
			return err
		}

		c.oauthGrant = token

		return nil
	}
}

// OAuth with oauth.
func OAuth(clientID, clientSecret string) Option {
	return func(c *Client) error {
		ctx := context.Background()
		conf := &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint:     bitbucket.Endpoint,
		}

		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog:\n%v", url)

		// Use the authorization code that is pushed to the redirect
		// URL. Exchange will do the handshake to retrieve the
		// initial access token. The HTTP Client returned by
		// conf.Client will refresh the token as necessary.
		var code string
		fmt.Printf("Enter the code in the return URL: ")
		if _, err := fmt.Scan(&code); err != nil {
			return err
		}

		token, err := conf.Exchange(ctx, code)
		if err != nil {
			return err
		}

		c.oauthGrant = token

		return nil
	}
}

// OAuthWithCode does the OAuth handshake with a given code.
func OAuthWithCode(clientID, clientSecret, code string) Option {
	return func(c *Client) error {
		ctx := context.Background()
		conf := &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint:     bitbucket.Endpoint,
		}

		token, err := conf.Exchange(ctx, code)
		if err != nil {
			return err
		}
		c.oauthGrant = token

		return nil
	}
}

// OAuthToken sets the oauth grant.
func OAuthToken(ot oauth2.Token) Option {
	return func(c *Client) error {
		c.oauthGrant = &ot
		return nil
	}
}

// OAuthBearerToken sets the oauth grant.
func OAuthBearerToken(t string) Option {
	return func(c *Client) error {
		c.bearerToken = &t
		return nil
	}
}
