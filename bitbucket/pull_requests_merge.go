package bitbucket

// MergePR merges the pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/merge#post
func (p *PullRequestsService) MergePR(owner, repoSlug string, pullRequestId int64) (*PullRequest, *Response, error) {
	result := new(PullRequest)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/merge", owner, repoSlug, pullRequestId)
	response, err := p.client.execute("POST", urlStr, result, nil)

	return result, response, err
}
