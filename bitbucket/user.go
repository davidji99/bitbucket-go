package bitbucket

import (
	"github.com/davidji99/simpleresty"
	"time"
)

// UserService handles communication with the user related methods
// of the Bitbucket API.
//
// This service only deals with returning information about the authenticated user.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/user
type UserService service

// User represents a Bitbucket user.
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
	AccountID     *string    `json:"account_id,omitempty"`
}

// Get returns the currently authenticated user.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/user#get
func (u *UserService) Get() (*User, *simpleresty.Response, error) {
	user := new(User)
	urlStr := u.client.http.RequestURL("/user")
	response, err := u.client.http.Get(urlStr, user, nil)

	return user, response, err
}
