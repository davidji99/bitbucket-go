package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"net/url"
)

// DownloadsService handles communication with the downloads related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/downloads
type DownloadsService service

// Artifacts represents a collection of artifacts (or files).
type Artifacts struct {
	PaginationInfo

	Values []*Artifact `json:"values,omitempty"`
}

// Artifact represents a file on Bitbucket.
type Artifact struct {
	Name          *string            `json:"name,omitempty"`
	Links         *ArtifactFileLinks `json:"links,omitempty"`
	DownloadCount *int64             `json:"downloads,omitempty"`
	User          *User              `json:"user,omitempty"`
	Type          *string            `json:"type,omitempty"`
	Size          *int64             `json:"size,omitempty"`
}

// ArtifactFileLinks represents the "links" object in a Bitbucket artifact.
type ArtifactFileLinks struct {
	Self *Link `json:"self,omitempty"`
}

// List returns a list of download links associated with the repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/downloads#get
func (d *DownloadsService) List(owner, repoSlug string, opts ...interface{}) (*Artifacts, *simpleresty.Response, error) {
	downloads := new(Artifacts)
	urlStr, urlStrErr := d.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/downloads", owner, repoSlug), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := d.client.http.Get(urlStr, downloads, nil)

	return downloads, response, err
}

// Delete the specified download artifact from the repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/downloads/%7Bfilename%7D#delete
func (d *DownloadsService) Delete(owner, repoSlug, fileName string) (*simpleresty.Response, error) {
	escapedFilename := url.QueryEscape(fileName)
	urlStr := d.client.http.RequestURL("/repositories/%s/%s/downloads/%s", owner, repoSlug, escapedFilename)
	response, err := d.client.http.Delete(urlStr, nil, nil)

	return response, err
}
