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
)
