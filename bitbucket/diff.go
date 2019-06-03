package bitbucket

import (
	"bytes"
)

// DiffService handles communication with the diff related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/diff
type DiffService service

type DiffGetOpts struct {
	Context          *int    `url:"context,omitempty"`
	Path             *string `url:"path,omitempty"`
	IgnoreWhitespace *string `url:"ignore_whitespace,omitempty"`
	Binary           *string `url:"binary,omitempty"`
}

// Get produces a raw, git-style diff for either a single commit (diffed against its first parent),
// or a revspec of 2 commits (e.g. 3a8b42..9ff173 where the first commit represents the source and the second commit the destination).
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/diff/%7Bspec%7D#get
func (d *DiffService) Get(owner, repoSlug, spec string, opts *DiffGetOpts) (interface{}, *Response, error) {
	urlStr := d.client.requestUrl("/repositories/%s/%s/diff/%s", owner, repoSlug, spec)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	req, reqErr := d.client.newRequest("GET", urlStr, nil, nil)
	if reqErr != nil {
		return nil, nil, reqErr
	}

	var buff bytes.Buffer
	response, err := d.client.doRequest(req, &buff, false)

	return buff.String(), response, err
}
