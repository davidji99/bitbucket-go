package bitbucket

import (
	"reflect"
	"time"
)

// IssueChanges represents a collection of changes on an issue,
type IssueChanges struct {
	PaginationInfo

	Values []*IssueChange `json:"values,omitempty"`
}

// IssueChange represents the individual change.
type IssueChange struct {
	ID        *int64                        `json:"id,omitempty"`
	Links     *IssueChangeLinks             `json:"links,omitempty"`
	Issue     *Issue                        `json:"issue,omitempty"`
	CreatedOn *time.Time                    `json:"created_on,omitempty"`
	User      *User                         `json:"user,omitempty"`
	Message   *Content                      `json:"message,omitempty"`
	Type      *string                       `json:"type,omitempty"`
	Changes   *map[string]map[string]string `json:"changes,omitempty"`
}

// IssueChangeLinks represents the "links" object in a Bitbucket issue change.
type IssueChangeLinks struct {
	Self *Link `json:"self,omitempty"`
	HTML *Link `json:"html,omitempty"`
}

// IssueChangeRequest represents a request to create change on an issue.
type IssueChangeRequest struct {
	AssigneeAccountID *string
	Kind              *string
	Priority          *string
	Component         *string
	Milestone         *string
	Version           *string
	Content           *string // represents the issue's description box.
	Message           *string
}

// ListChanges returns the list of all changes that have been made to the specified issue.
// Changes are returned in chronological order with the oldest change first.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/changes#get
func (i *IssuesService) ListChanges(owner, repoSlug string, id int64, opts ...interface{}) (*IssueChanges, *Response, error) {
	result := new(IssueChanges)
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/changes", owner, repoSlug, id)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := i.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// GetChange returns the specified issue change object.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/changes/%7Bchange_id%7D#get
func (i *IssuesService) GetChange(owner, repoSlug string, id, changeID int64, opts ...interface{}) (*IssueChange, *Response, error) {
	result := new(IssueChange)
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/changes/%v", owner, repoSlug, id, changeID)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := i.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// CreateChange makes a change to the specified issue.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/changes#post
func (i *IssuesService) CreateChange(owner, repoSlug string, id int64, io *IssueChangeRequest) (*IssueChange, *Response, error) {
	result := new(IssueChange)
	urlStr := i.client.requestURL("/repositories/%s/%s/issues/%v/changes", owner, repoSlug, id)
	response, err := i.client.execute("POST", urlStr, result, io.buildChangeRequestBody())

	return result, response, err
}

// issueChangeRequestBody represents the raw body for creating an issue change request.
type issueChangeRequestBody map[string]interface{}

// buildChangeRequestBody populates a map with the values from IssueChangeRequest
// and returns a nested map.
func (i *IssueChangeRequest) buildChangeRequestBody() *issueChangeRequestBody {
	body := make(issueChangeRequestBody)

	v := reflect.Indirect(reflect.ValueOf(i))
	t := reflect.TypeOf(*i)

	changes := map[string]map[string]string{}
	for i := 0; i < v.Type().NumField(); i++ {
		fieldName := toSnakeCase(t.Field(i).Name)
		fieldValue := reflect.Indirect(v.Field(i))

		if fieldValue.IsValid() && fieldName != "message" {
			changes[toSnakeCase(fieldName)] = map[string]string{"new": fieldValue.String()}
		}
	}
	body["changes"] = changes

	// Add the message to the body if not empty. We don't do this above because it's not nested under "changes".
	if i.GetMessage() != "" {
		body["message"] = map[string]string{"raw": i.GetMessage()}
	}

	return &body
}

// GetChanges returns a map of all changes for the issue.
func (i *IssueChange) GetChanges() map[string]map[string]string {
	if i == nil || i.Changes == nil {
		return make(map[string]map[string]string, 0)
	}

	return *i.Changes
}
