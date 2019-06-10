package bitbucket

// ForksService handles communication with the issue related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/forks
type ForksService service

// ForkRequest represents a request to a create a new fork.
type ForkRequest struct {
	SCM         *string               `json:"scm,omitempty"`
	Name        *string               `json:"name,omitempty"`
	MainBranch  *RepositoryMainBranch `json:"main_branch,omitempty"`
	IsPrivate   *bool                 `json:"is_private,omitempty"`
	Language    *string               `json:"language,omitempty"`
	ForkPolicy  *string               `json:"fork_policy,omitempty"`
	Description *string               `json:"description,omitempty"`
	HasWiki     *bool                 `json:"has_wiki,omitempty"`
	HasIssues   *string               `json:"has_issues,omitempty"`
	Parent      *Repository           `json:"parent,omitempty"`
	Owner       *User                 `json:"owner,omitempty"`
}

// List returns a paginated list of all the forks of the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/forks#get
func (f *ForksService) List(owner, repoSlug string, opts ...interface{}) (*Repositories, *Response, error) {
	result := new(Repositories)
	urlStr := f.client.requestURL("/repositories/%s/%s/forks", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := f.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Create creates a new fork of the specified repository.
//
// By default, forks are created under the authenticated user's account with the same name and slug of the original repository.
//
// The 'owner' & 'repoSlug' parameters represent the repository you want to fork into your account.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/forks#post
func (f *ForksService) Create(owner, repoSlug string, fo *ForkRequest) (*Repository, *Response, error) {
	result := new(Repository)
	urlStr := f.client.requestURL("/repositories/%s/%s/forks", owner, repoSlug)
	response, err := f.client.execute("POST", urlStr, result, fo)

	return result, response, err
}
