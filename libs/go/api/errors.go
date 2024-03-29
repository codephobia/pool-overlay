package api

import "errors"

var (
	// ErrEndpointMethodNotAllowed - Method supplied was not allowed.
	ErrEndpointMethodNotAllowed = errors.New("method not allowed")
	// ErrEndpointForbidden - API endpoint forbidden.
	ErrEndpointForbidden = errors.New("forbidden")
	// ErrInvalidPageNumber - Invalid page number.
	ErrInvalidPageNumber = errors.New("invalid page number")
	// ErrInternalServerError - Internal server error.
	ErrInternalServerError = errors.New("internal server error")
	// ErrInvalidPlayerID - Invalid player id.
	ErrInvalidPlayerID = errors.New("invalid player id")
	// ErrPlayerNotFound - Player not found.
	ErrPlayerNotFound = errors.New("player not found")
	// ErrInvalidPlayerDetails - Invalid player details.
	ErrInvalidPlayerDetails = errors.New("invalid player details")
	// ErrInvalidTeamID - Invalid team id.
	ErrInvalidTeamID = errors.New("invalid team id")
	// ErrTeamNotFound - Team not found.
	ErrTeamNotFound = errors.New("team not found")
	// ErrInvalidGameType - Game type invalid.
	ErrInvalidGameType = errors.New("invalid game type")
	// ErrInvalidGameVsMode - Game mode invalid.
	ErrInvalidGameVsMode = errors.New("invalid game mode")
	// ErrInvalidGameDirection - Game direction invalid.
	ErrInvalidGameDirection = errors.New("invalid game direction")
	// ErrUnableToBroadcastUpdate - Unable to broadcast update.
	ErrUnableToBroadcastUpdate = errors.New("unable to broadcast update")
	// ErrTournamentNotFound - Tournament not found.
	ErrTournamentNotFound = errors.New("tournament not found")
	// ErrInvalidTournamentID - Invalid tournament id.
	ErrInvalidTournamentID = errors.New("invalid tournament id")
	// ErrInvalidTournamentDetails - Invalid tournament details.
	ErrInvalidTournamentDetails = errors.New("invalid tournament details")
	// ErrInvalidTableNumber - Invalid table number.
	ErrInvalidTableNumber = errors.New("invalid table number")
	// ErrRemoveOnlyTable - Cannot remove only table.
	ErrRemoveOnlyTable = errors.New("cannot remove only table")
)
