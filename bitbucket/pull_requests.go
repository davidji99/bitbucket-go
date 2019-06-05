package bitbucket

import "time"

// PullRequestService handles communication with the pull requests related
// methods of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests
type PullRequestsService service

type PullRequests struct {
	PaginationInfo

	Values []*PullRequest `json:"values,omitempty"`
}

// PullRequest represents a Bitbucket pull request on a repository.
type PullRequest struct {
	Body              *PullRequestBody   `json:"rendered,omitempty"`
	Type              *string            `json:"type,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Links             *PullRequestLinks  `json:"links,omitempty"`
	Title             *string            `json:"title,omitempty"`
	CloseSourceBranch *bool              `json:"close_source_branch,omitempty"`
	Reviewers         []*User            `json:"Reviewers,omitempty"` // TODO: validate
	ID                *int64             `json:"id,omitempty"`
	Destination       *PullRequestBranch `json:"destination,omitempty"`
	CreatedOn         *time.Time         `json:"created_on,omitempty"`
	Summary           *BitbucketContent  `json:"summary,omitempty"`
	Source            *PullRequestBranch `json:"source,omitempty"`
	CommentCount      *int64             `json:"comment_count,omitempty"`
	State             *string            `json:"state,omitempty"`
	TaskCount         *int64             `json:"task_count,omitempty"`
	Participants      []*User            `json:"participants,omitempty"`
	Reason            *string            `json:"reason,omitempty"`
	UpdatedOn         *string            `json:"updated_on,omitempty"`
	Author            *User              `json:"author,omitempty"`
	MergeCommit       *string            `json:"merge_commit,omitempty"`
	ClosedBy          *User              `json:"closed_by,omitempty"`
}

// PullRequestLinks represents the "links" object in a Bitbucket pull request.
type PullRequestLinks struct {
	Decline  *BitbucketLink `json:"decline,omitempty"`
	Commits  *BitbucketLink `json:"commits,omitempty"`
	Self     *BitbucketLink `json:"self,omitempty"`
	Comments *BitbucketLink `json:"comments,omitempty"`
	Merge    *BitbucketLink `json:"merge,omitempty"`
	HTML     *BitbucketLink `json:"html,omitempty"`
	Activity *BitbucketLink `json:"activity,omitempty"`
	Diff     *BitbucketLink `json:"diff,omitempty"`
	Approve  *BitbucketLink `json:"approve,omitempty"`
	Statuses *BitbucketLink `json:"statuses,omitempty"`
}

// PullRequestBody represents the body of a Bitbucket pull request.
type PullRequestBody struct {
	Description *BitbucketContent `json:"description,omitempty"`
	Title       *BitbucketContent `json:"title,omitempty"`
}

// PullRequestBranch represents a branch associated with the pull request.
type PullRequestBranch struct {
	Commit     *Commit     `json:"commit,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Branch     *Branch     `json:"branch,omitempty"`
}

// NewPullRequestOpts represents a new pull request to be created.
type NewPullRequestOpts struct {
	Title             *string                        `json:"title,omitempty"`  // Required field
	Source            *NewPullRequestSourceOpts      `json:"source,omitempty"` // Required field
	Destination       *NewPullRequestDestinationOpts `json:"destination,omitempty"`
	Reviewers         []*NewPullRequestReviewerOpts  `json:"reviewers"`
	Description       *string                        `json:"description,omitempty"`
	CloseSourceBranch *bool                          `json:"close_source_branch"`
}

// UpdatePullRequestOpts represents the fields that are editable for an existing pull request.
type UpdatePullRequestOpts struct {
	Title       *string                        `json:"title,omitempty"` // Required field
	Description *string                        `json:"description,omitempty"`
	Source      *NewPullRequestSourceOpts      `json:"source,omitempty"`
	Destination *NewPullRequestDestinationOpts `json:"destination,omitempty"`
}

// PullRequestListOpts represents the filters and query parameters available when listing pull requests.
type PullRequestListOpts struct {
	// An array of pull request states that should be returned.
	// Valid options: MERGED, SUPERSEDED, OPEN, DECLINED. Case sensitive.
	// By default, only OPEN pull requests are returned.
	State []string `url:"state,omitempty"`

	FilterSortOpts
}

// NewPullRequestSourceOpts represents the source branch for the new pull request.
type NewPullRequestSourceOpts struct {
	Branch *Branch `json:"branch,omitempty"`
}

// NewPullRequestDestinationOpts represents the destination branch for the new pull request.
type NewPullRequestDestinationOpts struct {
	Branch *Branch `json:"branch,omitempty"`
}

// NewPullRequestReviewerOpts represent a reviewer for a pull request specified by the user's UUID.
type NewPullRequestReviewerOpts struct {
	UUID *string `json:"uuid,omitempty"`
}

// Branch represents a branch.
type Branch struct {
	Name *string `json:"name,omitempty"`
}

// List all pull requests for a given repository.
// Supports filtering by passing in a non-URI encoded query string. Reference: https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
// Example query string: source.repository.full_name != "main/repo" AND state = "OPEN" AND reviewers.username = "evzijst" AND destination.branch.name = "master"
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests#get
func (p *PullRequestsService) List(owner, repoSlug string, opts ...interface{}) (*PullRequests, *Response, error) {
	pullRequests := new(PullRequests)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, pullRequests, nil)

	return pullRequests, response, err
}

// Get a single pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D
func (p *PullRequestsService) Get(owner, repoSlug string, pullRequestId int64) (*PullRequest, *Response, error) {
	pr := new(PullRequest)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v", owner, repoSlug, pullRequestId)
	response, err := p.client.execute("GET", urlStr, pr, nil)

	return pr, response, err
}

// Create a new pull request.
// The minimum required fields to create a pull request are title and source, specified by a branch name.
// If the pull request's destination is not specified, it will default to the repository.mainbranch.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests#post
func (p *PullRequestsService) Create(owner, repoSlug string, po NewPullRequestOpts) (*PullRequest, *Response, error) {
	pr := new(PullRequest)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/", owner, repoSlug)
	response, err := p.client.execute("POST", urlStr, pr, po)

	return pr, response, err
}

// Update a pull request.
// This can be used to change the pull request's branches or description. Only open pull requests can be mutated.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D#put
func (p *PullRequestsService) Update(owner, repoSlug string, pullRequestId int64, po UpdatePullRequestOpts) (*PullRequest, *Response, error) {
	pr := new(PullRequest)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v", owner, repoSlug, pullRequestId)
	response, err := p.client.execute("PUT", urlStr, pr, po)

	return pr, response, err
}
