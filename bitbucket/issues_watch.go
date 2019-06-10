package bitbucket

// IsAuthUserWatching check whether the authenticated user is watching the specified issue.
//
// A 204 status code indicates that the user is watching this issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/watch#get
func (i *IssuesService) IsAuthUserWatching(owner, repoSlug string, id int64) (bool, *Response, error) {
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/watch", owner, repoSlug, id)
	response, err := i.client.execute("GET", urlStr, nil, nil)

	hasVoted := false
	if response.StatusCode == 204 {
		hasVoted = true
	}

	return hasVoted, response, err
}

// WatchIssue starts watching the specified issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/watch#put
func (i *IssuesService) WatchIssue(owner, repoSlug string, id int64) (*Response, error) {
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/watch", owner, repoSlug, id)
	response, err := i.client.execute("PUT", urlStr, nil, nil)

	return response, err
}

// StopWatchingIssue stops watching the specified issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/watch#delete
func (i *IssuesService) StopWatchingIssue(owner, repoSlug string, id int64) (*Response, error) {
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/watch", owner, repoSlug, id)
	response, err := i.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
