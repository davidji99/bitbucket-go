package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// ListFollowing returns the list of accounts this team is following.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/following#get
func (t *TeamsService) ListFollowing(teamUsername string, opts ...interface{}) (*Users, *simpleresty.Response, error) {
	result := new(Users)
	urlStr, urlStrErr := t.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/teams/%s/following", teamUsername), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := t.client.http.Get(urlStr, result, nil)

	return result, response, err
}
