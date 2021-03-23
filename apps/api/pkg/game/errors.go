package game

import "errors"

var (
	// ErrInvalidPlayerNumber - Invalid player number.
	ErrInvalidPlayerNumber = errors.New("invalid player number")
	// ErrInvalidPlayerID - Invalid player ID.
	ErrInvalidPlayerID = errors.New("invalid player id")
	// ErrInvalidTeamNumber - Invalid team number.
	ErrInvalidTeamNumber = errors.New("invalid team number")
)
