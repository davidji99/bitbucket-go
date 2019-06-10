package bitbucket

// ListFollowing returns the list of accounts this team is following.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/following#get
func (t *TeamsService) ListFollowing(teamUsername string, opts ...interface{}) (*Users, *Response, error) {
	result := new(Users)
	urlStr := t.client.requestURL("/teams/%s/following", teamUsername)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
