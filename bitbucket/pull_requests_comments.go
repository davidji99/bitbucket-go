package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// PRComments represents a collection of a PR's comments.
type PRComments struct {
	PaginationInfo

	Values []*PRComment `json:"values,omitempty"`
}

// PRComment represents a pull request comment.
type PRComment struct {
	Comment

	PullRequest *PullRequest `json:"pullrequest,omitempty"`
	Deleted     *bool        `json:"deleted,omitempty"`
}

// PRCommentRequest represents a request to create or update a pull request comment.
type PRCommentRequest struct {
	Content *Content `json:"content,omitempty"`
}

// ListComments returns a paginated list of the pull request's comments.
//
// This includes both global, inline comments and replies.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#get
func (p *PullRequestsService) ListComments(owner, repoSlug string, pullRequestID int64, opts ...interface{}) (*PRComments, *simpleresty.Response, error) {
	result := new(PRComments)
	urlStr, urlStrErr := p.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/pullrequests/%v/comments", owner, repoSlug, pullRequestID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := p.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// CreateComment creates a new pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#post
func (p *PullRequestsService) CreateComment(owner, repoSlug string, pullRequestID int64, po *PRCommentRequest) (*PRComment, *simpleresty.Response, error) {
	result := new(PRComment)
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/comments", owner, repoSlug, pullRequestID)
	response, err := p.client.http.Post(urlStr, result, po)

	return result, response, err
}

// GetComment returns a specific pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments/%7Bcomment_id%7D#get
func (p *PullRequestsService) GetComment(owner, repoSlug string, prID, cID int64, opts ...interface{}) (*PRComment, *simpleresty.Response, error) {
	result := new(PRComment)
	urlStr, urlStrErr := p.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/pullrequests/%v/comments/%v", owner, repoSlug, prID, cID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := p.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// UpdateComment updates a specific pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#put
func (p *PullRequestsService) UpdateComment(owner, repoSlug string, prID, cID int64, po *PRCommentRequest) (*PRComment, *simpleresty.Response, error) {
	result := new(PRComment)
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/comments/%v", owner, repoSlug, prID, cID)
	response, err := p.client.http.Put(urlStr, result, po)

	return result, response, err
}

// DeleteComment updates a specific pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#delete
func (p *PullRequestsService) DeleteComment(owner, repoSlug string, prID, cID int64) (*simpleresty.Response, error) {
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/comments/%v", owner, repoSlug, prID, cID)
	response, err := p.client.http.Delete(urlStr, nil, nil)

	return response, err
}
