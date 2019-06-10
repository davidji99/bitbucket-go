package bitbucket

import "time"

// TeamProjects represents a collection of team projects.
type TeamProjects struct {
	PaginationInfo

	Values []*TeamProject `json:"values,omitempty"`
}

// TeamProject represents a team project in Bitbucket.
type TeamProject struct {
	UUID        *string           `json:"uuid,omitempty"`
	Links       *TeamProjectLinks `json:"links,omitempty"`
	Description *string           `json:"description,omitempty"`
	CreatedOn   *time.Time        `json:"created_on,omitempty"`
	Key         *string           `json:"key,omitempty"`
	Owner       *User             `json:"owner,omitempty"`
	UpdatedOn   *time.Time        `json:"updated_on,omitempty"`
	Type        *string           `json:"type,omitempty"`
	IsPrivate   *bool             `json:"is_private,omitempty"`
	Name        *string           `json:"name,omitempty"`
}

// TeamProjectLinks represents the "links" object in a Bitbucket team project.
type TeamProjectLinks struct {
	Self   *Link `json:"self,omitempty"`
	HTML   *Link `json:"html,omitempty"`
	Avatar *Link `json:"avatar,omitempty"`
}

// TeamProjectRequest represents a request to create/update a team project.
type TeamProjectRequest struct {
	Name        *string `json:"name,omitempty"`
	Key         *string `json:"key,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPrivate   *bool   `json:"is_private,omitempty"`
}

// ListProjects returns each project a team has.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/projects/#get
func (t *TeamsService) ListProjects(teamUsername string, opts ...interface{}) (*TeamProjects, *Response, error) {
	result := new(TeamProjects)
	urlStr := t.client.requestURL("/teams/%s/projects/", teamUsername) // Trailing slash is required!
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// CreateProject creates a new project.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/projects/#post
func (t *TeamsService) CreateProject(teamUsername string, po *TeamProjectRequest) (*TeamProject, *Response, error) {
	result := new(TeamProject)
	urlStr := t.client.requestURL("/teams/%s/projects/", teamUsername) // Trailing slash is required!
	response, err := t.client.execute("POST", urlStr, result, po)

	return result, response, err
}

// UpdateProject updates an existing project
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/projects/%7Bproject_key%7D#put
func (t *TeamsService) UpdateProject(teamUsername, projectKey string, po *TeamProjectRequest) (*TeamProject, *Response, error) {
	result := new(TeamProject)
	urlStr := t.client.requestURL("/teams/%s/projects/%s", teamUsername, projectKey)
	response, err := t.client.execute("PUT", urlStr, result, po)

	return result, response, err
}

// DeleteProject deletes the specified project.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/projects/%7Bproject_key%7D#delete
func (t *TeamsService) DeleteProject(teamUsername, projectKey string) (*Response, error) {
	urlStr := t.client.requestURL("/teams/%s/projects/%s", teamUsername, projectKey)
	response, err := t.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
