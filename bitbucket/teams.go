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
	Self         *Link `json:"self,omitempty"`
	Hooks        *Link `json:"hooks,omitempty"`
	Repositories *Link `json:"repositories,omitempty"`
	Followers    *Link `json:"followers,omitempty"`
	HTML         *Link `json:"html,omitempty"`
	Avatar       *Link `json:"avatar,omitempty"`
	Following    *Link `json:"following,omitempty"`
	Members      *Link `json:"members,omitempty"`
	Projects     *Link `json:"projects,omitempty"`
	Snippets     *Link `json:"snippets,omitempty"`
}

type TeamListOpts struct {
	// Filters the teams based on the authenticated user's role on each team:
	//  - member: returns a list of all the teams which the caller is a member of at least one team group or repository owned by the team.
	//  - contributor: returns a list of teams which the caller has write access to at least one repository owned by the team.
	//  - admin: returns a list teams which the caller has team administrator access.
	Role string `url:"role,omitempty"`
}

// List returns all the teams that the authenticated user is associated with.
//
// Requires 'role' query parameter to be set.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams#get
func (t *TeamsService) List(opts ...interface{}) (*Teams, *Response, error) {
	teams := new(Teams)
	urlStr := t.client.requestUrl("/teams")
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
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
func (t *TeamsService) Get(teamUsername string, opts ...interface{}) (*Team, *Response, error) {
	team := new(Team)
	urlStr := t.client.requestUrl("/teams/%s", teamUsername)
	urlStr, addOptErr := addQueryParams(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, team, nil)

	return team, response, err
}
