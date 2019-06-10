package bitbucket

// CommitsService handles communication with the commits related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits
type CommitsService service

// Commits represent a collection of commits.
type Commits struct {
	PaginationInfo

	Values []*Commit `json:"values,omitempty"`
}

// List all commits for a given repository.
//
// Supports filtering by passing in a non-URI encoded query string. Refer to the API docs below.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits#get
func (c *CommitsService) List(owner, repoSlug string, opts ...interface{}) (*Commits, *Response, error) {
	result := new(Commits)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commits", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := c.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Get a commit revision.
//
// NOTE: Identical to GET /repositories/{username}/{repo_slug}/commits,
// except that POST allows clients to place the include and exclude parameters in the request body to avoid URL length issues.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits/%7Brevision%7D#post
func (c *CommitsService) ListSafe(owner, repoSlug string, opts ...interface{}) (*Commits, *Response, error) {
	// TODO: The name of this function isn't that great...feel free to suggest a new name and perhaps how this endpoint is suppose to work.
	result := new(Commits)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commits", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := c.client.execute("POST", urlStr, result, nil)

	return result, response, err
}

// GetRevision returns a commit revision. The results can return a collection of commits. Does not support any query parameters.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits/%7Brevision%7D#get
func (c *CommitsService) GetRevision(owner, repoSlug, revision string, opts ...interface{}) (*Commits, *Response, error) {
	commits := new(Commits)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commits/%s", owner, repoSlug, revision)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := c.client.execute("GET", urlStr, commits, nil)

	return commits, response, err
}
