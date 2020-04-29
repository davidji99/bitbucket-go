package bitbucket

import (
	"bytes"
	"github.com/davidji99/simpleresty"
)

// GetPatchRaw produces a raw patch for the specified pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/patch#get
func (p *PullRequestsService) GetPatchRaw(owner, repoSlug string, pid int64) (*bytes.Buffer, *simpleresty.Response, error) {
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/patch", owner, repoSlug, pid)

	var buff bytes.Buffer
	req := p.client.http.NewRequest()
	req.Method = simpleresty.GetMethod
	req.URL = urlStr
	req.Result = &buff

	response, reqErr := p.client.http.Dispatch(req)
	if reqErr != nil {
		return nil, nil, reqErr
	}

	return &buff, response, nil
}
