package apidocs

// swagger:route GET /game/update/teams game GameUpdateTeams
// Sets a team number to a team ID.
// responses:
//   200: gameResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdateTeams
type GameUpdateTeamsParam struct {
	// The team number to set a team on.
	//
	// in: query
	// required: true
	// min: 1
	// max: 2
	// example: 1
	TeamNum int `json:"teamNum"`

	// The ID of the team to set on the teamNum.
	//
	// in: query
	// required: true
	// min: 1
	// example: 1
	TeamID int `json:"teamID"`
}
