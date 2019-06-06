package bitbucket

import "time"

// Comment represents a generic comment on a Bitbucket pull request or issue.
type Comment struct {
	Links     *CommentLinks `json:"links,omitempty"`
	Content   *Content      `json:"content,omitempty"`
	CreatedOn *time.Time    `json:"created_on,omitempty"`
	User      *User         `json:"user,omitempty"`
	UpdatedOn *time.Time    `json:"updated_on,omitempty"`
	Type      *string       `json:"type,omitempty"`
	ID        *int64        `json:"id,omitempty"`
}

// CommentLinks represents the "links" object in a Bitbucket comment.
type CommentLinks struct {
	Self *Link `json:"self,omitempty"`
	HTML *Link `json:"html,omitempty"`
}
