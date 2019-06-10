package bitbucket

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
func (u *UserService) GetEmails(opts ...interface{}) (*UserEmails, *Response, error) {
	emails := new(UserEmails)
	urlStr := u.client.requestUrl("/user/emails")
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := u.client.execute("GET", urlStr, emails, nil)

	return emails, response, err
}
