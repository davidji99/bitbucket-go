package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

const componentSelfURLRegex = `http[sS]?:\/\/.*\/2.0\/repositories\/.*\/.*\/components/(\d+)`

// ComponentsService handles communication with the user related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/components
type ComponentsService service

// Components represent a collection of components.
type Components struct {
	PaginationInfo

	Values []*Component `json:"values,omitempty"`
}

// Component represents a Bitbucket repository component.
type Component struct {
	ID         *int64          `json:"-"` // This field is not present in the API response.
	Repository *Repository     `json:"repository,omitempty"`
	Type       *string         `json:"type,omitempty"`
	Name       *string         `json:"name,omitempty"`
	Links      *ComponentLinks `json:"links,omitempty"`
}

// ComponentLinks represents the "links" object in a Bitbucket component.
type ComponentLinks struct {
	Self *Link `json:"self,omitempty"`
}

// ComponentRequest represents an existing component to be added to an issue or pull request.
// There is no CREATE or UPDATE endpoint for the component resource.
type ComponentRequest struct {
	Name *string `json:"name,omitempty"`
}

// List all components that have been defined in the issue tracker.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/components#get
func (c *ComponentsService) List(owner, repoSlug string, opts ...interface{}) (*Components, *simpleresty.Response, error) {
	result := new(Components)
	urlStr, urlStrErr := c.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/components", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := c.client.http.Get(urlStr, result, nil)

	// Parse and store the component id
	for _, component := range result.Values {
		component.ID = parseForResourceID(componentSelfURLRegex, *component.Links.Self.HRef)
	}

	return result, response, err
}

// Get a single component.
// NOTE: The component ID is a numerical value, not the component name, that is visible in the links.self.href object.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/components/%7Bcomponent_id%7D#get
func (c *ComponentsService) Get(owner, repoSlug string, componentID int64, opts ...interface{}) (*Component, *simpleresty.Response, error) {
	component := new(Component)
	urlStr, urlStrErr := c.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/components/%v", owner, repoSlug, componentID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := c.client.http.Get(urlStr, component, nil)

	// Parse and store the component id
	component.ID = parseForResourceID(componentSelfURLRegex, *component.Links.Self.HRef)

	return component, response, err
}
