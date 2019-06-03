package bitbucket

// SearchCode searches for code in the repositories of the specified team.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/search/code#get
func (t *TeamsService) SearchCode(teamUsername string, opts *CodeSearchOpts) (*SearchCodeResults, *Response, error) {
	results := new(SearchCodeResults)
	urlStr := t.client.requestUrl("/teams/%s/search/code", teamUsername)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, results, nil)

	return results, response, err
}
