package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

// CommitService handles communication with the commit related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit
type CommitService service

// Commit represents a git commit in a Bitbucket repository.
type Commit struct {
	Rendered     *CommitMessageContent `json:"rendered,omitempty"`
	Hash         *string               `json:"hash,omitempty"`
	Repository   *Repository           `json:"repository,omitempty"`
	Links        *CommitLinks          `json:"links,omitempty"`
	Author       *User                 `json:"author,omitempty"`
	Summary      *Content              `json:"summary,omitempty"`
	Participants []*Participant        `json:"participants,omitempty"`
	Parents      []*Commit             `json:"parents,omitempty"`
	Date         *time.Time            `json:"date,omitempty"`
	Message      *string               `json:"message,omitempty"`
	Type         *string               `json:"type,omitempty"`
}

// CommitMessageContent represents the commit's message.
type CommitMessageContent struct {
	Message *Content `json:"message,omitempty"`
}

// Participant represents a user that interacted with a Bitbucket resource.
type Participant struct {
	Role           *string    `json:"role,omitempty"`
	ParticipatedOn *time.Time `json:"participated_on,omitempty"`
	Type           *string    `json:"type,omitempty"`
	Approved       *bool      `json:"approved,omitempty"`
	User           *User      `json:"user,omitempty"`
}

// CommitLinks represents the "links" object in a Bitbucket commit.
type CommitLinks struct {
	Self     *Link `json:"self,omitempty"`
	Comment  *Link `json:"comment,omitempty"`
	HTML     *Link `json:"html,omitempty"`
	Diff     *Link `json:"diff,omitempty"`
	Approve  *Link `json:"approve,omitempty"`
	Statuses *Link `json:"statuses,omitempty"`
}

// Get return the specified commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D#get
func (c *CommitService) Get(owner, repoSlug, sha string, opts ...interface{}) (*Commit, *simpleresty.Response, error) {
	results := new(Commit)
	urlStr, urlStrErr := c.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/commit/%s", owner, repoSlug, sha), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := c.client.http.Get(urlStr, results, nil)

	return results, response, err
}
