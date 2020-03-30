package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

// RefsService handles communication with the refs related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs
type RefsService service

// Refs represents a collection of refs.
type Refs struct {
	PaginationInfo

	Values []*Ref `json:"values,omitempty"`
}

// Ref represents the branches and tags in a repository.
type Ref struct {
	Heads                []*Commit  `json:"heads,omitempty"`
	Date                 *time.Time `json:"date,omitempty"`
	Message              *string    `json:"message,omitempty"`
	Name                 *string    `json:"name,omitempty"`
	Links                *RefLinks  `json:"links,omitempty"`
	DefaultMergeStrategy *string    `json:"default_merge_strategy,omitempty"`
	MergeStrategies      []*string  `json:"merge_strategies,omitempty"`
	Type                 *string    `json:"type,omitempty"`
	Target               *Commit    `json:"target,omitempty"`
}

// RefLinks represents the "links" object in a ref.
type RefLinks struct {
	Commits *Link `json:"commits,omitempty"`
	Self    *Link `json:"self,omitempty"`
	HTML    *Link `json:"html,omitempty"`
}

// ListAll returns the branches and tags in the repository.
//
// By default, results will be in the order the underlying source control system returns them and identical
// to the ordering one sees when running "$ git show-ref".
// Note that this follows simple lexical ordering of the ref names.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs#get
func (r *RefsService) ListAll(owner, repoSlug string, opts ...interface{}) (*Refs, *simpleresty.Response, error) {
	result := new(Refs)
	urlStr, urlStrErr := r.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/refs", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := r.client.http.Get(urlStr, result, nil)

	return result, response, err
}
