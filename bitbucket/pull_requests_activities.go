package bitbucket

import "time"

const (
	UpdateActivity   = "update"
	ApprovalActivity = "approval"
)

// PRActivities represents a collection of of pull request activity.
type PRActivities struct {
	PaginationInfo

	Values []*PRActivity `json:"values,omitempty"`
}

// PRActivity represents a pull request activity.
type PRActivity struct {
	Update      *PRUpdateActivity   `json:"update,omitempty"`
	Approval    *PRApprovalActivity `json:"approval,omitempty"`
	PullRequest *PullRequest        `json:"pull_request,omitempty"`
}

// PRUpdateActivity represents a pull request update activity.
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

// PRApprovalActivity represents a pull request approval activity.
type PRApprovalActivity struct {
	Date        *time.Time   `json:"date,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	User        *User        `json:"user,omitempty"`
}

// ListActivity returns a paginated list of all pull requests' activity log on a specified repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/activity#get
func (p *PullRequestsService) ListActivity(owner, repoSlug string, opts ...interface{}) (*PRActivities, *Response, error) {
	result := new(PRActivities)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/activity", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// GetActivity returns a paginated list of a single pull request's activity log in a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/pullrequests/%7Bpull_request_id%7D/activity#get
func (p *PullRequestsService) GetActivity(owner, repoSlug string, pullRequestId int64, opts ...interface{}) (*PRActivities, *Response, error) {
	result := new(PRActivities)
	urlStr := p.client.requestUrl("/repositories/%s/%s/pullrequests/%v/activity", owner, repoSlug, pullRequestId)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := p.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// GetActivityType returns the non-nil field representing the activity: an update or approval.
// It returns the activity object and its type.
func (p *PRActivity) GetActivityType() (interface{}, string) {
	if v := p.GetUpdate(); v != nil {
		return p.GetUpdate(), UpdateActivity
	}

	if v := p.GetApproval(); v != nil {
		return p.GetUpdate(), ApprovalActivity
	}

	return nil, ""
}
