package bitbucket

// Approve approves the specified commit as the authenticated user.
//
// This operation is only available to users that have explicit access to the repository.
// In contrast, just the fact that a repository is publicly accessible to users does not give them the ability to approve commits.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/approve#post
func (c *CommitService) Approve(owner, repoSlug, sha string) (*Participant, *Response, error) {
	results := new(Participant)
	urlStr := c.client.requestURL("/repositories/%s/%s/commit/%s/approve", owner, repoSlug, sha)
	response, err := c.client.execute("POST", urlStr, results, nil)

	return results, response, err
}

// UnApprove redacts/removes the authenticated user's approval of the specified commit.
//
// This operation is only available to users that have explicit access to the repository.
// In contrast, just the fact that a repository is publicly accessible to users does not give them the ability to approve commits.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/approve#delete
func (c *CommitService) UnApprove(owner, repoSlug, sha string) (*Response, error) {
	urlStr := c.client.requestURL("/repositories/%s/%s/commit/%s/approve", owner, repoSlug, sha)
	response, err := c.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
