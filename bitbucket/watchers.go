package bitbucket

// WatchersService handles communication with the watchers related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/watchers
type WatchersService service

// List returns a paginated list of all the watchers on the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/watchers#get
func (w *WatchersService) List(owner, repoSlug string, opts ...interface{}) (*Users, *Response, error) {
	results := new(Users)
	urlStr := w.client.requestUrl("/repositories/%s/%s/watchers", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := w.client.execute("GET", urlStr, results, nil)

	return results, response, err
}
