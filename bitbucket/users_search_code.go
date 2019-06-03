package bitbucket

// SearchCode searches for code in the repositories of the specified user.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/search/code#get
func (u *UsersService) SearchCode(userID string, opts *CodeSearchOpts) (*SearchCodeResults, *Response, error) {
	results := new(SearchCodeResults)
	urlStr := u.client.requestUrl("/users/%s/search/code", userID)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := u.client.execute("GET", urlStr, results, nil)

	return results, response, err
}
