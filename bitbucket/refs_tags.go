package bitbucket

// ListTags returns the tags in the repository.
// Results will be in the order the source control manager returns them.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/tags#get
func (r *RefsService) ListTags(owner, repoSlug string, opts ...interface{}) (*Refs, *Response, error) {
	result := new(Refs)
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/tags", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// CreateTag creates a new tag in the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/tags#post
func (r *RefsService) CreateTag(owner, repoSlug string, ro *RefRequest, opts ...interface{}) (*Ref, *Response, error) {
	result := new(Ref)
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/tags", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("POST", urlStr, result, ro)

	return result, response, err
}

// GetTag returns a tag object within the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/tags/%7Bname%7D#get
func (r *RefsService) GetTag(owner, repoSlug, name string, opts ...interface{}) (*Ref, *Response, error) {
	result := new(Ref)
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/tags/%s", owner, repoSlug, name)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := r.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// DeleteTag deletes a tag in the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/refs/tags/%7Bname%7D#delete
func (r *RefsService) DeleteTag(owner, repoSlug, name string, opts ...interface{}) (*Response, error) {
	urlStr := r.client.requestUrl("/repositories/%s/%s/refs/tags/%s", owner, repoSlug, name)
	response, err := r.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
