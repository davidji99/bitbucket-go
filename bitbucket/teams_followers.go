package bitbucket

// ListFollowers returns the list of accounts that are following this team.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/followers#get
func (t *TeamsService) ListFollowers(teamUsername string, opts ...interface{}) (*Users, *Response, error) {
	result := new(Users)
	urlStr := t.client.requestURL("/teams/%s/followers", teamUsername)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
