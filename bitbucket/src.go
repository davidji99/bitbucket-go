package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davidji99/simpleresty"
	"io"
	"net/url"
)

// SRCService handles communication with the src related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/src
type SRCService service

// SRCMetadata represents a Bitbucket file/folder's metadata on a repository.
type SRCMetadata struct {
	Mimetype   *string           `json:"mimetype,omitempty"`
	Links      *FileHistoryLinks `json:"links,omitempty"`
	Commit     *Commit           `json:"commit,omitempty"`
	Attributes []*string         `json:"attributes,omitempty"`
	Path       *string           `json:"path,omitempty"`
	Type       *string           `json:"type,omitempty"`
	Size       *int64            `json:"size,omitempty"`
}

// SRCGetOpts represents the query parameters available to SRC#Get requests.
type SRCGetOpts struct {
	// If provided, returns the contents of the repository and its subdirectories recursively
	// until the specified max_depth of nested directories. When omitted, this defaults to 1.
	MaxDepth int64 `url:"max_depth,omitempty"`
}

// srcFormatOpts represents the URL parameters to get the metadata.
//
// This is unexported by default in order to promote the distinction between the GetRaw and GetMetadata functions below.
type srcFormatOpts struct {
	Format string `url:"format,omitempty"`
}

// GetRaw retrieves the contents of a single file, or the contents of a directory at a specified revision.
//
// When path points to a file, this endpoint returns the raw contents. When path points to a directory instead of a file,
// the response is a paginated list of directory and file objects in the same order as the underlying SCM system would return them.
//
// Bitbucket API docs:https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/src/%7Bnode%7D/%7Bpath%7D#get
func (s *SRCService) GetRaw(owner, repoSlug, nodeRev, path string,
	opts ...interface{}) (fileContent *bytes.Buffer, folderContent *FileHistory, resp *simpleresty.Response, err error) {

	encPath := (&url.URL{Path: path}).String()
	urlStr, urlStrErr := s.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/src/%s/%s", owner, repoSlug, nodeRev, encPath), opts...)
	if urlStrErr != nil {
		return nil, nil, nil, urlStrErr
	}

	req := s.client.http.NewRequest()
	req.Method = simpleresty.GetMethod
	req.URL = urlStr

	resp, reqErr := s.client.http.Dispatch(req)
	if reqErr != nil {
		return nil, nil, nil, reqErr
	}

	// Attempt to unmarshal the response body as folderContent.
	// If it works, return as folderContent.
	decErr := json.NewDecoder(resp.Resp.RawResponse.Body).Decode(&folderContent)
	if decErr == io.EOF {
		decErr = nil
	}
	if decErr == nil {
		return nil, folderContent, resp, nil
	}

	// If the unmarshal above doesn't work, raw file content was returned by the API.
	// Handle and parse the raw content.
	var buff bytes.Buffer
	_, parseRawErr := io.Copy(&buff, resp.Resp.RawResponse.Body)
	if parseRawErr == nil {
		fileContent = &buff
		return fileContent, nil, resp, nil
	}

	// Return generic error if all else fails.
	return nil, nil, resp, fmt.Errorf("unable to get the raw or json content for your request")
}

// GetMetadata returns the JSON object describing the file or folder's properties,
// instead of returning the raw contents.
//
// Supports the Bitbucket querying/filtering syntax and so you could filter a directory
// listing to only include entries that match certain criteria.
//
// Bitbucket API docs:https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/src/%7Bnode%7D/%7Bpath%7D#get
func (s *SRCService) GetMetadata(owner, repoSlug, nodeRev, path string, opts ...interface{}) (*SRCMetadata, *simpleresty.Response, error) {
	result := new(SRCMetadata)
	encPath := (&url.URL{Path: path}).String()

	// Add format=meta URL parameter by default
	formatQueryParam := &srcFormatOpts{Format: "meta"}
	opts = append(opts, formatQueryParam)

	urlStr, urlStrErr := s.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/src/%s/%s", owner, repoSlug, nodeRev, encPath), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := s.client.http.Get(urlStr, result, nil)

	return result, response, err
}
