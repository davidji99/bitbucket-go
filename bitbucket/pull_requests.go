package bitbucket

import "time"

// PullRequestsService handles communication with the pull requests related
// methods of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests
type PullRequestsService service

// PullRequests represents a collection of pull requests.
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
	Reviewers         []*User            `json:"reviewers,omitempty"`
	ID                *int64             `json:"id,omitempty"`
	Destination       *PullRequestBranch `json:"destination,omitempty"`
	CreatedOn         *time.Time         `json:"created_on,omitempty"`
	Summary           *Content           `json:"summary,omitempty"`
	Source            *PullRequestBranch `json:"source,omitempty"`
	CommentCount      *int64             `json:"comment_count,omitempty"`
	State             *string            `json:"state,omitempty"`
	TaskCount         *int64             `json:"task_count,omitempty"`
	Participants      []*Participant     `json:"participants,omitempty"`
	Reason            *string            `json:"reason,omitempty"`
	UpdatedOn         *string            `json:"updated_on,omitempty"`
	Author            *User              `json:"author,omitempty"`
	MergeCommit       *string            `json:"merge_commit,omitempty"`
	ClosedBy          *User              `json:"closed_by,omitempty"`
}

// PullRequestLinks represents the "links" object in a Bitbucket pull request.
type PullRequestLinks struct {
	Decline  *Link `json:"decline,omitempty"`
	Commits  *Link `json:"commits,omitempty"`
	Self     *Link `json:"self,omitempty"`
	Comments *Link `json:"comments,omitempty"`
	Merge    *Link `json:"merge,omitempty"`
	HTML     *Link `json:"html,omitempty"`
	Activity *Link `json:"activity,omitempty"`
	Diff     *Link `json:"diff,omitempty"`
	Approve  *Link `json:"approve,omitempty"`
	Statuses *Link `json:"statuses,omitempty"`
}

// PullRequestBody represents the body of a Bitbucket pull request.
type PullRequestBody struct {
	Description *Content `json:"description,omitempty"`
	Title       *Content `json:"title,omitempty"`
}

// PullRequestBranch represents a branch associated with the pull request.
type PullRequestBranch struct {
	Commit     *Commit     `json:"commit,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Branch     *Branch     `json:"branch,omitempty"`
}

// PRRequest represents a request to create/update a pull request.
type PRRequest struct {
	Title             *string                   `json:"title,omitempty"`
	Source            *PRRequestSourceOpts      `json:"source,omitempty"`
	Destination       *PRRequestDestinationOpts `json:"destination,omitempty"`
	Reviewers         []*PRRequestReviewerOpts  `json:"reviewers"`
	Description       *string                   `json:"description,omitempty"`
	CloseSourceBranch *bool                     `json:"close_source_branch"`
}

// PullRequestListOpts represents the filters and query parameters available when listing pull requests.
type PullRequestListOpts struct {
	// An array of pull request states that should be returned.
	// Valid options: MERGED, SUPERSEDED, OPEN, DECLINED. Case sensitive.
	// By default, only OPEN pull requests are returned.
	State []string `url:"state,omitempty"`
}

// PRRequestSourceOpts represents the source branch for the pull request.
type PRRequestSourceOpts struct {
	Branch *Branch `json:"branch,omitempty"`
}

// PRRequestDestinationOpts represents the destination branch for the pull request.
type PRRequestDestinationOpts struct {
	Branch *Branch `json:"branch,omitempty"`
}

// PRRequestReviewerOpts represent a reviewer for a pull request specified by the user's UUID.
type PRRequestReviewerOpts struct {
	UUID *string `json:"uuid,omitempty"`
}

// Branch represents a branch.
type Branch struct {
	Name *string `json:"name,omitempty"`
}

// List returns all pull requests for a given repository.
// Supports filtering by passing in a non-URI encoded query string. Reference: https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
// Example query string: source.repository.full_name != "main/repo" AND state = "OPEN" AND reviewers.username = "evzijst" AND destination.branch.name = "master"
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests#get
func (p *PullRequestsService) List(owner, repoSlug string, opts ...interface{}) (*PullRequests, *Response, error) {
	result := new(PullRequests)
	urlStr := p.client.requestURL("/repositories/%s/%s/pullrequests", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Get returns a single pull request.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D
func (p *PullRequestsService) Get(owner, repoSlug string, pullRequestID int64, opts ...interface{}) (*PullRequest, *Response, error) {
	result := new(PullRequest)
	urlStr := p.client.requestURL("/repositories/%s/%s/pullrequests/%v", owner, repoSlug, pullRequestID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// ListByUser returns all pull requests authored by the specified user.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/pullrequests/%7Btarget_user%7D#get
func (p *PullRequestsService) ListByUser(targetUser string, opts ...interface{}) (*PullRequests, *Response, error) {
	result := new(PullRequests)
	urlStr := p.client.requestURL("/pullrequests/%s", targetUser)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Create a new pull request.
// The minimum required fields to create a pull request are title and source, specified by a branch name.
// If the pull request's destination is not specified, it will default to the repository.mainbranch.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests#post
func (p *PullRequestsService) Create(owner, repoSlug string, po *PRRequest) (*PullRequest, *Response, error) {
	result := new(PullRequest)
	urlStr := p.client.requestURL("/repositories/%s/%s/pullrequests", owner, repoSlug)
	response, err := p.client.execute("POST", urlStr, result, po)

	return result, response, err
}

// Update a pull request.
// This can be used to change the pull request's branches or description. Only open pull requests can be mutated.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D#put
func (p *PullRequestsService) Update(owner, repoSlug string, pullRequestID int64, po *PRRequest) (*PullRequest, *Response, error) {
	result := new(PullRequest)
	urlStr := p.client.requestURL("/repositories/%s/%s/pullrequests/%v", owner, repoSlug, pullRequestID)
	response, err := p.client.execute("PUT", urlStr, result, po)

	return result, response, err
}
