package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

// RepositoryHooks represents a collection of repository hooks.
type RepositoryHooks struct {
	PaginationInfo

	Values []*RepositoryHook `json:"values,omitempty"`
}

// RepositoryHook represents a repository hook.
type RepositoryHook struct {
	UUID        *string    `json:"uuid,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Description *string    `json:"description,omitempty"`
	SubjectType []*string  `json:"subject_type,omitempty"`
	Active      *bool      `json:"active,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Events      []*string  `json:"events,omitempty"`
}

// RepositoryHookRequest represents a request to create/update a hook.
type RepositoryHookRequest struct {
	Description *string   `json:"description,omitempty"`
	URL         *string   `json:"url,omitempty"`
	Active      *bool     `json:"active,omitempty"`
	Events      []*string `json:"events,omitempty"`
}

// ListHooks returns a paginated list of webhooks installed on a specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/hooks#get
func (r *RepositoriesService) ListHooks(owner, repoSlug string, opts ...interface{}) (*RepositoryHooks, *simpleresty.Response, error) {
	result := new(RepositoryHooks)
	urlStr, urlStrErr := r.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/hooks", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := r.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// CreateHook creates a new webhook on the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/hooks#post
func (r *RepositoriesService) CreateHook(owner, repoSlug string, rho *RepositoryHookRequest) (*RepositoryHook, *simpleresty.Response, error) {
	result := new(RepositoryHook)
	urlStr := r.client.http.RequestURL("/repositories/%s/%s/hooks", owner, repoSlug)
	response, err := r.client.http.Post(urlStr, result, rho)

	return result, response, err
}

// GetHook returns the webhook with the specified id installed on the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/hooks/%7Buid%7D#get
func (r *RepositoriesService) GetHook(owner, repoSlug, uid string, opts ...interface{}) (*RepositoryHook, *simpleresty.Response, error) {
	result := new(RepositoryHook)
	urlStr, urlStrErr := r.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/hooks/%s", owner, repoSlug, uid), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := r.client.http.Post(urlStr, result, nil)

	return result, response, err
}

// UpdateHook updates the specified webhook subscription.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/hooks/%7Buid%7D#put
func (r *RepositoriesService) UpdateHook(owner, repoSlug, uid string, rho *RepositoryHookRequest) (*RepositoryHook, *simpleresty.Response, error) {
	result := new(RepositoryHook)
	urlStr := r.client.http.RequestURL("/repositories/%s/%s/hooks/%s", owner, repoSlug, uid)
	response, err := r.client.http.Put(urlStr, result, rho)

	return result, response, err
}

// DeleteHook deletes the specified webhook subscription from the given repository.
// This is an irreversible operation.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/hooks/%7Buid%7D#delete
func (r *RepositoriesService) DeleteHook(owner, repoSlug, uid string) (*simpleresty.Response, error) {
	urlStr := r.client.http.RequestURL("/repositories/%s/%s/hooks/%s", owner, repoSlug, uid)
	response, err := r.client.http.Delete(urlStr, nil, nil)

	return response, err
}
