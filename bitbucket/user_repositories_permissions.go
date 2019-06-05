package bitbucket

// UserRepositoriesPermissions represents a collection of a user's permissions on repositories.
type UserRepositoriesPermissions struct {
	PaginationInfo

	Values []*UserRepositoriesPermission `json:"values,omitempty"`
}

type UserRepositoriesPermission struct {
	Type       *string     `json:"type,omitempty"`
	User       *User       `json:"user,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Permission *string     `json:"permission,omitempty"`
}

// ListReposPerms returns permissions for each repository the caller has explicit access to and the highest level of permission the caller has.
//
// This does not return public repositories that the user was not granted any specific permission in,
// and does not distinguish between direct and indirect privileges.
//
// Supports filtering by passing in a non-URI encoded query string. Reference: https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/user/permissions/repositories#get
func (u *UserService) ListReposPerms(opts ...interface{}) (*UserRepositoriesPermissions, *Response, error) {
	perms := new(UserRepositoriesPermissions)
	urlStr := u.client.requestUrl("/user/permissions/repositories")
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := u.client.execute("GET", urlStr, perms, nil)

	return perms, response, err
}
