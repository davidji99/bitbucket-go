package bitbucket

import "time"

// PullRequestService handles communication with the pull requests related
// methods of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests
type PullRequestsService service

type PullRequests struct {
	Page     int    `json:"page,omitempty"`
	Next     string `json:"next,omitempty"`
	Pagelen  int    `json:"pagelen,omitempty"`
	Size     int    `json:"size,omitempty"`
	Previous string `json:"previous,omitempty"`

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

type PullRequestBody struct {
	Description *BitbucketContent `json:"description,omitempty"`
	Title       *BitbucketContent `json:"title,omitempty"`
}

type PullRequestBranch struct {
	Commit     *Commit     `json:"commit,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Branch     struct {
		Name *string `json:"name,omitempty"`
	} `json:"branch,omitempty"`
}

type NewPullRequest struct {
	//State             string   `json:"state"`
	//CommentID         string   `json:"comment_id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	//CloseSourceBranch bool     `json:"close_source_branch"`
	Source      *CreatePullRequestSourceOpts      `json:"source,omitempty"`
	Destination *CreatePullRequestDestinationOpts `json:"destination,omitempty"`
	//SourceRepository  string   `json:"source_repository"`
	//DestinationCommit string   `json:"destination_repository"`
	//Message           string   `json:"message"`
	//Reviewers         []string `json:"reviewers"`
}

type CreatePullRequestSourceOpts struct {
	Branch *Branch `json:"branch,omitempty"`
}

type CreatePullRequestDestinationOpts struct {
	Branch *Branch `json:"branch,omitempty"`
}

type Branch struct {
	Name *string `json:"name,omitempty"`
}

func (p *PullRequestsService) List(owner, repo, opts string) ([]*PullRequest, *Response, error) {
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/", owner, repo)

	var pulls []*PullRequest
	response, err := p.client.execute("GET", urlStr, pulls, nil, opts)

	return pulls, response, err
}

func (p *PullRequestsService) Get(owner, repo, id string) (*PullRequest, *Response, error) {
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%s", owner, repo, id)

	result := new(PullRequest)
	response, err := p.client.execute("GET", urlStr, result, nil, "")

	return result, response, err
}

func (p *PullRequestsService) Create(owner, repo string, po NewPullRequest) (*PullRequest, *Response, error) {
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/", owner, repo)

	result := new(PullRequest)
	response, err := p.client.execute("POST", urlStr, result, po, "")

	return result, response, err
}

//func (p *PullRequestsService) Update(owner, repo, id string, po NewPullRequest) (*PullRequest, *Response, error) {
//	urlStr := p.client.BaseURL + "/repositories/" + owner + "/" + repo + "/pullrequests/" + id + "/patch"
//
//	result := new(PullRequest)
//	response, err := p.client.execute("POST", urlStr, result, po, "")
//
//	return result, response, err
//}
