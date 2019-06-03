package bitbucket

// ListFollowing returns the list of accounts this team is following.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/following#get
func (t *TeamsService) ListFollowing(teamUsername string) (*Users, *Response, error) {
	result := new(Users)
	urlStr := t.client.requestUrl("/teams/%s/following", teamUsername)
	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
