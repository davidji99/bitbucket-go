package bitbucket

import (
	"bytes"
	"fmt"
	"github.com/davidji99/simpleresty"
)

// GetDiffRaw produces a raw, git-style diff for the pull requests
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/diff#get
func (p *PullRequestsService) GetDiffRaw(owner, repoSlug string, pid int64) (*bytes.Buffer, *simpleresty.Response, error) {
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/diff", owner, repoSlug, pid)

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

// GetDiff returns the diff stat for the specified pull request.
//
// Diff stat responses contain a record for every path modified by the commit and lists the number of lines added and removed for each file.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/diffstat#get
func (p *PullRequestsService) GetDiff(owner, repoSlug string, pid int64, opts ...interface{}) (*Diffs, *simpleresty.Response, error) {
	result := new(Diffs)
	urlStr, urlStrErr := p.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/pullrequests/%v/diffstat", owner, repoSlug, pid), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := p.client.http.Get(urlStr, result, nil)

	return result, response, err
}
