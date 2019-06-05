package bitbucket

import (
	"time"
)

// RepositoriesService handles communication with the repository related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories
type RepositoriesService service

// Repositories represent a collection of repositories.
type Repositories struct {
	PaginationInfo

	Values []*Issue `json:"values,omitempty"`
}

// Repository represents a Bitbucket repository.
type Repository struct {
	SCM         *string               `json:"scm,omitempty"`
	Website     *string               `json:"page,omitempty"`
	HasIssues   *bool                 `json:"has_issues,omitempty"`
	HasWiki     *bool                 `json:"has_wiki,omitempty"`
	Language    *string               `json:"language,omitempty"`
	ForkPolicy  *string               `json:"fork_policy,omitempty"`
	Links       *RepositoryLinks      `json:"links,omitempty"`
	Name        *string               `json:"name,omitempty"`
	CreatedOn   *time.Time            `json:"created_on,omitempty"`
	MainBranch  *RepositoryMainBranch `json:"main_branch,omitempty"`
	FullName    *string               `json:"full_name,omitempty"`
	Owner       *User                 `json:"owner,omitempty"`
	UpdatedOn   *time.Time            `json:"updated_on,omitempty"`
	Size        *int64                `json:"size,omitempty"`
	Type        *string               `json:"type,omitempty"`
	Slug        *string               `json:"slug,omitempty"`
	IsPrivate   *bool                 `json:"is_private,omitempty"`
	Description *string               `json:"description,omitempty"`
	Parent      *Repository           `json:"parent,omitempty"`
}

// RepositoryMainBranch represents the primary branch set for a repository.
type RepositoryMainBranch struct {
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
}

// RepositoryCloneLink represents the specific clone related links in a Bitbucket repository.
type RepositoryCloneLink struct {
	HRef *string `json:"href,omitempty"`
	Name *string `json:"name,omitempty"`
}

// RepositoryLinks represents the "links" object in a Bitbucket repository.
type RepositoryLinks struct {
	Clone        []*RepositoryCloneLink `json:"clone,omitempty"`
	Watchers     *BitbucketLink         `json:"watchers,omitempty"`
	Branches     *BitbucketLink         `json:"branches,omitempty"`
	Tags         *BitbucketLink         `json:"tags,omitempty"`
	Commits      *BitbucketLink         `json:"commits,omitempty"`
	Downloads    *BitbucketLink         `json:"downloads,omitempty"`
	Source       *BitbucketLink         `json:"source,omitempty"`
	HTML         *BitbucketLink         `json:"html,omitempty"`
	Avatar       *BitbucketLink         `json:"avatar,omitempty"`
	Forks        *BitbucketLink         `json:"forks,omitempty"`
	Self         *BitbucketLink         `json:"self,omitempty"`
	PullRequests *BitbucketLink         `json:"pull_requests,omitempty"`
}

// RepositoryListOpts represents the filters and query parameters available when listing repositories.
type RepositoryListOpts struct {
	// Filters the result based on the authenticated user's role on each repository.
	// Valid roles:
	// - member: returns repositories to which the user has explicit read access
	// - contributor: returns repositories to which the user has explicit write access
	// - admin: returns repositories to which the user has explicit administrator access
	// - owner: returns all repositories owned by the current user
	Role string `url:"role,omitempty"`

	FilterSortOpts
}

// RepositoryRequest represents a request to create/update a repository.
// TODO: might need to break this apart as some fields aren't editable.
type RepositoryRequest struct {
	// Valid options for SCM are git or hg.
	SCM *string `json:"scm,omitempty"` // Required field.

	// Description of the new repository
	Description *string `json:"description,omitempty"`

	// Valid options: no_public_forks, no_forks, allow_forks
	ForkPolicy *string `json:"fork_policy,omitempty"`

	HasWiki   *bool   `json:"has_wiki,omitempty"`
	HasIssues *bool   `json:"has_issues,omitempty"`
	Name      *string `json:"name,omitempty"`

	// In order to set the project for the newly created repository,
	// pass in either the project key or the project UUID as part of the request body as shown in the examples below:
	Project struct {
		Key *string `json:"key,omitempty"`
	} `json:"project,omitempty"`
}

type RepositoryDeleteOpts struct {
	//If a repository has been moved to a new location, use this parameter to show users a friendly message
	// in the Bitbucket UI that the repository has moved to a new location.
	// However, a GET to this endpoint will still return a 404.
	RedirectTo string `url:"redirect_to,omitempty"`
}

// List all public repositories.
// Supports filtering by passing in a non-URI encoded query string. Reference: https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
// Example query string: parent.owner.username = "bitbucket"
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories#get
func (r *RepositoriesService) ListAll(opts ...interface{}) (*Repositories, *Response, error) {
	repositories := new(Repositories)
	urlStr := r.client.requestUrl("/repositories")
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, repositories, nil)

	return repositories, response, err
}

// List all repositories owned by the specified account or UUID.
// Accepts a query parameter for 'role.
// Supports filtering by passing in a non-URI encoded query string. Reference: https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
// Example query string: parent.owner.username = "bitbucket"
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D#get
func (r *RepositoriesService) List(owner string, opts ...interface{}) (*Repositories, *Response, error) {
	repositories := new(Repositories)
	urlStr := r.client.requestUrl("/repositories/%s", owner)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, repositories, nil)

	return repositories, response, err
}

// Get a single repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#get
func (r *RepositoriesService) Get(owner, repoSlug string) (*Repository, *Response, error) {
	repo := new(Repository)
	urlStr := r.client.requestUrl("/repositories/%s/%s", owner, repoSlug)
	response, err := r.client.execute("GET", urlStr, repo, nil)

	return repo, response, err
}

// Create a new repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#post
func (r *RepositoriesService) Create(owner string, rr *RepositoryRequest) (*Repository, *Response, error) {
	repo := new(Repository)
	urlStr := r.client.requestUrl("/repositories/%s/%s", owner, rr.GetName())
	response, err := r.client.execute("POST", urlStr, repo, rr)

	return repo, response, err
}

// Update a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#put
func (r *RepositoriesService) Update(owner, repoSlug string, rr *RepositoryRequest) (*Repository, *Response, error) {
	repo := new(Repository)
	urlStr := r.client.requestUrl("/repositories/%s/%s", owner, repoSlug)
	response, err := r.client.execute("PUT", urlStr, repo, rr)

	return repo, response, err
}

// Delete a repository.
// This is an irreversible operation.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#delete
func (r *RepositoriesService) Delete(owner, repoSlug string, opts ...interface{}) (*Response, error) {
	urlStr := r.client.requestUrl("/repositories/%s/%s", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, addOptErr
	}

	response, err := r.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}

// GetName returns the Name field if it's non-nil, empty string otherwise.
func (rr *RepositoryRequest) GetName() string {
	if rr == nil || rr.Name == nil {
		return ""
	}
	return *rr.Name
}
