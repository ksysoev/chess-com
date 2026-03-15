package chesscom

import (
	"errors"
	"fmt"
)

// Sentinel errors returned by the client for specific HTTP status codes.
var (
	// ErrNotFound is returned when the API responds with HTTP 404.
	ErrNotFound = errors.New("resource not found")

	// ErrGone is returned when the API responds with HTTP 410, indicating
	// that the resource is permanently unavailable and should not be requested again.
	ErrGone = errors.New("resource permanently unavailable")

	// ErrRateLimited is returned when the API responds with HTTP 429.
	// The caller should back off before retrying.
	ErrRateLimited = errors.New("rate limited by server")
)

// APIError represents an unexpected HTTP error response from the Chess.com API.
type APIError struct {
	Status     string
	StatusCode int
}

// Error implements the error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("chess.com api error: %s", e.Status)
}

// newAPIError creates a new APIError.
func newAPIError(statusCode int, status string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Status:     status,
	}
}
