package types

import (
	"fmt"
)

// ParseError is returned when the response from the ECommerce system
// doesn't follow "KEY: value" format, or when a field has an unexpected type .
type ParseError struct {
	// Underlying error
	Err error

	// Response body that couldn't be parsed
	Body string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("error parsing response: %s", e.Err)
}

// Unwrap returns the underlying error, for usage with errors.As.
func (e *ParseError) Unwrap() error {
	return e.Err
}

// ECommError is returned when the ECommerce system responds with
// a non-200 status, or when the response body starts with "error:".
type ECommError struct {
	// HTTP status code
	Code int

	// Response body
	Body string
}

func (e *ECommError) Error() string {
	return fmt.Sprintf("maib ecomm returned %d: %s", e.Code, e.Body)
}

// ValidationError is triggered before sending the request to the
// ECommerce system, if the request has failed validation.
type ValidationError struct {
	// Which field is malformed.
	Field PayloadField

	// Human-readable explanation of the requirements.
	Description string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("malformed field %s: %s", e.Field, e.Description)
}
