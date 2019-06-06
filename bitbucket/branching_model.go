package bitbucket

// BranchingModelService handles communication with the branching model related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branching-model
type BranchingModelService service

// BranchingModel represents the brancing model for a repository.
type BranchingModel struct {
	Development *BMBranch        `json:"development,omitempty"`
	BranchTypes []*BMBranchTypes `json:"values,omitempty"`
	Production  *BMBranch        `json:"production,omitempty"`
	Type        *string          `json:"type,omitempty"`
	Links       *BMLinks         `json:"links,omitempty"`
}

// BMBranch represents the git branches you want the model to be applied to.
type BMBranch struct {
	Name    *string `json:"name,omitempty"`
	IsValid *bool   `json:"is_valid,omitempty"`
	Branch  *struct {
		Type   *string `json:"type,omitempty"`
		Name   *string `json:"name,omitempty"`
		Target *struct {
			Hash *string `json:"hash,omitempty"`
		} `json:"target,omitempty"`
	} `json:"branch,omitempty"`
	UseMainbranch *bool `json:"use_mainbranch,omitempty"`
}

// BMBranchTypes represents the branch prefix configurations for new branches.
type BMBranchTypes struct {
	Kind    *string `json:"kind,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
	Prefix  *string `json:"prefix,omitempty"`
}

// BMLinks represents the "links" object in a Bitbucket branching model.
type BMLinks struct {
	Self *Link `json:"self,omitempty"`
}

// BranchModelRequest represents a request to update an existing branching model.
type BranchModelRequest struct {
	Development *BMBranchUpdateOpts `json:"development,omitempty"`
	Production  *BMBranchUpdateOpts `json:"production,omitempty"`
	BranchTypes []*BMBranchTypes    `json:"branch_types,omitempty"`
}

// BMBranchUpdateOpts represents the fields available when updating the development/production branches.
// Tips:
//  - If development/production branch are not using `master` and you wish to switch to it, you will need
//    to set `Name` field to an empty string along with `UseMainbranch: true`.
type BMBranchUpdateOpts struct {
	UseMainbranch *bool   `json:"use_mainbranch,omitempty"`
	Enabled       *bool   `json:"enabled,omitempty"`
	Name          *string `json:"name,omitempty"`
}

// Get returns the branching model as applied to the repository. This view is read-only.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branching-model
func (bm *BranchingModelService) Get(owner, repoSlug string, opts ...interface{}) (*BranchingModel, *Response, error) {
	result := new(BranchingModel)
	urlStr := bm.client.requestUrl("/repositories/%s/%s/issues/branching-model", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := bm.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// GetRaw returns the branching model's raw configuration for a repository.
//
// A client wishing to see the branching model with its actual current branches should use the 'Get' function above.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branching-model/settings#get
func (bm *BranchingModelService) GetRaw(owner, repoSlug string, opts ...interface{}) (*BranchingModel, *Response, error) {
	result := new(BranchingModel)
	urlStr := bm.client.requestUrl("/repositories/%s/%s/issues/branching-model/settings", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := bm.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// Update update the branching model configuration for a repository.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branching-model/settings#put
func (bm *BranchingModelService) Update(owner, repoSlug string, bo *BranchModelRequest, opts ...interface{}) (*BranchingModel, *Response, error) {
	result := new(BranchingModel)
	urlStr := bm.client.requestUrl("/repositories/%s/%s/issues/branching-model/settings", owner, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := bm.client.execute("PUT", urlStr, result, bo)

	return result, response, err
}
