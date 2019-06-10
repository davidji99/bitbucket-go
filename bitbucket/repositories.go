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
	Watchers     *Link                  `json:"watchers,omitempty"`
	Branches     *Link                  `json:"branches,omitempty"`
	Tags         *Link                  `json:"tags,omitempty"`
	Commits      *Link                  `json:"commits,omitempty"`
	Downloads    *Link                  `json:"downloads,omitempty"`
	Source       *Link                  `json:"source,omitempty"`
	HTML         *Link                  `json:"html,omitempty"`
	Avatar       *Link                  `json:"avatar,omitempty"`
	Forks        *Link                  `json:"forks,omitempty"`
	Self         *Link                  `json:"self,omitempty"`
	PullRequests *Link                  `json:"pull_requests,omitempty"`
}

// RepositoryListQueryParams represents the filters and query parameters available when listing repositories.
type RepositoryListQueryParams struct {
	// Filters the result based on the authenticated user's role on each repository.
	// Valid roles:
	// - member: returns repositories to which the user has explicit read access
	// - contributor: returns repositories to which the user has explicit write access
	// - admin: returns repositories to which the user has explicit administrator access
	// - owner: returns all repositories owned by the current user
	Role string `url:"role,omitempty"`
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

// RepositoryDeleteQueryParam represents the query parameter available when deleting a repository.
type RepositoryDeleteQueryParam struct {
	//If a repository has been moved to a new location, use this parameter to show users a friendly message
	// in the Bitbucket UI that the repository has moved to a new location.
	// However, a GET to this endpoint will still return a 404.
	RedirectTo string `url:"redirect_to,omitempty"`
}

// ListPublic returns all public repositories.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories#get
func (r *RepositoriesService) ListPublic(opts ...interface{}) (*Repositories, *Response, error) {
	result := new(Repositories)
	urlStr := r.client.requestURL("/repositories")
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// List all repositories owned by the specified account or UUID.
//
// Accepts a query parameter for 'role.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D#get
func (r *RepositoriesService) List(owner string, opts ...interface{}) (*Repositories, *Response, error) {
	result := new(Repositories)
	urlStr := r.client.requestURL("/repositories/%s", owner)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Get a single repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#get
func (r *RepositoriesService) Get(owner, repoSlug string, opts ...interface{}) (*Repository, *Response, error) {
	result := new(Repository)
	urlStr := r.client.requestURL("/repositories/%s/%s", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Create a new repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#post
func (r *RepositoriesService) Create(owner string, rr *RepositoryRequest) (*Repository, *Response, error) {
	result := new(Repository)
	urlStr := r.client.requestURL("/repositories/%s/%s", owner, rr.GetName())
	response, err := r.client.execute("POST", urlStr, result, rr)

	return result, response, err
}

// Update a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#put
func (r *RepositoriesService) Update(owner, repoSlug string, rr *RepositoryRequest) (*Repository, *Response, error) {
	result := new(Repository)
	urlStr := r.client.requestURL("/repositories/%s/%s", owner, repoSlug)
	response, err := r.client.execute("PUT", urlStr, result, rr)

	return result, response, err
}

// Delete a repository.
// This is an irreversible operation.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#delete
func (r *RepositoriesService) Delete(owner, repoSlug string, deleteOpt *RepositoryDeleteQueryParam) (*Response, error) {
	urlStr := r.client.requestURL("/repositories/%s/%s", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, deleteOpt)
	if addOptErr != nil {
		return nil, addOptErr
	}

	response, err := r.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
