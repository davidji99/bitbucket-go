package bitbucket

import "time"

// IssueComments represents a collection of issue comments.
type IssueComments struct {
	PaginationInfo

	Values []*IssueChange `json:"values,omitempty"`
}

// IssueComment represents a comment on an issue.
type IssueComment struct {
	ID        *int64             `json:"id,omitempty"`
	Type      *string            `json:"type,omitempty"`
	Links     *IssueCommentLinks `json:"links,omitempty"`
	Issue     *Issue             `json:"issue,omitempty"`
	Content   *Content           `json:"content,omitempty"`
	CreatedOn *time.Time         `json:"created_on,omitempty"`
	UpdatedOn *time.Time         `json:"updated_on,omitempty"`
	User      *User              `json:"user,omitempty"`
}

// IssueCommentLinks represents the "links" object in a Bitbucket issue comment.
type IssueCommentLinks struct {
	Self *Link `json:"self,omitempty"`
	HTML *Link `json:"html,omitempty"`
}

// IssueCommentRequest represents a request to create/update an issue comment.
type IssueCommentRequest struct {
	Content *Content `json:"content,omitempty"`
}

// ListComments returns a paginated list of all comments that were made on the specified issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/comments#get
func (i *IssuesService) ListComments(owner, repoSlug string, id int64, opts ...interface{}) (*IssueComments, *Response, error) {
	result := new(IssueComments)
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/comments", owner, repoSlug, id)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := i.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// CreateComment creates a new issue comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/comments#post
func (i *IssuesService) CreateComment(owner, repoSlug string, id int64, io *IssueCommentRequest) (*IssueComment, *Response, error) {
	result := new(IssueComment)
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/comments", owner, repoSlug, id)
	response, err := i.client.execute("POST", urlStr, result, io)

	return result, response, err
}

// GetComment returns the specified issue comment object.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/comments/%7Bcomment_id%7D#get
func (i *IssuesService) GetComment(owner, repoSlug string, id, commentID int64, opts ...interface{}) (*IssueComment, *Response, error) {
	result := new(IssueComment)
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/comments/%v", owner, repoSlug, id, commentID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := i.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// UpdateComment updates an existing issue comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/comments/%7Bcomment_id%7D#put
func (i *IssuesService) UpdateComment(owner, repoSlug string, id, commentID int64, io *IssueCommentRequest) (*IssueComment, *Response, error) {
	result := new(IssueComment)
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/comments/%v", owner, repoSlug, id, commentID)
	response, err := i.client.execute("PUT", urlStr, result, io)

	return result, response, err
}

// DeleteComment deletes an existing issue comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/comments/%7Bcomment_id%7D#delete
func (i *IssuesService) DeleteComment(owner, repoSlug string, id, commentID int64) (*Response, error) {
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/comments/%v", owner, repoSlug, id, commentID)
	response, err := i.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
