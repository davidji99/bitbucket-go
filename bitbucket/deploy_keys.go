package bitbucket

import "time"

// DeployKeysService handles communication with the deploy keys related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/deploy-keys
type DeployKeysService service

// DeployKeys represents a collection of deploy keys.
type DeployKeys struct {
	PaginationInfo

	Values []*DeployKey `json:"values,omitempty"`
}

// DeployKey represents a deploy key aka access key on a repository.
type DeployKey struct {
	ID         *int64          `json:"id,omitempty"`
	Key        *string         `json:"key,omitempty"`
	Label      *string         `json:"label,omitempty"`
	Type       *string         `json:"type,omitempty"`
	CreatedOn  *time.Time      `json:"created_on,omitempty"`
	Repository *Repository     `json:"repository,omitempty"`
	Links      *DeployKeyLinks `json:"links,omitempty"`
	LastUsed   *time.Time      `json:"last_used,omitempty"`
	Comment    *string         `json:"comment,omitempty"`
	AddedOn    *time.Time      `json:"added_on,omitempty"`
}

// DeployKeyLinks represents the "links" object in a Bitbucket deploy key.
type DeployKeyLinks struct {
	Self *Link `json:"self,omitempty"`
}

// DeployKeyRequest represents a request to create/update a deploy/access key.
type DeployKeyRequest struct {
	Key   *string `json:"key,omitempty"`
	Label *string `json:"label,omitempty"`
}

// List returns all deploy-keys belonging to a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/deploy-keys#get
func (dk *DeployKeysService) List(owner, repoSlug string, opts ...interface{}) (*DeployKeys, *Response, error) {
	result := new(DeployKeys)
	urlStr := dk.client.requestURL("/repositories/%s/%s/deploy-keys", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := dk.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Add creates a new deploy key in a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/deploy-keys#post
func (dk *DeployKeysService) Add(owner, repoSlug string, do *DeployKeyRequest) (*DeployKey, *Response, error) {
	result := new(DeployKey)
	urlStr := dk.client.requestURL("/repositories/%s/%s/deploy-keys", owner, repoSlug)
	response, err := dk.client.execute("POST", urlStr, result, do)

	return result, response, err
}

// Get returns the deploy key belonging to a specific key.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/deploy-keys/%7Bkey_id%7D#get
func (dk *DeployKeysService) Get(owner, repoSlug string, keyID int64, opts ...interface{}) (*DeployKey, *Response, error) {
	result := new(DeployKey)
	urlStr := dk.client.requestURL("/repositories/%s/%s/deploy-keys/%v", owner, repoSlug, keyID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := dk.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Update modifies an existing key's label and/or comment.
// The same key needs to be passed in but the comment and label can change.
//
// For security reasons, you can't modify the contents of an access key. To update, delete and re-add the key.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/deploy-keys/%7Bkey_id%7D#put
func (dk *DeployKeysService) Update(owner, repoSlug string, keyID int64, do *DeployKeyRequest) (*DeployKey, *Response, error) {
	result := new(DeployKey)
	urlStr := dk.client.requestURL("/repositories/%s/%s/deploy-keys/%v", owner, repoSlug, keyID)
	response, err := dk.client.execute("PUT", urlStr, result, do)

	return result, response, err
}

// Remove deletes a deploy key from a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/deploy-keys/%7Bkey_id%7D#delete
func (dk *DeployKeysService) Remove(owner, repoSlug string, keyID int64) (*Response, error) {
	urlStr := dk.client.requestURL("/repositories/%s/%s/deploy-keys/%v", owner, repoSlug, keyID)
	response, err := dk.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
