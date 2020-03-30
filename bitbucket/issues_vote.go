package bitbucket

import "github.com/davidji99/simpleresty"

// HasCurrentUserVoted check whether the authenticated user has voted for this issue.
//
// A 204 status code indicates that the user has voted, while a 404 implies they haven't.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/vote#get
func (i *IssuesService) HasCurrentUserVoted(owner, repoSlug string, id int64) (bool, *simpleresty.Response, error) {
	urlStr := i.client.http.RequestURL("/repositories/%s/%s/issues/%v/vote", owner, repoSlug, id)
	response, err := i.client.http.Get(urlStr, nil, nil)
	if err != nil {
		return false, nil, err
	}

	hasVoted := false
	if response.StatusCode == 204 {
		hasVoted = true
	}

	return hasVoted, response, nil
}

// Vote adds a vote on behalf of the authenticated user only.
//
// The 204 status code indicates that the operation was successful.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/vote#put
func (i *IssuesService) Vote(owner, repoSlug string, id int64) (*simpleresty.Response, error) {
	urlStr := i.client.http.RequestURL("/repositories/%s/%s/issues/%v/vote", owner, repoSlug, id)
	response, err := i.client.http.Put(urlStr, nil, nil)

	return response, err
}

// RemoveVote retract your vote.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/vote#delete
func (i *IssuesService) RemoveVote(owner, repoSlug string, id int64) (*simpleresty.Response, error) {
	urlStr := i.client.http.RequestURL("/repositories/%s/%s/issues/%v/vote", owner, repoSlug, id)
	response, err := i.client.http.Delete(urlStr, nil, nil)

	return response, err
}
