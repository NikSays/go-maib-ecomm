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

type malformedFieldEnum string

// malformedFieldEnum contains the names of the fields that can
// fail validation, and thus will be returned in [ErrMalformedPayload]
const (
	FieldTransactionId   malformedFieldEnum = "trans_id"
	FieldAmount          malformedFieldEnum = "amount"
	FieldCurrency        malformedFieldEnum = "currency"
	FieldClientIpAddress malformedFieldEnum = "client_ip_addr"
	FieldDescription     malformedFieldEnum = "description"
	FieldLanguage        malformedFieldEnum = "language"
)

// ErrMalformedPayload is triggered before sending the request
// to MAIB EComm, if an error was encountered in payload input.
type ErrMalformedPayload struct {
	// Which field is malformed
	Field malformedFieldEnum

	// Human-readable explanation of the requirements
	Description string
}

func (e ErrMalformedPayload) Error() string {
	return fmt.Sprintf("malformed field %s (%s)", e.Field, e.Description)
}

// For internal use, since the same descriptions are repeated
var (
	errMalformedTransactionID = ErrMalformedPayload{
		Field:       FieldTransactionId,
		Description: "not 28 characters in base64",
	}
	errMalformedAmount = ErrMalformedPayload{
		Field:       FieldAmount,
		Description: "either 0 or more than 12 digits",
	}
	errMalformedCurrency = ErrMalformedPayload{
		Field:       FieldCurrency,
		Description: "invalid ISO 4217 3-number code",
	}
	errMalformedClientIP = ErrMalformedPayload{
		Field:       FieldClientIpAddress,
		Description: "invalid IP address",
	}
	errMalformedDescription = ErrMalformedPayload{
		Field:       FieldDescription,
		Description: "more than 125 characters",
	}
	errMalformedLanguage = ErrMalformedPayload{
		Field:       FieldLanguage,
		Description: "either 0 or more than 32 characters",
	}
)
