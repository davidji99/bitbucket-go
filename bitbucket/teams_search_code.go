package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// SearchCode searches for code in the repositories of the specified team.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/search/code#get
func (t *TeamsService) SearchCode(teamUsername string, opts ...interface{}) (*SearchCodeResults, *simpleresty.Response, error) {
	results := new(SearchCodeResults)
	urlStr, urlStrErr := t.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/teams/%s/search/code", teamUsername), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := t.client.http.Get(urlStr, results, nil)

	return results, response, err
}
