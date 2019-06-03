package bitbucket

import "net/url"

// DownloadsService handles communication with the downloads related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/downloads
type DownloadsService service

type Downloads struct {
	Pagination

	Values []*DownloadArtifact `json:"values,omitempty"`
}

type DownloadArtifact struct {
	Name          *string            `json:"name,omitempty"`
	Links         *DownloadFileLinks `json:"links,omitempty"`
	DownloadCount *int64             `json:"downloads,omitempty"`
	User          *User              `json:"user,omitempty"`
	Type          *string            `json:"type,omitempty"`
	Size          *int64             `json:"size,omitempty"`
}

type DownloadFileLinks struct {
	Self *BitbucketLink `json:"self,omitempty"`
}

// List returns a list of download links associated with the repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/downloads#get
func (d *DownloadsService) List(owner, repoSlug string) (*Downloads, *Response, error) {
	downloads := new(Downloads)
	urlStr := d.client.requestUrl("/repositories/%s/%s/downloads", owner, repoSlug)
	response, err := d.client.execute("GET", urlStr, downloads, nil)

	return downloads, response, err
}

// Delete the specified download artifact from the repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/downloads/%7Bfilename%7D#delete
func (d *DownloadsService) Delete(owner, repoSlug, fileName string) (*Response, error) {
	encodedFileName := url.QueryEscape(fileName) // Encode the file name
	urlStr := d.client.requestUrl("/repositories/%s/%s/downloads/%s", owner, repoSlug, encodedFileName)
	response, err := d.client.execute("DELETE", urlStr, nil, nil)

	return response, err
}

// TODO: implement create, get#single
