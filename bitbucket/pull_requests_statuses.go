package bitbucket

// ListStatuses returns all statuses (e.g. build results) for the given pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/statuses#get
func (p *PullRequestsService) ListStatuses(owner, repoSlug string, pid int64, opts ...interface{}) (*CommitStatuses, *Response, error) {
	result := new(CommitStatuses)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/statuses", owner, repoSlug, pid)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
