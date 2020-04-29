package bitbucket

import "github.com/davidji99/simpleresty"

// Approve approves the specified pull request as the authenticated user.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/approve#post
func (p *PullRequestsService) Approve(owner, repoSlug string, pullRequestID int64) (*Participant, *simpleresty.Response, error) {
	result := new(Participant)
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/approve", owner, repoSlug, pullRequestID)
	response, err := p.client.http.Post(urlStr, result, nil)

	return result, response, err
}

// RemoveApproval redact the authenticated user's approval of the specified pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/approve#delete
func (p *PullRequestsService) RemoveApproval(owner, repoSlug string, pullRequestID int64) (*simpleresty.Response, error) {
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/approve", owner, repoSlug, pullRequestID)
	response, err := p.client.http.Delete(urlStr, nil, nil)

	return response, err
}
