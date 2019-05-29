package bitbucket

import "time"

// CommitsService handles communication with the commit related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commits
type CommitsService service

// Issues represent a collection of issues.
type Commits struct {
	Page     int    `json:"page,omitempty"`
	Next     string `json:"next,omitempty"`
	Pagelen  int    `json:"pagelen,omitempty"`
	Size     int    `json:"size,omitempty"`
	Previous string `json:"previous,omitempty"`

	Values []*Commit `json:"values,omitempty"`
}

type Commit struct {
	Rendered struct {
		Message *BitbucketContent `json:"message,omitempty"`
	} `json:"rendered,omitempty"`
	Hash         *string           `json:"hash,omitempty"`
	Repository   *Repository       `json:"repository,omitempty"`
	Links        *CommitLinks      `json:"links,omitempty"`
	Author       *User             `json:"author,omitempty"`
	Summary      *BitbucketContent `json:"summary,omitempty"`
	Participants []*User           `json:"participants,omitempty"`
	Parent       []*Commit         `json:"parent,omitempty"`
	Date         *time.Time        `json:"date,omitempty"`
	Message      *string           `json:"message,omitempty"`
	Type         *string           `json:"type,omitempty"`
}

type CommitLinks struct {
	Self     *BitbucketLink `json:"self,omitempty"`
	Comment  *BitbucketLink `json:"comment,omitempty"`
	HTML     *BitbucketLink `json:"html,omitempty"`
	Diff     *BitbucketLink `json:"diff,omitempty"`
	Approve  *BitbucketLink `json:"approve,omitempty"`
	Statuses *BitbucketLink `json:"statuses,omitempty"`
}
