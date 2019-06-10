package bitbucket

// PullRequestReview represents a review on a pull request.
type PullRequestReview struct {
	CommitParticipant

	Approved *bool   `json:"approved,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// Approve approves the specified pull request as the authenticated user.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/approve#post
func (p *PullRequestsService) Approve(owner, repoSlug string, pullRequestID int64) (*PullRequestReview, *Response, error) {
	result := new(PullRequestReview)
	urlStr := p.client.requestURL("/repositories/%s/%s/pullrequests/%v/approve", owner, repoSlug, pullRequestID)
	response, err := p.client.execute("POST", urlStr, result, nil)

	return result, response, err
}

// RemoveApproval redact the authenticated user's approval of the specified pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/approve#delete
func (p *PullRequestsService) RemoveApproval(owner, repoSlug string, pullRequestID int64) (*Response, error) {
	urlStr := p.client.requestURL("/repositories/%s/%s/pullrequests/%v/approve", owner, repoSlug, pullRequestID)
	response, err := p.client.execute("POST", urlStr, nil, nil)

	return response, err
}
