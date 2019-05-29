package bitbucket

import "time"

// UsersService handles communication with the user related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users
type UsersService service

type User struct {
	Username      *string    `json:"username,omitempty"`
	Website       *string    `json:"website,omitempty"`
	DisplayName   *string    `json:"display_name,omitempty"`
	UUID          *string    `json:"uuid,omitempty"`
	Links         *UserLinks `json:"links,omitempty"`
	Nickname      *string    `json:"nickname,omitempty"`
	CreatedOn     *time.Time `json:"created_on,omitempty"`
	IsStaff       *bool      `json:"is_staff,omitempty"`
	Location      *string    `json:"location,omitempty"`
	AccountStatus *string    `json:"account_status,omitempty"`
	Type          *string    `json:"type,omitempty"`
	AccountId     *string    `json:"account_id,omitempty"`
}

type UserLinks struct {
	Hooks        *BitbucketLink `json:"hooks,omitempty"`
	Self         *BitbucketLink `json:"self,omitempty"`
	Repositories *BitbucketLink `json:"repositories,omitempty"`
	HTML         *BitbucketLink `json:"html,omitempty"`
	Followers    *BitbucketLink `json:"followers,omitempty"`
	Avatar       *BitbucketLink `json:"avatar,omitempty"`
	Following    *BitbucketLink `json:"following,omitempty"`
	Snippets     *BitbucketLink `json:"snippet,omitempty"`
}
