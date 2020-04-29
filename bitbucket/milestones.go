package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

const milestoneSelfURL = `http[sS]?:\/\/.*\/2.0\/repositories\/.*\/.*\/milestones/(\d+)`

// MilestonesService handles communication with the milestone related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/milestones
type MilestonesService service

// Milestones represent a collection of milestones.
type Milestones struct {
	PaginationInfo

	Values []*Milestone `json:"values,omitempty"`
}

// Milestone represents a Bitbucket repository milestone.
type Milestone struct {
	ID         *int64          `json:"-"` // This field is not present in the API response.
	Repository *Repository     `json:"repository,omitempty"`
	Type       *string         `json:"type,omitempty"`
	Name       *string         `json:"name,omitempty"`
	Links      *MilestoneLinks `json:"links,omitempty"`
}

// MilestoneLinks represents the "links" object in a Bitbucket milestone.
type MilestoneLinks struct {
	Self *Link `json:"self,omitempty"`
}

// MilestoneRequest represents an existing milestone to be added to an issue or pull request.
// There is no CREATE or UPDATE endpoint for the milestone resource.
type MilestoneRequest struct {
	Name *string `json:"name,omitempty"`
}

// List all milestones that have been defined in the issue tracker.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/milestones#get
func (m *MilestonesService) List(owner, repoSlug string, opts ...interface{}) (*Milestones, *simpleresty.Response, error) {
	result := new(Milestones)
	urlStr, urlStrErr := m.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/milestones", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := m.client.http.Get(urlStr, result, nil)

	// Parse and store the milestone id
	for _, milestone := range result.Values {
		milestone.ID = parseForResourceID(milestoneSelfURL, *milestone.Links.Self.HRef)
	}

	return result, response, err
}

// Get returns a single milestone.
// NOTE: The milestone ID is a numerical value, not the component name, that is visible in the links.self.href object.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/milestones/%7Bmilestone_id%7D#get
func (m *MilestonesService) Get(owner, repoSlug string, milestoneID int64, opts ...interface{}) (*Milestone, *simpleresty.Response, error) {
	result := new(Milestone)
	urlStr, urlStrErr := m.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/milestones/%v", owner, repoSlug, milestoneID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := m.client.http.Get(urlStr, result, nil)

	// Parse and store the milestone id
	result.ID = parseForResourceID(milestoneSelfURL, *result.Links.Self.HRef)

	return result, response, err
}
