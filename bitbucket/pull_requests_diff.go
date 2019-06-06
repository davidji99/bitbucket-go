package bitbucket

import "bytes"

// GetDiffRaw produces a raw, git-style diff for the pull requests
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/diff#get
func (p *PullRequestsService) GetDiffRaw(owner, repoSlug string, pid int64) (*bytes.Buffer, *Response, error) {
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/diff", owner, repoSlug, pid)

	req, reqErr := p.client.newRequest("GET", urlStr, nil, nil)
	if reqErr != nil {
		return nil, nil, reqErr
	}

	var buff bytes.Buffer
	response, err := p.client.doRequest(req, &buff, false)

	return &buff, response, err
}

// GetDiff returns the diff stat for the specified pull request.
//
// Diff stat responses contain a record for every path modified by the commit and lists the number of lines added and removed for each file.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/diffstat#get
func (p *PullRequestsService) GetDiff(owner, repoSlug string, pid int64, opts ...interface{}) (*Diffs, *Response, error) {
	result := new(Diffs)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/diffstat", owner, repoSlug, pid)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
