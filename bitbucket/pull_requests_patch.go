package bitbucket

import "bytes"

// GetPatchRaw produces a raw patch for the specified pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/patch#get
func (p *PullRequestsService) GetPatchRaw(owner, repoSlug string, pid int64) (*bytes.Buffer, *Response, error) {
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/patch", owner, repoSlug, pid)

	req, reqErr := p.client.newRequest("GET", urlStr, nil, nil)
	if reqErr != nil {
		return nil, nil, reqErr
	}

	var buff bytes.Buffer
	response, err := p.client.doRequest(req, &buff, false)

	return &buff, response, err
}
