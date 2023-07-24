package maib

import (
	"errors"
)

// Errors encountered in response from MAIB EComm.
var (
	// ErrMAIB is returned when response from MAIB EComm starts with "error".
	ErrMAIB = errors.New("MAIB EComm returned an error")

	// ErrParse is returned when response from MAIB EComm doesn't follow "KEY: value" format,
	// or when a field is of an unexpected type.
	ErrParse = errors.New("couldn't parse response")
)
