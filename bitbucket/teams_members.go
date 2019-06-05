package bitbucket

// ListMembers returns all members of the specified team.
//
// Any member of any of the team's groups is considered a member of the team.
// This includes users in groups that may not actually have access to any of the team's repositories.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/members#get
func (t *TeamsService) ListMembers(teamUsername string, opts ...interface{}) (*Users, *Response, error) {
	result := new(Users)
	urlStr := t.client.requestUrl("/teams/%s/members", teamUsername)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
