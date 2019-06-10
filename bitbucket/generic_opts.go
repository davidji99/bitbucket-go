package bitbucket

// GenericOpts represents all of the available generic query parameters Bitbucket has to offer.
//
// Each query parameter may or may not work so your mileage may vary for certain resources.
// This struct is available out of convenience.
type GenericOpts struct {
	ListOpts
	FilterSortOpts
	PartialRespOpts
}

// ListOpts specifies the optional parameters to various List methods that support pagination.
type ListOpts struct {
	// For paginated result sets, page of results to retrieve.
	Page int64 `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	// Globally, the minimum length is 10 and the maximum is 100. Some APIs may specify a different default.
	Pagelen int64 `url:"pagelen,omitempty"`
}

// FilterSortOpts represents the querying and sorting mechanism available
// to certain Bitbucket API resources that return multiple results in a response.
//
// Bitbucket API Docs: https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering#query-sort
type FilterSortOpts struct {
	// Query is the raw non-URL encoded query string.
	// Note that the entire query string is put in the Query field.
	// This library will take care of URL encoding the string for you.
	Query string `url:"q,omitempty"`

	// In principle, every field that can be queried can also be used as a key for sorting.
	// By default the sort order is ascending. To reverse the order, prefix the field name with a hyphen (e.g. ?sort=-updated_on).
	// Only one field can be sorted on. Compound fields (e.g. sort on state first, followed by updated_on) are not supported.
	Sort string `url:"sort,omitempty"`
}

// PartialRespOpts represents the URL parameter to request a partial response and to add or remove
// specific fields from a response.
//
// Bitbucket API Docs: https://developer.atlassian.com/bitbucket/api/2/reference/meta/partial-response
type PartialRespOpts struct {
	// The fields parameter can contain a list of multiple comma-separated field names (e.g. fields=owner.username,uuid,links.self.href).
	Fields string `url:"fields,omitempty"`
}
