package bitbucket

import "time"

// IssuesService handles communication with the issue related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues
type IssuesService service

// Issues represent a collection of issues.
type Issues struct {
	Page     int    `json:"page,omitempty"`
	Next     string `json:"next,omitempty"`
	Pagelen  int    `json:"pagelen,omitempty"`
	Size     int    `json:"size,omitempty"`
	Previous string `json:"previous,omitempty"`

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

// IssueLinks represents the "links" object in a Bitbucket issue,
type IssueLinks struct {
	Self        *BitbucketLink `json:"self,omitempty"`
	Attachments *BitbucketLink `json:"attachments,omitempty"`
	Watch       *BitbucketLink `json:"watch,omitempty"`
	Comments    *BitbucketLink `json:"comments,omitempty"`
	HTML        *BitbucketLink `json:"html,omitempty"`
	Vote        *BitbucketLink `json:"vote,omitempty"`
}

type IssueContent struct {
	Raw    *string `json:"raw,omitempty"`
	Markup *string `json:"markup,omitempty"`
	Html   *string `json:"html,omitempty"`
	Type   *string `json:"type,omitempty"`
}

type IssueRequest struct {
	Title    *string                  `json:"title,omitempty"`
	Kind     *string                  `json:"kind,omitempty"`
	Priority *string                  `json:"priority,omitempty"`
	Content  *CreateIssueContentOpts  `json:"content,omitempty"`
	Assignee *CreateIssueAssigneeOpts `json:"assignee,omitempty"`
}

type CreateIssueContentOpts struct {
	Raw *string `json:"raw,omitempty"`
}

type CreateIssueAssigneeOpts struct {
	Username *string `json:"username,omitempty"`
}

func (i *IssuesService) List(owner, repoSlug string) (*Issues, *Response, error) {
	issues := new(Issues)
	urlStr := i.client.requestUrl("/repositories/%s/%s/issues", owner, repoSlug)

	response, err := i.client.execute("GET", urlStr, &issues, nil, "")
	return issues, response, err
}

func (i *IssuesService) Get(owner, repoSlug, issueId string) (*Issue, *Response, error) {
	result := new(Issue)
	urlStr := i.client.requestUrl("/repositories/%s/%s/issues/%s", owner, repoSlug, issueId)
	response, err := i.client.execute("GET", urlStr, result, nil, "")

	return result, response, err
}

func (i *IssuesService) Create(owner, repoSlug string, io *IssueRequest) (*Issue, *Response, error) {
	result := new(Issue)
	urlStr := i.client.requestUrl("/repositories/%s/%s/issues", owner, repoSlug)
	response, err := i.client.execute("POST", urlStr, result, io, "")

	return result, response, err
}
