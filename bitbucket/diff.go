package bitbucket

import (
	"bytes"
)

// DiffService handles communication with the diff related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/diff
type DiffService service

// Diffs represents a collection of diffs.
type Diffs struct {
	PaginationInfo

	Values []*Diff `json:"values,omitempty"`
}

// Diff represents a code diff on Bitbucket.
type Diff struct {
	Status       *string   `json:"state,omitempty"`
	Old          *CodeFile `json:"old,omitempty"`
	New          *CodeFile `json:"new,omitempty"`
	LinesRemoved *int64    `json:"lines_removed,omitempty"`
	LinesAdded   *int64    `json:"lines_added,omitempty"`
	Type         *string   `json:"type,omitempty"`
}

// DiffGetOpts represents the query parameters available when getting the raw of a diff.
type DiffGetOpts struct {
	Context          *int    `url:"context,omitempty"`
	Path             *string `url:"path,omitempty"`
	IgnoreWhitespace *string `url:"ignore_whitespace,omitempty"`
	Binary           *string `url:"binary,omitempty"`
}

// GetRaw produces a raw, git-style diff for either a single commit (diffed against its first parent),
// or a revspec of 2 commits (e.g. 3a8b42..9ff173 where the first commit represents the source and the second commit the destination).
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/diff/%7Bspec%7D#get
func (d *DiffService) GetRaw(owner, repoSlug, spec string) (*bytes.Buffer, *Response, error) {
	urlStr := d.client.requestUrl("/repositories/%s/%s/diff/%s", owner, repoSlug, spec)

	req, reqErr := d.client.newRequest("GET", urlStr, nil, nil)
	if reqErr != nil {
		return nil, nil, reqErr
	}

	var buff bytes.Buffer
	response, err := d.client.doRequest(req, &buff, false)

	return &buff, response, err
}

// Get returns the diff stat for the specified commit.
//
// Diff stat responses contain a record for every path modified by the commit and lists the number of lines added and removed for each file.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/diffstat/%7Bspec%7D#get
func (d *DiffService) Get(owner, repoSlug, spec string, opts ...interface{}) (*Diffs, *Response, error) {
	result := new(Diffs)
	urlStr := d.client.requestUrl("/repositories/%s/%s/diffstat/%s", owner, repoSlug, spec)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := d.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
