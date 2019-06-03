package bitbucket

import "time"

const (
	UpdateActivity   = "update"
	ApprovalActivity = "approval"
)

type PRActivities struct {
	PaginationInfo

	Values []*PRActivity `json:"values,omitempty"`
}

type PRActivity struct {
	Update      *PRUpdateActivity   `json:"update,omitempty"`
	Approval    *PRApprovalActivity `json:"approval,omitempty"`
	PullRequest *PullRequest        `json:"pull_request,omitempty"`
}

type PRUpdateActivity struct {
	Description *string            `json:"description,omitempty"`
	Title       *string            `json:"title,omitempty"`
	Destination *PullRequestBranch `json:"destination,omitempty"`
	Reason      *string            `json:"reason,omitempty"`
	Source      *PullRequestBranch `json:"source,omitempty"`
	State       *string            `json:"state,omitempty"`
	Author      *User              `json:"author,omitempty"`
	Date        *time.Time         `json:"date,omitempty"`
}

type PRApprovalActivity struct {
	Date        *time.Time   `json:"date,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	User        *User        `json:"user,omitempty"`
}

// ListActivity returns a paginated list of all pull requests' activity log on a specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/activity#get
func (p *PullRequestsService) ListActivity(owner, repoSlug string, opts *ListPaginationOpts) (*PRActivities, *Response, error) {
	result := new(PRActivities)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/activity", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// GetActivity returns a paginated list of a single pull request's activity log in a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/activity#get
func (p *PullRequestsService) GetActivity(owner, repoSlug string, pullRequestId int64) (*PRActivities, *Response, error) {
	result := new(PRActivities)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/activity", owner, repoSlug, pullRequestId)
	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// GetActivityType returns the non-nil field representing the activity: an update or approval.
// It returns the activity object and its type.
func (a *PRActivity) GetActivityType() (interface{}, string) {
	if v, _ := a.GetUpdate(); v != nil {
		return a.GetUpdate()
	}

	if v, _ := a.GetApproval(); v != nil {
		return a.GetUpdate()
	}

	return nil, ""
}

// GetUpdate returns the Update field if it's non-nil, nil otherwise.
func (a *PRActivity) GetUpdate() (interface{}, string) {
	if a == nil || a.Update == nil {
		return nil, ""
	}
	return *a.Update, UpdateActivity
}

// GetApproval returns the Approval field if it's non-nil, nil otherwise.
func (a *PRActivity) GetApproval() (interface{}, string) {
	if a == nil || a.Approval == nil {
		return nil, ""
	}
	return *a.Approval, ApprovalActivity
}
