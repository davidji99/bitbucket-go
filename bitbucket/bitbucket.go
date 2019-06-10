package bitbucket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/bitbucket"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// DefaultPageLength represents the default page length returned from API calls.
const DefaultPageLength = 10

const (
	apiBaseURL = "https://api.bitbucket.org/2.0"
	userAgent  = "go-bitbucket"
)

// A Client manages communication with the Bitbucket API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Bitbucket API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL string

	// User agent used when communicating with the Bitbucket API.
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Services used for talking to different parts of the Bitbucket API.
	BranchRestrictions *BranchRestrictionsService
	Commit             *CommitService
	Commits            *CommitsService
	Components         *ComponentsService
	DefaultReviewers   *DefaultReviewersService
	DeployKeys         *DeployKeysService
	Diff               *DiffService
	Downloads          *DownloadsService
	FileHistory        *FileHistoryService
	Forks              *ForksService
	HookEvents         *HookEventsService
	Issues             *IssuesService
	Milestones         *MilestonesService
	Patch              *PatchService
	PullRequests       *PullRequestsService
	Refs               *RefsService
	Repositories       *RepositoriesService
	SRC                *SRCService
	Teams              *TeamsService
	User               *UserService
	Users              *UsersService
	Versions           *VersionsService
	Watchers           *WatchersService

	Pagelen uint64
	Auth    *auth
}

type auth struct {
	appID, secret  string
	user, password string
	token          oauth2.Token
	bearerToken    string
}

type service struct {
	client *Client // TODO: rename this to API
}

