package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

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
func (t *TeamsService) ListProjects(teamUsername string, opts ...interface{}) (*TeamProjects, *simpleresty.Response, error) {
	result := new(TeamProjects)
	urlStr, urlStrErr := t.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/teams/%s/projects/", teamUsername), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := t.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// CreateProject creates a new project.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/projects/#post
func (t *TeamsService) CreateProject(teamUsername string, po *TeamProjectRequest) (*TeamProject, *simpleresty.Response, error) {
	result := new(TeamProject)
	urlStr := t.client.http.RequestURL("/teams/%s/projects/", teamUsername) // Trailing slash is required!
	response, err := t.client.http.Post(urlStr, result, po)

	return result, response, err
}

// UpdateProject updates an existing project
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/projects/%7Bproject_key%7D#put
func (t *TeamsService) UpdateProject(teamUsername, projectKey string, po *TeamProjectRequest) (*TeamProject, *simpleresty.Response, error) {
	result := new(TeamProject)
	urlStr := t.client.http.RequestURL("/teams/%s/projects/%s", teamUsername, projectKey)
	response, err := t.client.http.Put(urlStr, result, po)

	return result, response, err
}

// DeleteProject deletes the specified project.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/projects/%7Bproject_key%7D#delete
func (t *TeamsService) DeleteProject(teamUsername, projectKey string) (*simpleresty.Response, error) {
	urlStr := t.client.http.RequestURL("/teams/%s/projects/%s", teamUsername, projectKey)
	response, err := t.client.http.Delete(urlStr, nil, nil)

	return response, err
}
