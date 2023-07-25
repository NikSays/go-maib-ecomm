package types

import (
	"errors"
	"fmt"
)

// Errors encountered in response from MAIB EComm.
var (
	// ErrMAIB is returned when response from MAIB EComm starts with "error".
	ErrMAIB = errors.New("MAIB EComm returned an error")

	// ErrParse is returned when response from MAIB EComm doesn't follow "KEY: value" format,
	// or when a field is of an unexpected type.
	ErrParse = errors.New("couldn't parse response")
)

// ErrMalformedPayload is triggered before sending the request
// to MAIB EComm, if an error was encountered in payload input.
type ErrMalformedPayload struct {
	// Which field is malformed
	Field PayloadField

	// Human-readable explanation of the requirements
	Description string
}

func (e ErrMalformedPayload) Error() string {
	return fmt.Sprintf("malformed field %s (%s)", e.Field, e.Description)
}