// Link represents a single link object from Bitbucket object links.
type Link struct {
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

// NewOAuthClientCredentials uses the Client Credentials Grant oauth2 flow to authenticate to Bitbucket
func NewOAuthClientCredentials(i, s string) *Client {
	a := &auth{appID: i, secret: s}
	ctx := context.Background()
	conf := &clientcredentials.Config{
		ClientID:     i,
		ClientSecret: s,
		TokenURL:     bitbucket.Endpoint.TokenURL,
	}

	tok, err := conf.Token(ctx)
	if err != nil {
		log.Fatal(err)
	}
	a.token = *tok
	return injectClient(a)

}

// NewOAuth creates a new oauth.
func NewOAuth(clientID, clientSecret string) *Client {
	a := &auth{appID: clientID, secret: clientSecret}
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
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	a.token = *tok
	return injectClient(a)
}

// NewOAuthWithCode finishes the OAuth handshake with a given code
// and returns a *Client
func NewOAuthWithCode(i, s, c string) (*Client, string) {
	a := &auth{appID: i, secret: s}
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     i,
		ClientSecret: s,
		Endpoint:     bitbucket.Endpoint,
	}

	tok, err := conf.Exchange(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	a.token = *tok
	return injectClient(a), tok.AccessToken
}

// NewOAuthToken creates a new oauth with otken.
func NewOAuthToken(t oauth2.Token) *Client {
	a := &auth{token: t}
	return injectClient(a)
}

// NewOAuthbearerToken creates a new oauth with the bearer token.
func NewOAuthbearerToken(t string) *Client {
	a := &auth{bearerToken: t}
	return injectClient(a)
}

// NewBasicAuth creates a new client using username and preferably an app password.
func NewBasicAuth(u, p string) *Client {
	a := &auth{user: u, password: p}
	return injectClient(a)
}

// injectClient adds all resouce services to the client.
func injectClient(a *auth) *Client {
	c := &Client{Auth: a, Pagelen: DefaultPageLength, BaseURL: apiBaseURL, UserAgent: userAgent, client: new(http.Client)}
	c.common.client = c
	c.BranchRestrictions = (*BranchRestrictionsService)(&c.common)
	c.Commit = (*CommitService)(&c.common)
	c.Commits = (*CommitsService)(&c.common)
	c.Components = (*ComponentsService)(&c.common)
	c.DefaultReviewers = (*DefaultReviewersService)(&c.common)
	c.DeployKeys = (*DeployKeysService)(&c.common)
	c.Diff = (*DiffService)(&c.common)
	c.Downloads = (*DownloadsService)(&c.common)
	c.FileHistory = (*FileHistoryService)(&c.common)
	c.Forks = (*ForksService)(&c.common)
	c.HookEvents = (*HookEventsService)(&c.common)
	c.Issues = (*IssuesService)(&c.common)
	c.Milestones = (*MilestonesService)(&c.common)
	c.Patch = (*PatchService)(&c.common)
	c.PullRequests = (*PullRequestsService)(&c.common)
	c.Refs = (*RefsService)(&c.common)
	c.Repositories = (*RepositoriesService)(&c.common)
	c.SRC = (*SRCService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.User = (*UserService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Versions = (*VersionsService)(&c.common)
	c.Watchers = (*WatchersService)(&c.common)

	return c
}

func (c *Client) requestURL(template string, args ...interface{}) string {
	if len(args) == 1 && args[0] == "" {
		return c.BaseURL + template
	}
	return c.BaseURL + fmt.Sprintf(template, args...)
}

func (c *Client) newRequest(method string, urlStr string, v, body interface{}) (*http.Request, error) {
	// Use pagination if changed from default value
	const DecRadix = 10
	if strings.Contains(urlStr, "/repositories/") {
		if c.Pagelen != DefaultPageLength {
			urlObj, err := url.Parse(urlStr)
			if err != nil {
				return nil, err
			}
			q := urlObj.Query()
			q.Set("pagelen", strconv.FormatUint(c.Pagelen, DecRadix))
			urlObj.RawQuery = q.Encode()
			urlStr = urlObj.String()
		}
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, urlStr, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) execute(method string, urlStr string, v, body interface{}) (*Response, error) {
	req, reqErr := c.newRequest(method, urlStr, v, body)
	if reqErr != nil {
		return nil, reqErr
	}

	response, err := c.doRequest(req, v, false)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) doRequest(req *http.Request, v interface{}, emptyResponse bool) (*Response, error) {
	c.addAuthHeaders(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if emptyResponse {
		return nil, nil
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}

	response := newResponse(resp)

	err = CheckResponse(resp)

	if err != nil {
		// Special case for AcceptedErrors. If an AcceptedError
		// has been encountered, the response's payload will be
		// added to the AcceptedError and returned.
		//
		// Issue #1022
		aerr, ok := err.(*AcceptedError)
		if ok {
			b, readErr := ioutil.ReadAll(resp.Body)
			if readErr != nil {
				return response, readErr
			}

			aerr.Raw = b
			return response, aerr
		}

		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, copyErr := io.Copy(w, resp.Body)
			if copyErr != nil {
				return nil, copyErr
			}
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

func (c *Client) addAuthHeaders(req *http.Request) {
	if c.Auth.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.Auth.bearerToken)
	}

	if c.Auth.user != "" && c.Auth.password != "" {
		req.SetBasicAuth(c.Auth.user, c.Auth.password)
	} else if c.Auth.token.Valid() {
		c.Auth.token.SetAuthHeader(req)
	}
	return
}

// Response represents a response returned from this client.
type Response struct {
	*http.Response

	Page    int
	Next    string
	PageLen int
	size    int
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Body     []byte
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, u, e.Response.StatusCode, e.Message)
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		errorResponse.Body = data

		var raw interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			errorResponse.Message = string(data)
		} else {
			errorResponse.Message = parseError(raw)
		}
	}

	return errorResponse
}

func parseError(raw interface{}) string {
	switch raw := raw.(type) {
	case string:
		return raw

	case []interface{}:
		var errs []string
		for _, v := range raw {
			errs = append(errs, parseError(v))
		}
		return fmt.Sprintf("[%s]", strings.Join(errs, ", "))

	case map[string]interface{}:
		var errs []string
		for k, v := range raw {
			errs = append(errs, fmt.Sprintf("{%s: %s}", k, parseError(v)))
		}
		sort.Strings(errs)
		return strings.Join(errs, ", ")

	default:
		return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
	}
}

// addQueryParams takes a slice of opts and adds each field as escaped URL query parameters to s.
// Each element in opts must be a struct whose fields contain "url" tags.
//
// Based on: https://github.com/google/go-github/blob/master/github/github.go#L226
func addQueryParams(s string, opts ...interface{}) (string, error) {
	// Handle if opts is nil
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Slice && v.IsNil() {
		return s, nil
	}

	// Parse URL
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	fulQS := url.Values{}
	for _, opt := range opts {
		//// Handle scenario when no opts are passed which means opts is a slice containing one empty slice.
		//v := reflect.ValueOf(opt)
		//if v.Kind() == reflect.Slice && v.IsNil() {
		//	return s, nil
		//}

		qs, err := query.Values(opt)
		if err != nil {
			return s, err
		}

		for k, v := range qs {
			fulQS[k] = v
		}
	}

	u.RawQuery = fulQS.Encode()
	return u.String(), nil
}
