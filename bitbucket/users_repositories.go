package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// ListRepositories fetches all repositories owned by a user/team.
//
// This includes private repositories, but filtered down to the ones that the calling user has access to.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/repositories#get
func (u *UsersService) ListRepositories(userID string, opts ...interface{}) (*Repositories, *simpleresty.Response, error) {
	repos := new(Repositories)
	urlStr, urlStrErr := u.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/users/%s/repositories", userID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := u.client.http.Get(urlStr, repos, nil)

	return repos, response, err
}
