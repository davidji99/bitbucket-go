package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// UserEmails represents a collection of user emails.
type UserEmails struct {
	PaginationInfo

	Values []*UserEmail `json:"values,omitempty"`
}

// UserEmail represents an individual user's email address.
type UserEmail struct {
	IsPrimary   *bool           `json:"is_primary,omitempty"`
	IsConfirmed *bool           `json:"is_confirmed,omitempty"`
	Type        *string         `json:"type,omitempty"`
	Email       *string         `json:"email,omitempty"`
	Links       *UserEmailLinks `json:"links,omitempty"`
}

// UserEmailLinks represents the "links" object in a Bitbucket user email.
type UserEmailLinks struct {
	Self *Link `json:"self,omitempty"`
}

// GetEmails returns all the authenticated user's email addresses. Both confirmed and unconfirmed.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/user/emails/%7Bemail%7D#get
func (u *UserService) GetEmails(opts ...interface{}) (*UserEmails, *simpleresty.Response, error) {
	emails := new(UserEmails)
	urlStr, urlStrErr := u.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/user/emails"), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := u.client.http.Get(urlStr, emails, nil)

	return emails, response, err
}
