package bitbucket

import "time"

// RepositoriesService handles communication with the repository related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D
type RepositoriesService service

// Repositories represent a collection of repositories.
type Repositories struct {
	Page     int    `json:"page,omitempty"`
	Next     string `json:"next,omitempty"`
	PageLen  int    `json:"pagelen,omitempty"`
	Size     int    `json:"size,omitempty"`
	Previous string `json:"previous,omitempty"`

	Values []*Issue `json:"values,omitempty"`
}

// Repository represents a Bitbucket repository.
type Repository struct {
	SCM        *string          `json:"scm,omitempty"`
	Website    *string          `json:"page,omitempty"`
	HasWiki    *bool            `json:"has_wiki,omitempty"`
	Language   *string          `json:"language,omitempty"`
	ForkPolicy *string          `json:"fork_policy,omitempty"`
	Links      *RepositoryLinks `json:"links,omitempty"`
	Name       *string          `json:"name,omitempty"`
	CreatedOn  *time.Time       `json:"created_on,omitempty"`
	MainBranch struct {
		Type *string `json:"type,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"main_branch,omitempty"`
	FullName    *string    `json:"full_name,omitempty"`
	Owner       *User      `json:"owner,omitempty"`
	UpdatedOn   *time.Time `json:"updated_on,omitempty"`
	Size        *int64     `json:"size,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Slug        *string    `json:"slug,omitempty"`
	IsPrivate   *bool      `json:"is_private,omitempty"`
	Description *string    `json:"description,omitempty"`
}

type RepositoryCloneLink struct {
	HRef *string `json:"href,omitempty"`
	Name *string `json:"name,omitempty"`
}

type RepositoryLinks struct {
	Clone        *RepositoryCloneLink `json:"clone,omitempty"`
	Watchers     *BitbucketLink       `json:"watchers,omitempty"`
	Branches     *BitbucketLink       `json:"branches,omitempty"`
	Tags         *BitbucketLink       `json:"tags,omitempty"`
	Commits      *BitbucketLink       `json:"commits,omitempty"`
	Downloads    *BitbucketLink       `json:"downloads,omitempty"`
	Source       *BitbucketLink       `json:"source,omitempty"`
	HTML         *BitbucketLink       `json:"html,omitempty"`
	Avatar       *BitbucketLink       `json:"avatar,omitempty"`
	Forks        *BitbucketLink       `json:"forks,omitempty"`
	Self         *BitbucketLink       `json:"self,omitempty"`
	PullRequests *BitbucketLink       `json:"pull_requests,omitempty"`
}
