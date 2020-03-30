package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// WatchersService handles communication with the watchers related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/watchers
type WatchersService service

// List returns a paginated list of all the watchers on the specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/watchers#get
func (w *WatchersService) List(owner, repoSlug string, opts ...interface{}) (*Users, *simpleresty.Response, error) {
	results := new(Users)
	urlStr, urlStrErr := w.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/watchers", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := w.client.http.Get(urlStr, results, nil)

	return results, response, err
}
