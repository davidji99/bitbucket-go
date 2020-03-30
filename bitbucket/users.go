package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// UsersService handles communication with the users related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users
type UsersService service

// Users represents a collection of users.
type Users struct {
	PaginationInfo

	Values []*User `json:"values,omitempty"`
}

// UserLinks represents the "links" object in a Bitbucket user.
type UserLinks struct {
	Hooks        *Link `json:"hooks,omitempty"`
	Self         *Link `json:"self,omitempty"`
	Repositories *Link `json:"repositories,omitempty"`
	HTML         *Link `json:"html,omitempty"`
	Followers    *Link `json:"followers,omitempty"`
	Avatar       *Link `json:"avatar,omitempty"`
	Following    *Link `json:"following,omitempty"`
	Snippets     *Link `json:"snippet,omitempty"`
}

// GetByID fetches a single user.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D#get
func (u *UsersService) GetByID(userID string, opts ...interface{}) (*User, *simpleresty.Response, error) {
	user := new(User)
	urlStr, urlStrErr := u.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/users/%s", userID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := u.client.http.Get(urlStr, user, nil)

	return user, response, err
}
