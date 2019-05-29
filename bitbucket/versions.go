package bitbucket

// VersionsService handles communication with the version related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/versions
type VersionsService service

type Versions struct {
	Page     int    `json:"page,omitempty"`
	Next     string `json:"next,omitempty"`
	Pagelen  int    `json:"pagelen,omitempty"`
	Size     int    `json:"size,omitempty"`
	Previous string `json:"previous,omitempty"`

	Values []*Version `json:"values,omitempty"`
}

type Version struct {
	Repository *Repository   `json:"repository,omitempty"`
	Type       *string       `json:"type,omitempty"`
	Name       *string       `json:"name,omitempty"`
	Links      *VersionLinks `json:"links,omitempty"`
}

type VersionLinks struct {
	Self *BitbucketLink `json:"self,omitempty"`
}
