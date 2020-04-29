package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// ListStatuses returns all statuses (e.g. build results) for the given pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/statuses#get
func (p *PullRequestsService) ListStatuses(owner, repoSlug string, pid int64, opts ...interface{}) (*CommitStatuses, *simpleresty.Response, error) {
	result := new(CommitStatuses)
	urlStr, urlStrErr := p.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/pullrequests/%v/statuses", owner, repoSlug, pid), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := p.client.http.Get(urlStr, result, nil)

	return result, response, err
}
