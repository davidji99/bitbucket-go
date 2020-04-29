package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// ListTeamRepositories returns the list of accounts that are following this team.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/repositories#get
func (t *TeamsService) ListTeamRepositories(teamUsername string, opts ...interface{}) (*Repositories, *simpleresty.Response, error) {
	result := new(Repositories)
	urlStr, urlStrErr := t.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/teams/%s/repositories", teamUsername), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := t.client.http.Get(urlStr, result, nil)

	return result, response, err
}
