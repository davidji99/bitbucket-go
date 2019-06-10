package bitbucket

import "net/url"

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
func (d *DownloadsService) List(owner, repoSlug string, opts ...interface{}) (*Artifacts, *Response, error) {
	downloads := new(Artifacts)
	urlStr := d.client.requestUrl("/repositories/%s/%s/downloads", owner, repoSlug)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := d.client.execute("GET", urlStr, downloads, nil)

	return downloads, response, err
}

// Delete the specified download artifact from the repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/downloads/%7Bfilename%7D#delete
func (d *DownloadsService) Delete(owner, repoSlug, fileName string) (*Response, error) {
	escapedFilename := url.QueryEscape(fileName)
	urlStr := d.client.requestUrl("/repositories/%s/%s/downloads/%s", owner, repoSlug, escapedFilename)
	response, err := d.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}
