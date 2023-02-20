package bitbucket

import "github.com/davidji99/simpleresty"

// MergePrRequest represents a request to merge a pull request.
type MergePrRequest struct {
	// Type of merge. Required
	Type string `json:"type"`

	// The commit message that will be used on the resulting commit.
	Message string `json:"message,omitempty"`

	// Whether the source branch should be deleted. If this is not provided, we fallback to the value used when the
	// pull request was created, which defaults to False
	CloseSourceBranch *bool `json:"close_source_branch,omitempty"`

	// The merge strategy that will be used to merge the pull request. Default: merge_commit
	// Valid values: merge_commit, squash, fast_forward
	MergeStrategy string `json:"merge_strategy,omitempty"`
}

// MergePR merges the pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-merge-post
func (p *PullRequestsService) MergePR(workspace, repoSlug string, pullRequestID int64, opts *MergePrRequest) (*PullRequest, *simpleresty.Response, error) {
	result := new(PullRequest)
	urlStr := p.client.http.RequestURL("/repositories/%s/%s/pullrequests/%v/merge", workspace, repoSlug, pullRequestID)
	response, err := p.client.http.Post(urlStr, result, opts)

	return result, response, err
}
