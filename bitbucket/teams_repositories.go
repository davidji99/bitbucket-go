package bitbucket

// ListTeamRepositories returns the list of accounts that are following this team.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/repositories#get
func (t *TeamsService) ListTeamRepositories(teamUsername string, opts *ListPaginationOpts) (*Repositories, *Response, error) {
	result := new(Repositories)
	urlStr := t.client.requestUrl("/teams/%s/repositories", teamUsername)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
