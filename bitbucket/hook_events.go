package bitbucket

import (
	"fmt"
	"github.com/davidji99/simpleresty"
)

// HookEventsService handles communication with the hook events related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/hook_events
type HookEventsService service

// HookEvents represents a collection of hook events.
type HookEvents struct {
	PaginationInfo

	Values []*HookEvent `json:"values,omitempty"`
}

// HookEvent represents a hook event.
type HookEvent struct {
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	Event       *string `json:"event,omitempty"`
	Label       *string `json:"label,omitempty"`
}

// HookEventTypes represents hook events.
type HookEventTypes struct {
	Repository *HookEventTypesLinks `json:"repository,omitempty"`
	Team       *HookEventTypesLinks `json:"team,omitempty"`
	User       *HookEventTypesLinks `json:"user,omitempty"`
}

// HookEventTypesLinks represents the "links" object in a Bitbucket hook event type.
type HookEventTypesLinks struct {
	Events *Link `json:"events,omitempty"`
}

// List returns the webhook resource or subject types on which webhooks can be registered.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/hook_events#get
func (h *HookEventsService) List(opts ...interface{}) (*HookEventTypes, *simpleresty.Response, error) {
	result := new(HookEventTypes)
	urlStr, urlStrErr := h.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/hook_events"), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := h.client.http.Get(urlStr, result, nil)

	return result, response, err
}

// Get returns a paginated list of all valid webhook events for the specified entity.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/hook_events/%7Bsubject_type%7D#get
func (h *HookEventsService) Get(subjectType string, opts ...interface{}) (*HookEvents, *simpleresty.Response, error) {
	result := new(HookEvents)
	urlStr, urlStrErr := h.client.http.RequestURLWithQueryParams(
		fmt.Sprintf("/hook_events/%s", subjectType), opts...)
	if urlStrErr != nil {
		return nil, nil, urlStrErr
	}

	response, err := h.client.http.Get(urlStr, result, nil)

	return result, response, err
}
