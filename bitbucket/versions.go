package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

const versionSelfURLRegex = `http[sS]?:\/\/.*\/2.0\/repositories\/.*\/.*\/versions/(\d+)`

// VersionsService handles communication with the version related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/versions
type VersionsService service

// Versions represents a collection of versions.
type Versions struct {
	PaginationInfo

	Values []*Version `json:"values,omitempty"`
}

// Version represents a Bitbucket repository version.
type Version struct {
	ID         *int64        `json:"-"` // This field is not present in the API response.
	Repository *Repository   `json:"repository,omitempty"`
	Type       *string       `json:"type,omitempty"`
	Name       *string       `json:"name,omitempty"`
	Links      *VersionLinks `json:"links,omitempty"`
}

// VersionLinks represents the "links" object in a Bitbucket version.
type VersionLinks struct {
	Self *Link `json:"self,omitempty"`
}

// VersionRequest represents an EXISTING version to be added to an issue or pull request.
// There is no CREATE or UPDATE endpoint for the version resource.
type VersionRequest struct {
	Name *string `json:"name,omitempty"`
}

// List all versions that have been defined in the issue tracker.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/versions#get
func (v *VersionsService) List(owner, repoSlug string, opts ...interface{}) (*Versions, *simpleresty.Response, error) {
	versions := new(Versions)
	urlStr, urlStrErr := v.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/versions", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := v.client.http.Get(urlStr, versions, nil)

	// Parse and store the version id
	for _, version := range versions.Values {
		version.ID = parseForResourceID(versionSelfURLRegex, *version.Links.Self.HRef)
	}

	return versions, response, err
}

// Get a single version.
// NOTE: The version ID is a numerical value, not the version name, that is visible in the links.self.href object.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/versions/%7Bversion_id%7D#get
func (v *VersionsService) Get(owner, repoSlug string, versionID int64, opts ...interface{}) (*Version, *simpleresty.Response, error) {
	version := new(Version)
	urlStr, urlStrErr := v.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/versions/%v", owner, repoSlug, versionID), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := v.client.http.Get(urlStr, version, nil)

	// Parse and store the version id
	version.ID = parseForResourceID(versionSelfURLRegex, *version.Links.Self.HRef)

	return version, response, err
}
