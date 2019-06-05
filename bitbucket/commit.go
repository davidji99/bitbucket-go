package bitbucket

import "time"

// CommitService handles communication with the commit related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit
type CommitService service

// Commits represents a git commit in a Bitbucket repository.
type Commit struct {
	Rendered struct {
		Message *BitbucketContent `json:"message,omitempty"`
	} `json:"rendered,omitempty"`
	Hash         *string              `json:"hash,omitempty"`
	Repository   *Repository          `json:"repository,omitempty"`
	Links        *CommitLinks         `json:"links,omitempty"`
	Author       *User                `json:"author,omitempty"`
	Summary      *BitbucketContent    `json:"summary,omitempty"`
	Participants []*CommitParticipant `json:"participants,omitempty"`
	Parents      []*Commit            `json:"parents,omitempty"`
	Date         *time.Time           `json:"date,omitempty"`
	Message      *string              `json:"message,omitempty"`
	Type         *string              `json:"type,omitempty"`
}

// CommitParticipant represents a user that interacted with a commit.
type CommitParticipant struct {
	Role           *string    `json:"role,omitempty"`
	ParticipatedOn *time.Time `json:"participated_on,omitempty"`
	User           *User      `json:"user,omitempty"`
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

// GetCommit return the specified commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D#get
func (c *CommitService) Get(owner, repoSlug, sha string, opts ...interface{}) (*Commit, *Response, error) {
	results := new(Commit)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commit/%s", owner, repoSlug, sha)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := c.client.execute("GET", urlStr, results, nil)

	return results, response, err
}
