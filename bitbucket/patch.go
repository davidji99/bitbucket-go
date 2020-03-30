package bitbucket

import (
	"bytes"
	"github.com/davidji99/simpleresty"
)

// PatchService handles communication with the patch related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/patch
type PatchService service

// GetRaw produces a raw patch for a single commit (diffed against its first parent),
// or a patch-series for a revspec of 2 commits (e.g. 3a8b42..9ff173 where the first commit
// represents the source and the second commit the destination).
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/patch/%7Bspec%7D#get
func (p *PatchService) GetRaw(owner, repoSlug, spec string) (*bytes.Buffer, *simpleresty.Response, error) {
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/patch/%s", owner, repoSlug, spec)

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
