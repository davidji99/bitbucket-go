package bitbucket

import "time"

// UsersSSHKeys represents a collection of user ssh keys.
type UsersSSHKeys struct {
	PaginationInfo

	Values []*UsersSSHKey `json:"values,omitempty"`
}

// UsersSSHKey represents a user ssh key added to Bitbucket.
type UsersSSHKey struct {
	Comment   *string           `json:"comment,omitempty"`
	CreatedOn *time.Time        `json:"created_on,omitempty"`
	Key       *string           `json:"key,omitempty"`
	Label     *string           `json:"labels,omitempty"`
	LastUsed  *time.Time        `json:"last_used,omitempty"`
	Links     *UsersSSHKeyLinks `json:"links,omitempty"`
	Owner     *User             `json:"owner,omitempty"`
	Type      *string           `json:"values,omitempty"`
	UUID      *string           `json:"uuid,omitempty"`
}

// UsersSSHKeyLinks represents the "links" object in a Bitbucket user ssh key.
type UsersSSHKeyLinks struct {
	Self *Link `json:"self,omitempty"`
}

// SSHKeyAddRequest represents a request to add a SSH key.
type SSHKeyAddRequest struct {
	Key   *string `json:"key,omitempty"`
	Label *string `json:"label,omitempty"`
}

// ListSSHKeys returns a paginated list of the user's SSH public keys.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/ssh-keys#get
func (u *UsersService) ListSSHKeys(userID string, opts ...interface{}) (*UsersSSHKeys, *Response, error) {
	sshKeys := new(UsersSSHKeys)
	urlStr := u.client.requestUrl("/users/%s/ssh-keys", userID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := u.client.execute("GET", urlStr, sshKeys, nil)

	return sshKeys, response, err
}

// AddSSHKey adds a new SSH public key to the specified user account and returns the resulting key.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/ssh-keys#post
func (u *UsersService) AddSSHKey(userID string, newKey *SSHKeyAddRequest) (*UsersSSHKey, *Response, error) {
	sshKey := new(UsersSSHKey)
	urlStr := u.client.requestUrl("/users/%s/ssh-keys", userID)
	response, err := u.client.execute("POST", urlStr, sshKey, newKey)

	return sshKey, response, err
}
