package bitbucket

// ComponentsService handles communication with the user related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/components
type ComponentsService service

// Components represent a collection of issues.
type Components struct {
	Page     int    `json:"page,omitempty"`
	Next     string `json:"next,omitempty"`
	Pagelen  int    `json:"pagelen,omitempty"`
	Size     int    `json:"size,omitempty"`
	Previous string `json:"previous,omitempty"`

	Values []*Component `json:"values,omitempty"`
}

type Component struct {
	Repository *Repository     `json:"repository,omitempty"`
	Type       *string         `json:"type,omitempty"`
	Name       *string         `json:"name,omitempty"`
	Links      *ComponentLinks `json:"links,omitempty"`
}

type ComponentLinks struct {
	Self *BitbucketLink `json:"self,omitempty"`
}
