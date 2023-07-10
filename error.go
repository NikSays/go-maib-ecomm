package maib

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

type ErrMalformedPayload struct {
	Field       string
	Description string
}

func (e ErrMalformedPayload) Error() string {
	return fmt.Sprintf("malformed field %s (%s)", e.Field, e.Description)
}

// Errors encountered in payload input.
var (
	errMalformedTransactionID = ErrMalformedPayload{
		Field:       "TransactionID",
		Description: "not 28 characters in base64",
	}
	errMalformedAmount = ErrMalformedPayload{
		Field:       "Amount",
		Description: "either 0 or more than 12 digits",
	}
	errMalformedCurrency = ErrMalformedPayload{
		Field:       "Currency",
		Description: "invalid ISO 4217 3-number code",
	}
	errMalformedClientIP = ErrMalformedPayload{
		Field:       "ClientIPAddress",
		Description: "invalid IP address",
	}
	errMalformedDescription = ErrMalformedPayload{
		Field:       "Description",
		Description: "more than 125 characters",
	}
	errMalformedLanguage = ErrMalformedPayload{
		Field:       "Language",
		Description: "either 0 or more than 32 characters",
	}
)
