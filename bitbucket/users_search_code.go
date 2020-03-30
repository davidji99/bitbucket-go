package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// SearchCode searches for code in the repositories of the specified user.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/search/code#get
func (u *UsersService) SearchCode(userID string, opts ...interface{}) (*SearchCodeResults, *simpleresty.Response, error) {
	results := new(SearchCodeResults)
	urlStr, urlStrErr := u.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/users/%s/search/code", userID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := u.client.http.Get(urlStr, results, nil)

	return results, response, err
}
