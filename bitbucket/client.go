package bitbucket

import (
	"encoding/base64"
	"github.com/davidji99/simpleresty"
	"golang.org/x/oauth2"
	"time"
)

const (
	// DefaultAPIBaseURL for the base API URL.
	DefaultAPIBaseURL = "https://api.bitbucket.org/2.0"

	// DefaultUserAgent for the API calls.
	DefaultUserAgent = "bitbucket-go"

	// DefaultPageLength represents the default page length returned from API calls.
	DefaultPageLength = 10
)

// A Client manages communication with the Bitbucket API.
type Client struct {
	http *simpleresty.Client // HTTP client used to communicate with the API.

	// baseURL for API requests. Defaults to the public Bitbucket APIv2 URL.
	baseURL string

	// userAgent used when communicating with the Bitbucket APIv2.
	userAgent string

	// Custom HTTPHeaders
	customHTTPHeaders map[string]string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// basicAuth represents the base64 encoded string for authentication
	basicAuth *string

	// bearerToken
	bearerToken *string

	// oauthGrant
	oauthGrant *oauth2.Token

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
}

type service struct {
	client *Client
}

// New creates a new client using username and an app password, essentially basic authentication.
func New(username, appPassword string, opts ...Option) (*Client, error) {
	c := &Client{
		http:        simpleresty.New(),
		baseURL:     DefaultAPIBaseURL,
		userAgent:   DefaultUserAgent,
		basicAuth:   nil,
		bearerToken: nil,
	}

	// Define any user custom Client settings
	if optErr := c.parseOptions(opts...); optErr != nil {
		return nil, optErr
	}

	// Generate base64 string using username & appPassword
	a := base64.StdEncoding.EncodeToString([]byte(username + ":" + appPassword))
	c.basicAuth = &a

	// Setup the client with default settings
	c.setupClient()

	// Inject services
	c.injectServices()

	return c, nil
}

// setupClient sets common headers and other configurations.
func (c *Client) setupClient() {
	// Set Base URL for the http client
	c.http.SetBaseURL(c.baseURL)

	// Set basic headers
	c.http.SetHeader("Content-type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", c.userAgent).
		SetTimeout(1 * time.Minute).
		SetAllowGetMethodPayload(true)

	// Set Auth headers
	c.addAuthHeaders()

	// Set additional headers
	if c.customHTTPHeaders != nil {
		c.http.SetHeaders(c.customHTTPHeaders)
	}
}

func (c *Client) addAuthHeaders() {
	if c.bearerToken != nil {
		c.http.SetHeader("Authorization", "Bearer "+*c.bearerToken)
	}

	if c.basicAuth != nil {
		c.http.SetHeader("Authorization", "Basic "+*c.basicAuth)
		return
	}

	if c.oauthGrant.Valid() {
		c.oauthGrant.SetAuthHeader(c.http.R().RawRequest)
	}
	return
}

// injectClient adds all resource services to the client.
func (c *Client) injectServices() *Client {
	//c := &Client{Auth: a, Pagelen: DefaultPageLength, BaseURL: apiBaseURL, UserAgent: userAgent, client: new(http.Client)}
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

// parseOptions parses the supplied options functions and returns a configured *Client instance.
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}
