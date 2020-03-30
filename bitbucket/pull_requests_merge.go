package bitbucket

import "github.com/davidji99/simpleresty"

// MergePR merges the pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/merge#post
func (p *PullRequestsService) MergePR(owner, repoSlug string, pullRequestID int64) (*PullRequest, *simpleresty.Response, error) {
	result := new(PullRequest)
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/merge", owner, repoSlug, pullRequestID)
	response, err := p.client.http.Post(urlStr, result, nil)

	return result, response, err
}
