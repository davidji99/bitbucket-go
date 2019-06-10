package bitbucket

import "time"

// UserHooks represent a user's hooks.
type UserHooks struct {
	PaginationInfo

	Values []*Issue `json:"values,omitempty"`
}

// UserHook represents a user hook.
type UserHook struct {
	UUID        *string    `json:"uuid,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Description *string    `json:"description,omitempty"`
	SubjectType *string    `json:"subject_type,omitempty"`
	Active      *bool      `json:"active,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Events      []*string  `json:"events,omitempty"`
}

// ListHooks fetches all hooks for a user.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/hooks#get
func (u *UsersService) ListHooks(userID string, opts ...interface{}) (*UserHooks, *Response, error) {
	hooks := new(UserHooks)
	urlStr := u.client.requestURL("/users/%s/hooks", userID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := u.client.execute("GET", urlStr, hooks, nil)

	return hooks, response, err
}

// GetHook fetches a single hook for a user.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/hooks/%7Buid%7D#get
func (u *UsersService) GetHook(userID, hookID string, opts ...interface{}) (*UserHook, *Response, error) {
	hook := new(UserHook)
	urlStr := u.client.requestURL("/users/%s/hooks/%s", userID, hookID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := u.client.execute("GET", urlStr, hook, nil)

	return hook, response, err
}

// DeleteHook deletes a single hook for a user.
//
// Accepts the user's UUID, account_id, or username. Recommend to use UUID or account_id.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/users/%7Busername%7D/hooks/%7Buid%7D#delete
func (u *UsersService) DeleteHook(userID, hookID string) (*Response, error) {
	urlStr := u.client.requestURL("/users/%s/hooks/%s", userID, hookID)
	response, err := u.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
