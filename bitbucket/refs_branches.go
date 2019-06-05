package bitbucket

// RefRequest represents a request to create a new branch.
type RefRequest struct {
	Name   *string `json:"name,omitempty"`
	Target struct {
		Hash *string `json:"hash,omitempty"`
	} `json:"target,omitempty"`
}

// ListBranches returns a list of all open branches within the specified repository.
// Results will be in the order the source control manager returns them.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/branches#get
func (r *RefsService) ListBranches(owner, repoSlug string, opts ...interface{}) (*Refs, *Response, error) {
	result := new(Refs)
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/branches", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// CreateBranch creates a new branch in the specified repository.
//
// The branch name should not include any prefixes (e.g. refs/heads).
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/branches#post
func (r *RefsService) CreateBranch(owner, repoSlug string, ro *RefRequest, opts ...interface{}) (*Ref, *Response, error) {
	result := new(Ref)
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/branches", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("POST", urlStr, result, ro)

	return result, response, err
}

// GetBranch returns a branch object within the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/branches/%7Bname%7D#get
func (r *RefsService) GetBranch(owner, repoSlug, name string, opts ...interface{}) (*Ref, *Response, error) {
	result := new(Ref)
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/branches/%s", owner, repoSlug, name)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// DeleteBranch deletes a branch in the specified repository.
//
// The main branch is not allowed to be deleted and will return a 400 response.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/branches/%7Bname%7D#delete
func (r *RefsService) DeleteBranch(owner, repoSlug, name string, opts ...interface{}) (*Response, error) {
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/branches/%s", owner, repoSlug, name)
	response, err := r.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
