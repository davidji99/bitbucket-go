package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// ListMembers returns all members of the specified team.
//
// Any member of any of the team's groups is considered a member of the team.
// This includes users in groups that may not actually have access to any of the team's repositories.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/members#get
func (t *TeamsService) ListMembers(teamUsername string, opts ...interface{}) (*Users, *simpleresty.Response, error) {
	result := new(Users)
	urlStr, urlStrErr := t.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/teams/%s/members", teamUsername), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := t.client.http.Get(urlStr, result, nil)

	return result, response, err
}
