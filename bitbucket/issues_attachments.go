package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"net/url"
)

// ListAttachments returns all attachments for this issue.
//
// This returns the files' meta data. This does not return the files' actual contents.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/attachments#get
func (i *IssuesService) ListAttachments(owner, repoSlug string, id int64, opts ...interface{}) (*Artifacts, *simpleresty.Response, error) {
	result := new(Artifacts)
	urlStr, urlStrErr := i.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/repositories/%s/%s/issues/%v/attachments", owner, repoSlug, id), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := i.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// DeleteAttachment deletes an attachment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/issues/%7Bissue_id%7D/attachments/%7Bpath%7D#delete
func (i *IssuesService) DeleteAttachment(owner, repoSlug string, id int64, filePath string) (*simpleresty.Response, error) {
	escFilePath := url.QueryEscape(filePath)
	urlStr := i.client.http.RequestURL("/repositories/%s/%s/issues/%v/attachments/%s", owner, repoSlug, id, escFilePath)
	response, err := i.client.http.Delete(urlStr, nil, nil)

	return response, err
}
