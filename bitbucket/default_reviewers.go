package bitbucket

// DefaultReviewersService handles communication with the default reviewers related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/default-reviewers
type DefaultReviewersService service

// List returns the repository's default reviewers.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/default-reviewers#get
func (dr *DefaultReviewersService) List(owner, repoSlug string, opts ...interface{}) (*Users, *Response, error) {
	result := new(Users)
	urlStr := dr.client.requestURL("/repositories/%s/%s/default-reviewers", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := dr.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Get returns the specified reviewer.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// This can be used to test whether a user is among the repository's default reviewers list. A 404 indicates that that specified user is not a default reviewer.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/default-reviewers/%7Btarget_username%7D#get
func (dr *DefaultReviewersService) Get(owner, repoSlug, userID string, opts ...interface{}) (*User, *Response, error) {
	result := new(User)
	urlStr := dr.client.requestURL("/repositories/%s/%s/default-reviewers/%s", owner, repoSlug, userID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := dr.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Add adds the specified user to the repository's list of default reviewers.
// This method is idempotent. Adding a user a second time has no effect.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/default-reviewers/%7Btarget_username%7D#put
func (dr *DefaultReviewersService) Add(owner, repoSlug, userID string) (*User, *Response, error) {
	result := new(User)
	urlStr := dr.client.requestURL("/repositories/%s/%s/default-reviewers/%s", owner, repoSlug, userID)
	response, err := dr.client.execute("PUT", urlStr, result, nil)

	return result, response, err
}

// Remove removes a default reviewer from the repository.
// This method is idempotent. Removing a user a second time has no effect.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/default-reviewers/%7Btarget_username%7D#delete
func (dr *DefaultReviewersService) Remove(owner, repoSlug, userID string) (*Response, error) {
	urlStr := dr.client.requestURL("/repositories/%s/%s/default-reviewers/%s", owner, repoSlug, userID)
	response, err := dr.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
