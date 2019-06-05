package bitbucket

import "time"

// TeamsService handles communication with the teams related methods
// of the Bitbucket API.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams
type TeamsService service

// Teams represents a collection of teams.
type Teams struct {
	PaginationInfo

	Values []*Team `json:"values,omitempty"`
}

// Team represents a Bitbucket team.
type Team struct {
	Username      *string    `json:"username,omitempty"`
	Nickname      *string    `json:"nickname,omitempty"`
	AccountStatus *string    `json:"account_status,omitempty"`
	DisplayName   *string    `json:"display_name,omitempty"`
	CreatedOn     *time.Time `json:"created_on,omitempty"`
	UUID          *string    `json:"uuid,omitempty"`
	Has2FAEnabled *string    `json:"has_2fa_enabled,omitempty"`
	Website       *string    `json:"website,omitempty"`
	Links         *TeamLinks `json:"links,omitempty"`
}

// TeamLinks represents the "links" object in a Bitbucket team.
type TeamLinks struct {
	Self         *BitbucketLink `json:"self,omitempty"`
	Hooks        *BitbucketLink `json:"hooks,omitempty"`
	Repositories *BitbucketLink `json:"repositories,omitempty"`
	Followers    *BitbucketLink `json:"followers,omitempty"`
	HTML         *BitbucketLink `json:"html,omitempty"`
	Avatar       *BitbucketLink `json:"avatar,omitempty"`
	Following    *BitbucketLink `json:"following,omitempty"`
	Members      *BitbucketLink `json:"members,omitempty"`
	Projects     *BitbucketLink `json:"projects,omitempty"`
	Snippets     *BitbucketLink `json:"snippets,omitempty"`
}

type TeamListOpts struct {
	// Filters the teams based on the authenticated user's role on each team:
	//  - member: returns a list of all the teams which the caller is a member of at least one team group or repository owned by the team.
	//  - contributor: returns a list of teams which the caller has write access to at least one repository owned by the team.
	//  - admin: returns a list teams which the caller has team administrator access.
	Role string `url:"role,omitempty"`

	ListOpts
}

// List returns all the teams that the authenticated user is associated with.
//
// Requires 'role' query parameter to be set.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams#get
func (t *TeamsService) List(opts ...interface{}) (*Teams, *Response, error) {
	teams := new(Teams)
	urlStr := t.client.requestUrl("/teams")
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, teams, nil)

	return teams, response, err
}

// Get the public information associated with a team.
//
// If the team's profile is private, location, website and created_on elements are omitted.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D#get
func (t *TeamsService) Get(teamUsername string) (*Team, *Response, error) {
	team := new(Team)
	urlStr := t.client.requestUrl("/teams/%s", teamUsername)
	response, err := t.client.execute("GET", urlStr, team, nil)

	return team, response, err
}
