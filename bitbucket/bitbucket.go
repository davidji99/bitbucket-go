package bitbucket

import (
	"golang.org/x/oauth2"
)

type auth struct {
	appID, secret  string
	user, password string
	token          oauth2.Token
	bearerToken    string
}

// Link represents a single link object from Bitbucket object links.
type Link struct {
	Name *string `json:"name,omitempty"`
	HRef *string `json:"href,omitempty"`
}

// Content represents content found in a Bitbucket resource.
type Content struct {
	Raw    *string `json:"raw,omitempty"`
	Markup *string `json:"markup,omitempty"`
	HTML   *string `json:"html,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// PaginationInfo represents the pagination data returned on most LIST functions.
//
// Bitbucket API Docs: https://developer.atlassian.com/bitbucket/api/2/reference/meta/pagination
type PaginationInfo struct {
	// Page number of the current results. This is an optional element that is not provided in all responses.
	Page *int64 `json:"page,omitempty"`

	//  Link to the next page if it exists. The last page of a collection does not have this value.
	//  Use this link to navigate the result set and refrain from constructing your own URLs.
	Next *string `json:"next,omitempty"`

	// Current number of objects on the existing page.
	Pagelen *int64 `json:"pagelen,omitempty"`

	// Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Size *int64 `json:"size,omitempty"`

	//Link to previous page if it exists. A collections first page does not have this value.
	// This is an optional element that is not provided in all responses.
	// Some result sets strictly support forward navigation and never provide previous links.
	// Clients must anticipate that backwards navigation is not always available.
	// Use this link to navigate the result set and refrain from constructing your own URLs.
	Previous *string `json:"previous,omitempty"`
}

//
//// NewOAuthClientCredentials uses the Client Credentials Grant oauth2 flow to authenticate to Bitbucket
//func NewOAuthClientCredentials(i, s string) *Client {
//	a := &auth{appID: i, secret: s}
//	ctx := context.Background()
//	conf := &clientcredentials.Config{
//		ClientID:     i,
//		ClientSecret: s,
//		TokenURL:     bitbucket.Endpoint.TokenURL,
//	}
//
//	tok, err := conf.Token(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	a.token = *tok
//	return injectClient(a)
//
//}
//
//// NewOAuth creates a new oauth.
//func NewOAuth(clientID, clientSecret string) *Client {
//	a := &auth{appID: clientID, secret: clientSecret}
//	ctx := context.Background()
//	conf := &oauth2.Config{
//		ClientID:     clientID,
//		ClientSecret: clientSecret,
//		Endpoint:     bitbucket.Endpoint,
//	}
//
//	// Redirect user to consent page to ask for permission
//	// for the scopes specified above.
//	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
//	fmt.Printf("Visit the URL for the auth dialog:\n%v", url)
//
//	// Use the authorization code that is pushed to the redirect
//	// URL. Exchange will do the handshake to retrieve the
//	// initial access token. The HTTP Client returned by
//	// conf.Client will refresh the token as necessary.
//	var code string
//	fmt.Printf("Enter the code in the return URL: ")
//	if _, err := fmt.Scan(&code); err != nil {
//		log.Fatal(err)
//	}
//	tok, err := conf.Exchange(ctx, code)
//	if err != nil {
//		log.Fatal(err)
//	}
//	a.token = *tok
//	return injectClient(a)
//}

//// NewOAuthWithCode finishes the OAuth handshake with a given code
//// and returns a *Client
//func NewOAuthWithCode(i, s, c string) (*Client, string) {
//	a := &auth{appID: i, secret: s}
//	ctx := context.Background()
//	conf := &oauth2.Config{
//		ClientID:     i,
//		ClientSecret: s,
//		Endpoint:     bitbucket.Endpoint,
//	}
//
//	tok, err := conf.Exchange(ctx, c)
//	if err != nil {
//		log.Fatal(err)
//	}
//	a.token = *tok
//	return injectClient(a), tok.AccessToken
//}
//
//// NewOAuthToken creates a new oauth with otken.
//func NewOAuthToken(t oauth2.Token) *Client {
//	a := &auth{token: t}
//	return injectClient(a)
//}
//
//// NewOAuthbearerToken creates a new oauth with the bearer token.
//func NewOAuthbearerToken(t string) *Client {
//	a := &auth{bearerToken: t}
//	return injectClient(a)
//}
