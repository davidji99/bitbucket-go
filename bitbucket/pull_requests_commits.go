package bitbucket

// ListCommits returns a paginated list of a pull request's commits.
//
// These are the commits that are being merged into the destination branch when the pull requests gets accepted.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/commits#get
func (p *PullRequestsService) ListCommits(owner, repoSlug string, pullRequestID int64, opts ...interface{}) (*Commits, *Response, error) {
	result := new(Commits)
	urlStr := p.client.requestURL("/repositories/%s/%s/pullrequests/%v/commits", owner, repoSlug, pullRequestID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
