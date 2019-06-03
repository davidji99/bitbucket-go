package bitbucket

// SearchCodeResults represents the results from a search query.
type SearchCodeResults struct {
	Pagination

	QuerySubstituted *bool             `json:"query_substituted,omitempty"`
	Values           *SearchCodeResult `json:"values,omitempty"`
}

// SearchCodeResult represents the individual search query result.
type SearchCodeResult struct {
	Type              *string               `json:"type,omitempty"`
	ContentMatchCount *int64                `json:"content_match_count,omitempty"`
	ContentMatches    []*SearchContentMatch `json:"content_matches,omitempty"`
	PathMatches       []*SearchMatch        `json:"path_matches,omitempty"`
}

// SearchContentMatch represents the content code lines that match a search result.
type SearchContentMatch struct {
	Lines *SearchContentMatchLine `json:"lines,omitempty"`
}

// SearchContentMatchLine represents the specific line(s) that match a content result.
type SearchContentMatchLine struct {
	Line     *int64         `json:"line,omitempty"`
	Segments []*SearchMatch `json:"segments,omitempty"`
}

// SearchMatch represents the content of a search result code line.
type SearchMatch struct {
	Text  *string `json:"text,omitempty"`
	Match *bool   `json:"match,omitempty"`
}

// SearchCodeFile represents the information regarding the file that matched a search query.
type SearchCodeFile struct {
	Path  *string              `json:"path,omitempty"`
	Type  *string              `json:"type,omitempty"`
	Links *SearchCodeFileLinks `json:"links,omitempty"`
}

// SearchCodeFileLinks represents the "links" object in a Bitbucket search code result file.
type SearchCodeFileLinks struct {
	Self         *BitbucketLink `json:"self,omitempty"`
	Repositories *BitbucketLink `json:"repositories,omitempty"`
}

// CodeSearchOpts represents the query parameters available when searching for code.
type CodeSearchOpts struct {
	// The search query
	SearchQuery string `url:"search_query,omitempty"`

	// Which page of the search results to retrieve
	Page string `url:"page,omitempty"`

	// How many search results to retrieve per page
	Pagelen string `url:"pagelen,omitempty"`
}
