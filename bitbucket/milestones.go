package bitbucket

const milestoneSelfUrl = `http[sS]?:\/\/.*\/2.0\/repositories\/.*\/.*\/milestones/(\d+)`

// MilestonesService handles communication with the milestone related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/milestones
type MilestonesService service

// Milestones represent a collection of milestones.
type Milestones struct {
	Pagination

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
	Self *BitbucketLink `json:"self,omitempty"`
}

// MilestoneRequest represents an existing milestone to be added to an issue or pull request.
// There is no CREATE or UPDATE endpoint for the milestone resource.
type MilestoneRequest struct {
	Name *string `json:"name,omitempty"`
}

// List all milestones that have been defined in the issue tracker.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/milestones#get
func (m *MilestonesService) List(owner, repoSlug string) (*Milestones, *Response, error) {
	milestones := new(Milestones)
	urlStr := m.client.requestUrl("/repositories/%s/%s/milestones", owner, repoSlug)
	response, err := m.client.execute("GET", urlStr, milestones, nil)

	// Parse and store the milestonoe id
	for _, milestone := range milestones.Values {
		milestone.ID = parseForResourceId(milestoneSelfUrl, *milestone.Links.Self.HRef)
	}

	return milestones, response, err
}

// Get a single milestone.
// NOTE: The milestone ID is a numerical value, not the component name, that is visible in the links.self.href object.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/milestones/%7Bmilestone_id%7D#get
func (m *MilestonesService) Get(owner, repoSlug string, milestoneID int64) (*Milestone, *Response, error) {
	milestone := new(Milestone)
	urlStr := m.client.requestUrl("/repositories/%s/%s/milestones/%v", owner, repoSlug, milestoneID)
	response, err := m.client.execute("GET", urlStr, milestone, nil)

	// Parse and store the milestone id
	milestone.ID = parseForResourceId(milestoneSelfUrl, *milestone.Links.Self.HRef)

	return milestone, response, err
}
