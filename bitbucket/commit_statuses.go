package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

// CommitStatuses represent a collection of a commit's statuses.
type CommitStatuses struct {
	PaginationInfo

	Values []*CommitStatus `json:"values,omitempty"`
}

// CommitStatus represents a commit status.
type CommitStatus struct {
	Links       *CSLinks   `json:"links,omitempty"`
	UUID        *string    `json:"uuid,omitempty"`
	Key         *string    `json:"key,omitempty"`
	Refname     *string    `json:"refname,omitempty"`
	URL         *string    `json:"url,omitempty"`
	State       *string    `json:"state,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedOn   *time.Time `json:"created_on,omitempty"`
	UpdatedOn   *time.Time `json:"updated_on,omitempty"`
}

// CSLinks represents the "links" object in a Bitbucket commit status.
type CSLinks struct {
	Self *Link `json:"self,omitempty"`
	HTML *Link `json:"html,omitempty"`
}

// CommitStatusRequest represents a new commit status.
type CommitStatusRequest struct {
	URL         *string `json:"url,omitempty"`
	State       *string `json:"state,omitempty"`
	Key         *string `json:"key,omitempty"`
	Refname     *string `json:"refname,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// ListStatuses returns all statuses (e.g. build results) for a specific commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses#get
func (c *CommitService) ListStatuses(owner, repoSlug, sha string, opts ...interface{}) (*CommitStatuses, *simpleresty.Response, error) {
	results := new(CommitStatuses)
	urlStr, urlStrErr := c.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/commit/%s/statuses", owner, repoSlug, sha), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := c.client.http.Get(urlStr, results, nil)

	return results, response, err
}

// CreateStatus creates a new build status against the specified commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build#post
func (c *CommitService) CreateStatus(owner, repoSlug, sha string, co *CommitStatusRequest) (*CommitStatus, *simpleresty.Response, error) {
	results := new(CommitStatus)
	urlStr := c.client.http.RequestURL("/repositories/%s/%s/commit/%s/statuses/build", owner, repoSlug, sha)
	response, err := c.client.http.Post(urlStr, results, co)

	return results, response, err
}

// UpdateStatus update the current status of a build status object on the specific commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build/%7Bkey%7D#put
func (c *CommitService) UpdateStatus(owner, repoSlug, sha, key string, co *CommitStatusRequest) (*CommitStatus, *simpleresty.Response, error) {
	results := new(CommitStatus)
	urlStr := c.client.http.RequestURL("/repositories/%s/%s/commit/%s/statuses/build/%s", owner, repoSlug, sha, key)
	response, err := c.client.http.Put(urlStr, results, co)

	return results, response, err
}

// GetStatusByBuild returns the specified build status for a commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build/%7Bkey%7D#get
func (c *CommitService) GetStatusByBuild(owner, repoSlug, sha, key string, opts ...interface{}) (*CommitStatus, *simpleresty.Response, error) {
	results := new(CommitStatus)
	urlStr, urlStrErr := c.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/commit/%s/statuses/build/%s", owner, repoSlug, sha, key), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := c.client.http.Get(urlStr, results, nil)

	return results, response, err
}
