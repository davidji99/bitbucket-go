package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// BranchRestrictionsService handles communication with the branch restrictions related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branch-restrictions
type BranchRestrictionsService service

// BranchRestrictions represent a collection of branch restrictions.
type BranchRestrictions struct {
	PaginationInfo

	Values []*BranchRestriction `json:"values,omitempty"`
}

// BranchRestriction represents a Bitbucket repository branch restriction.
type BranchRestriction struct {
	ID              *int64   `json:"id,omitempty"`
	Kind            *string  `json:"kind,omitempty"`
	Users           []*User  `json:"users,omitempty"`
	Pattern         *string  `json:"pattern,omitempty"`
	Value           *int64   `json:"value,omitempty"`
	BranchMatchKind *string  `json:"branch_match_kind,omitempty"`
	Type            *string  `json:"type,omitempty"`
	Links           *BRLinks `json:"links,omitempty"`
}

// BRLinks represents the "links" object in a Bitbucket branch restriction.
type BRLinks struct {
	Self *Link `json:"self,omitempty"`
}

// BRRequest represents a request to create/update a branch restriction.
type BRRequest struct {
	Kind            *string `json:"kind,omitempty"`
	BranchMatchKind *string `json:"branch_match_kind,omitempty"`
	BranchType      *string `json:"branch_type,omitempty"`
	Pattern         *string `json:"pattern,omitempty"`
}

// BranchRestrictionListOpts represents the query parameters available to listing all branch restrictions.
type BranchRestrictionListOpts struct {
	// Branch restrictions of this type
	Kind string `url:"kind,omitempty"`

	// Branch restrictions applied to branches of this pattern
	Pattern string `url:"pattern,omitempty"`
}

// List returns a paginated list of all branch restrictions on the repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branch-restrictions#get
func (br *BranchRestrictionsService) List(owner, repoSlug string, opts ...interface{}) (*BranchRestrictions, *simpleresty.Response, error) {
	result := new(BranchRestrictions)
	urlStr, urlStrErr := br.client.http.RequestURLWithQueryParams(fmt.Sprintf("/repositories/%s/%s/branch-restrictions", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := br.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// Get Returns a specific branch restriction.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branch-restrictions/%7Bid%7D#get
func (br *BranchRestrictionsService) Get(owner, repoSlug, id string, opts ...interface{}) (*BranchRestriction, *simpleresty.Response, error) {
	result := new(BranchRestriction)
	urlStr, urlStrErr := br.client.http.RequestURLWithQueryParams(fmt.Sprintf("/repositories/%s/%s/branch-restrictions/%s", owner, repoSlug, id), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := br.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// Update updates an existing branch restriction rule.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branch-restrictions/%7Bid%7D#put
func (br *BranchRestrictionsService) Update(owner, repoSlug string, brID int64, bo *BRRequest) (*BranchRestriction, *simpleresty.Response, error) {
	result := new(BranchRestriction)
	urlStr := br.client.http.RequestURL("/repositories/%s/%s/branch-restrictions/%v", owner, repoSlug, brID)

	response, err := br.client.http.Put(urlStr, result, bo)

	return result, response, err
}

// Create creates a new branch restriction rule for a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branch-restrictions#post
func (br *BranchRestrictionsService) Create(owner, repoSlug string, bo *BRRequest) (*BranchRestriction, *simpleresty.Response, error) {
	result := new(BranchRestriction)
	urlStr := br.client.http.RequestURL("/repositories/%s/%s/branch-restrictions", owner, repoSlug)

	response, err := br.client.http.Post(urlStr, result, bo)

	return result, response, err
}

// Delete an existing branch restriction rule.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branch-restrictions/%7Bid%7D#delete
func (br *BranchRestrictionsService) Delete(owner, repoSlug string, brID int64) (*simpleresty.Response, error) {
	urlStr := br.client.http.RequestURL("/repositories/%s/%s/branch-restrictions/%v", owner, repoSlug, brID)

	response, err := br.client.http.Delete(urlStr, nil, nil)

	return response, err
}
