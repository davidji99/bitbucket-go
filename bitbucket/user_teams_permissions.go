package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// UserTeamsPermissions represents a collection of a user's permissions on repositories.
type UserTeamsPermissions struct {
	PaginationInfo

	Values []*UserTeamsPermission `json:"values,omitempty"`
}

// UserTeamsPermission represents a user's team permission.
type UserTeamsPermission struct {
	Type       *string `json:"type,omitempty"`
	User       *User   `json:"user,omitempty"`
	Team       *Team   `json:"team,omitempty"`
	Permission *string `json:"permission,omitempty"`
}

// ListTeamsPerms returns permissions for each team the caller is a member of and the highest level of privilege the caller has.
//
// If a user is a member of multiple groups with distinct roles, only the highest level is returned.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/user/permissions/teams#get
func (u *UserService) ListTeamsPerms(opts ...interface{}) (*UserTeamsPermissions, *simpleresty.Response, error) {
	perms := new(UserTeamsPermissions)
	urlStr, urlStrErr := u.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/user/permissions/teams"), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := u.client.http.Get(urlStr, perms, nil)

	return perms, response, err
}
