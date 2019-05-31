package bitbucket

import "time"

// CommitsService handles communication with the commit related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits
type CommitsService service

// Commits represent a collection of commits.
type Commits struct {
	Page     int    `json:"page,omitempty"`
	Next     string `json:"next,omitempty"`
	Pagelen  int    `json:"pagelen,omitempty"`
	Size     int    `json:"size,omitempty"`
	Previous string `json:"previous,omitempty"`

	Values []*Commit `json:"values,omitempty"`
}

// Commits represents a git commit in a Bitbucket repository.
type Commit struct {
	Rendered struct {
		Message *BitbucketContent `json:"message,omitempty"`
	} `json:"rendered,omitempty"`
	Hash         *string           `json:"hash,omitempty"`
	Repository   *Repository       `json:"repository,omitempty"`
	Links        *CommitLinks      `json:"links,omitempty"`
	Author       *User             `json:"author,omitempty"`
	Summary      *BitbucketContent `json:"summary,omitempty"`
	Participants []*User           `json:"participants,omitempty"`
	Parents      []*Commit         `json:"parents,omitempty"`
	Date         *time.Time        `json:"date,omitempty"`
	Message      *string           `json:"message,omitempty"`
	Type         *string           `json:"type,omitempty"`
}

// CommitLinks represents the "links" object in a Bitbucket commit.
type CommitLinks struct {
	Self     *BitbucketLink `json:"self,omitempty"`
	Comment  *BitbucketLink `json:"comment,omitempty"`
	HTML     *BitbucketLink `json:"html,omitempty"`
	Diff     *BitbucketLink `json:"diff,omitempty"`
	Approve  *BitbucketLink `json:"approve,omitempty"`
	Statuses *BitbucketLink `json:"statuses,omitempty"`
}

// List all commits for a given repository.
// Supports filtering by passing in a non-URI encoded query string. Refer to the API docs below.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits#get
func (c *CommitsService) List(owner, repoSlug string) (*Commits, *Response, error) {
	commits := new(Commits)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commits", owner, repoSlug)
	response, err := c.client.execute("GET", urlStr, &commits, nil)

	return commits, response, err
}

// Get a commit revision. The results can return a collection of commits. Does not support any query parameters.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits/%7Brevision%7D#get
func (c *CommitsService) Get(owner, repoSlug, revision string) (*Commits, *Response, error) {
	commits := new(Commits)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commits/%s", owner, repoSlug, revision)
	response, err := c.client.execute("GET", urlStr, commits, nil)

	return commits, response, err
}

// Get a commit revision.
// NOTE: Identical to GET /repositories/{username}/{repo_slug}/commits,
// except that POST allows clients to place the include and exclude parameters in the request body to avoid URL length issues.
// TODO: The name of this function isn't that great...feel free to suggest a new name and perhaps how this endpoint is suppose to work.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits/%7Brevision%7D#post
func (c *CommitsService) ListSafe(owner, repoSlug string) (*Commits, *Response, error) {
	commits := new(Commits)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commits", owner, repoSlug)
	response, err := c.client.execute("POST", urlStr, commits, nil)

	return commits, response, err
}
