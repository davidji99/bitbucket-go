package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

// CommitComments represent a collection of a commit's comments.
type CommitComments struct {
	PaginationInfo

	Values []*CommitComment `json:"values,omitempty"`
}

// CommitComment represents a commit comment.
type CommitComment struct {
	ID        *int64         `json:"id,omitempty"`
	Links     *CCLinks       `json:"links,omitempty"`
	Deleted   *bool          `json:"deleted,omitempty"`
	Content   *Content       `json:"content,omitempty"`
	CreatedOn *time.Time     `json:"created_on,omitempty"`
	User      *User          `json:"user,omitempty"`
	Commit    *Commit        `json:"commit,omitempty"`
	UpdatedOn *time.Time     `json:"updated_on,omitempty"`
	Type      *string        `json:"type,omitempty"`
	Parent    *CommitComment `json:"parent,omitempty"`
}

// CCLinks represents the "links" object in a Bitbucket commit comment.
type CCLinks struct {
	Self *Link `json:"self,omitempty"`
	HTML *Link `json:"html,omitempty"`
}

// CommitCommentRequest represents a new commit comment.
type CommitCommentRequest struct {
	Content       *Content       `json:"content,omitempty"`
	ParentComment *CommitComment `json:"parent,omitempty"`
}

// ListComments returns the commit's comments.
//
// This includes both global and inline comments.
// The default sorting is oldest to newest and can be overridden with the sort query parameter.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/comments#get
func (c *CommitService) ListComments(owner, repoSlug, sha string, opts ...interface{}) (*CommitComments, *simpleresty.Response, error) {
	results := new(CommitComments)
	urlStr, urlStrErr := c.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/commit/%s/comments", owner, repoSlug, sha), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := c.client.http.Get(urlStr, results, nil)

	return results, response, err
}

// CreateComment creates new comment on the specified commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/comments#post
func (c *CommitService) CreateComment(owner, repoSlug, sha string, co *CommitCommentRequest) (*CommitComment, *simpleresty.Response, error) {
	results := new(CommitComment)
	urlStr := c.client.http.RequestURL("/repositories/%s/%s/commit/%s/comments", owner, repoSlug, sha)
	response, err := c.client.http.Post(urlStr, results, co)

	return results, response, err
}

// GetComment returns the specified commit comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/comments/%7Bcomment_id%7D#get
func (c *CommitService) GetComment(owner, repoSlug, sha string, cID int64, opts ...interface{}) (*CommitComment, *simpleresty.Response, error) {
	results := new(CommitComment)
	urlStr, urlStrErr := c.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/commit/%s/comments/%v", owner, repoSlug, sha, cID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := c.client.http.Get(urlStr, results, nil)

	return results, response, err
}
