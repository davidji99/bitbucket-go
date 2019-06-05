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
)

const DEFAULT_PAGE_LENGTH = 10

const (
	apiBaseURL = "https://api.bitbucket.org/2.0"
	userAgent  = "go-bitbucket"
)

// A Client manages communication with the Bitbucket API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

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

// BitbucketLink represents a single link object from Bitbucket object links.
type BitbucketLink struct {
	HRef *string `json:"href,omitempty"`
}

// BitbucketContent represents content found in a Bitbucket resource.
type BitbucketContent struct {
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

// FilterSortOpts represents the querying and sorting mechanism available
// to certain Bitbucket API resources that return multiple results in a response.
//
// Bitbucket API Docs: https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering#query-sort
type FilterSortOpts struct {
	// Query is the raw non-URL encoded query string.
	// Note that the entire query string is put in the Query field.
	// This library will take care of URL encoding the string for you.
	Query string `url:"q,omitempty"`

	// In principle, every field that can be queried can also be used as a key for sorting.
	// By default the sort order is ascending. To reverse the order, prefix the field name with a hyphen (e.g. ?sort=-updated_on).
	// Only one field can be sorted on. Compound fields (e.g. sort on state first, followed by updated_on) are not supported.
	Sort string `url:"sort,omitempty"`
}

// ListOpts specifies the optional parameters to various List methods that support pagination.
type ListOpts struct {
	// For paginated result sets, page of results to retrieve.
	Page int64 `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	// Globally, the minimum length is 10 and the maximum is 100. Some APIs may specify a different default.
	Pagelen int64 `url:"pagelen,omitempty"`
}

// PartialRespOpts represents the URL parameter to request a partial response and to add or remove
// specific fields from a response.
//
// Bitbucket API Docs: https://developer.atlassian.com/bitbucket/api/2/reference/meta/partial-response
type PartialRespOpts struct {
	Fields string `url:"fields,omitempty"`
}

// Uses the Client Credentials Grant oauth2 flow to authenticate to Bitbucket
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

func NewOAuth(i, s string) *Client {
	a := &auth{appID: i, secret: s}
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     i,
		ClientSecret: s,
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

func NewOAuthToken(t oauth2.Token) *Client {
	a := &auth{token: t}
	return injectClient(a)
}

func NewOAuthbearerToken(t string) *Client {
	a := &auth{bearerToken: t}
	return injectClient(a)
}

func NewBasicAuth(u, p string) *Client {
	a := &auth{user: u, password: p}
	return injectClient(a)
}

func injectClient(a *auth) *Client {
	c := &Client{Auth: a, Pagelen: DEFAULT_PAGE_LENGTH, BaseURL: apiBaseURL, UserAgent: userAgent, client: new(http.Client)}
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

func (c *Client) requestUrl(template string, args ...interface{}) string {
	if len(args) == 1 && args[0] == "" {
		return c.BaseURL + template
	}
	return c.BaseURL + fmt.Sprintf(template, args...)
}

func (c *Client) newRequest(method string, urlStr string, v, body interface{}) (*http.Request, error) {
	// Use pagination if changed from default value
	const DEC_RADIX = 10
	if strings.Contains(urlStr, "/repositories/") {
		if c.Pagelen != DEFAULT_PAGE_LENGTH {
			urlObj, err := url.Parse(urlStr)
			if err != nil {
				return nil, err
			}
			q := urlObj.Query()
			q.Set("pagelen", strconv.FormatUint(c.Pagelen, DEC_RADIX))
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

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, fmt.Errorf(resp.Status)
	}

	if emptyResponse {
		return nil, nil
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}

	response := newResponse(resp)

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
			errorResponse.Message = "failed to parse unknown error format"
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

//// addOptions adds the parameters in opt as URL query parameters to s. opt
//// must be a struct whose fields may contain "url" tags.
//// Credit: https://github.com/google/go-github/blob/master/github/github.go#L226
//func addOptions(s string, opt interface{}) (string, error) {
//	v := reflect.ValueOf(opt)
//	if v.Kind() == reflect.Ptr && v.IsNil() {
//		return s, nil
//	}
//
//	u, err := url.Parse(s)
//	if err != nil {
//		return s, err
//	}
//
//	qs, err := query.Values(opt)
//	if err != nil {
//		return s, err
//	}
//
//	u.RawQuery = qs.Encode()
//	return u.String(), nil
//}

// addOptions takes a slice of opts and adds the parameters in each opt as URL query parameters to s.
// each opt must be a struct whose fields may contain "url" tags.
// Based on: https://github.com/google/go-github/blob/master/github/github.go#L226
func addOptions(s string, opts ...interface{}) (string, error) {
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
