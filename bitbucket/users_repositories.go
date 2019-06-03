package bitbucket

// ListRepositories fetches all repositories owned by a user/team.
//
// This includes private repositories, but filtered down to the ones that the calling user has access to.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/repositories#get
func (u *UsersService) ListRepositories(userID string, opts *ListPaginationOpts) (*Repositories, *Response, error) {
	repos := new(Repositories)
	urlStr := u.client.requestUrl("/users/%s/repositories", userID)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := u.client.execute("GET", urlStr, repos, nil)

	return repos, response, err
}
