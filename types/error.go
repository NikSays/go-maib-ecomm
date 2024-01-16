package types

import (
	"fmt"
)

// ErrParse is returned when the response from the ECommerce system
// doesn't follow "KEY: value" format, or when a field has an unexpected type.
type ErrParse struct {
	// Underlying error
	Reason error
}

func (e ErrParse) Error() string {
	return fmt.Sprintf("error parsing response: %s", e.Reason)
}

func (e ErrParse) Unwrap() error {
	return e.Reason
}

// ErrMAIB is returned when the ECommerce system responds with
// a non-200 status, or when the response body starts with "error:".
type ErrMAIB struct {
	// HTTP status code
	Code int

	// Response body
	Body string
}

func (e ErrMAIB) Error() string {
	return fmt.Sprintf("maib ecomm returned %d: %s", e.Code, e.Body)
}

// ErrMalformedPayload is triggered before sending the request to the
// ECommerce system, if the request has failed validation.
type ErrMalformedPayload struct {
	// Which field is malformed.
	Field PayloadField

	// Human-readable explanation of the requirements.
	Description string
}

func (e ErrMalformedPayload) Error() string {
	return fmt.Sprintf("malformed field %s: %s", e.Field, e.Description)
}
