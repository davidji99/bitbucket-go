package bitbucket

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

// ListComments returns a paginated list of the pull request's comments.
//
// This includes both global, inline comments and replies.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#get
func (p *PullRequestsService) ListComments(owner, repoSlug string, pullRequestId int64, opts ...interface{}) (*PRComments, *Response, error) {
	result := new(PRComments)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/comments", owner, repoSlug, pullRequestId)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// CreateComment creates a new pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#post
func (p *PullRequestsService) CreateComment(owner, repoSlug string, pullRequestId int64, po *Content) (*PRComment, *Response, error) {
	result := new(PRComment)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/comments", owner, repoSlug, pullRequestId)
	response, err := p.client.execute("POST", urlStr, result, po)

	return result, response, err
}

// GetComment returns a specific pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments/%7Bcomment_id%7D#get
func (p *PullRequestsService) GetComment(owner, repoSlug string, prID, cID int64, opts ...interface{}) (*PRComment, *Response, error) {
	result := new(PRComment)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/comments/%v", owner, repoSlug, prID, cID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// UpdateComment updates a specific pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#put
func (p *PullRequestsService) UpdateComment(owner, repoSlug string, prID, cID int64, po *Content) (*PRComment, *Response, error) {
	result := new(PRComment)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/comments/%v", owner, repoSlug, prID, cID)
	response, err := p.client.execute("PUT", urlStr, result, po)

	return result, response, err
}

// DeleteComment updates a specific pull request comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/comments#delete
func (p *PullRequestsService) DeleteComment(owner, repoSlug string, prID, cID int64) (*Response, error) {
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/comments/%v", owner, repoSlug, prID, cID)
	response, err := p.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
