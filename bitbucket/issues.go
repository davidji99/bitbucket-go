package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

// IssuesService handles communication with the issue related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues
type IssuesService service

// Issues represent a collection of issues.
type Issues struct {
	PaginationInfo

	Values []*Issue `json:"values,omitempty"`
}

// Issue represents a Bitbucket issue on a repository.
type Issue struct {
	Priority   *string       `json:"priority,omitempty"`
	Kind       *string       `json:"kind,omitempty"`
	Repository *Repository   `json:"repository,omitempty"`
	Links      *IssueLinks   `json:"links,omitempty"`
	Reporter   *User         `json:"reporter,omitempty"`
	Title      *string       `json:"title,omitempty"`
	Component  *Component    `json:"component,omitempty"`
	Votes      *int          `json:"votes,omitempty"`
	Watches    *int          `json:"watches,omitempty"`
	Content    *IssueContent `json:"content,omitempty"`
	Assignee   *User         `json:"assignee,omitempty"`
	State      *string       `json:"state,omitempty"`
	Type       *string       `json:"type,omitempty"`
	Version    *Version      `json:"version,omitempty"`
	EditedOn   *time.Time    `json:"edited_on,omitempty"`
	CreatedOn  *time.Time    `json:"created_on,omitempty"`
	Milestone  *Milestone    `json:"milestone,omitempty"`
	UpdatedOn  *time.Time    `json:"updated_on,omitempty"`
	ID         *int64        `json:"id,omitempty"`
}

// IssueLinks represents the "links" object in a Bitbucket issue.
type IssueLinks struct {
	Self        *Link `json:"self,omitempty"`
	Attachments *Link `json:"attachments,omitempty"`
	Watch       *Link `json:"watch,omitempty"`
	Comments    *Link `json:"comments,omitempty"`
	HTML        *Link `json:"html,omitempty"`
	Vote        *Link `json:"vote,omitempty"`
}

// IssueContent represents the Description box in the Bitbucket issue UI.
type IssueContent struct {
	Raw    *string `json:"raw,omitempty"`
	Markup *string `json:"markup,omitempty"`
	HTML   *string `json:"html,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// IssueRequest represents a request to create/update an issue.
type IssueRequest struct {
	Title     *string                   `json:"title,omitempty"`    // Required field.
	Kind      *string                   `json:"kind,omitempty"`     // Required field.
	Priority  *string                   `json:"priority,omitempty"` // Required field.
	Content   *IssueRequestContentOpts  `json:"content,omitempty"`
	Component *ComponentRequest         `json:"component,omitempty"`
	Milestone *MilestoneRequest         `json:"milestone,omitempty"`
	Version   *VersionRequest           `json:"version,omitempty"`
	Assignee  *IssueRequestAssigneeOpts `json:"assignee,omitempty"`
}

// IssueRequestContentOpts represents the Description box when creating/updating a new issue.
type IssueRequestContentOpts struct {
	Raw  *string `json:"raw,omitempty"`
	HTML *string `json:"html,omitempty"`
}

// IssueRequestAssigneeOpts represents the Bitbucket user to be assigned when creating/updating a new issue.
type IssueRequestAssigneeOpts struct {
	Username *string `json:"username,omitempty"`
}

// List returns all issues for a given repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues#get
func (i *IssuesService) List(owner, repoSlug string, opts ...interface{}) (*Issues, *simpleresty.Response, error) {
	result := new(Issues)
	urlStr, urlStrErr := i.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/issues", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := i.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// Get a single issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D#get
func (i *IssuesService) Get(owner, repoSlug string, issueID int64, opts ...interface{}) (*Issue, *simpleresty.Response, error) {
	result := new(Issue)
	urlStr, urlStrErr := i.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/issues/%v", owner, repoSlug, issueID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := i.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// Create a new issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues#post
func (i *IssuesService) Create(owner, repoSlug string, io *IssueRequest) (*Issue, *simpleresty.Response, error) {
	result := new(Issue)
	urlStr := i.client.http.RequestURL("/repositories/%s/%s/issues", owner, repoSlug)
	response, err := i.client.http.Post(urlStr, result, io)

	return result, response, err
}

// Update an issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D#put
func (i *IssuesService) Update(owner, repoSlug string, issueID int64, io *IssueRequest) (*Issue, *simpleresty.Response, error) {
	result := new(Issue)
	urlStr := i.client.http.RequestURL("/repositories/%s/%s/issues/%v", owner, repoSlug, issueID)
	response, err := i.client.http.Put(urlStr, result, io)

	return result, response, err
}

// Delete the specified issue. This requires write access to the repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D#delete
func (i *IssuesService) Delete(owner, repoSlug string, issueID int64) (*simpleresty.Response, error) {
	urlStr := i.client.http.RequestURL("/repositories/%s/%s/issues/%v", owner, repoSlug, issueID)
	response, err := i.client.http.Delete(urlStr, nil, nil)

	return response, err
}
