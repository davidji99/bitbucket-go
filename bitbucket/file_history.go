package bitbucket

// FileHistoryService handles communication with the file history related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/filehistory
type FileHistoryService service

// FileHistory represents the history of a file.
type FileHistory struct {
	PaginationInfo

	Values []*SRCMetadata `json:"values,omitempty"`
}

// FileHistoryLinks represents the "links" object in a Bitbucket file history.
type FileHistoryLinks struct {
	Self    *Link `json:"self,omitempty"`
	Meta    *Link `json:"meta,omitempty"`
	History *Link `json:"history,omitempty"`
}

// FileHistoryListOpts represents the unique query parameters for file history.
type FileHistoryListOpts struct {
	// When true, Bitbucket will follow the history of the file across renames (this is the default behavior).
	// This can be turned off by specifying false.
	Renames bool `url:"renames,omitempty"`
}

// Get returns a paginated list of commits that modified the specified file.
// Commits are returned in reverse chronological order.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/filehistory/%7Bnode%7D/%7Bpath%7D#get
func (fh *FileHistoryService) Get(owner, repoSlug, nodeRev, path string, opts ...interface{}) (*FileHistory, *Response, error) {
	result := new(FileHistory)
	urlStr := fh.client.requestUrl("/repositories/%s/%s/filehistory/%s/%s", owner, repoSlug, nodeRev, path)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := fh.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
