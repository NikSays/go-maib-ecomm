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

type malformedInputEnum string

// malformedInputEnum contains the names of the parameters and payload fields that can
// fail validation, and thus will be returned in [ErrMalformedPayload]
const (
	FieldTransactionId   malformedInputEnum = "trans_id"
	FieldAmount          malformedInputEnum = "amount"
	FieldCurrency        malformedInputEnum = "currency"
	FieldClientIpAddress malformedInputEnum = "client_ip_addr"
	FieldDescription     malformedInputEnum = "description"
	FieldLanguage        malformedInputEnum = "language"
	FieldBillerClientId  malformedInputEnum = "biller_client_id"
	FieldPerspayeeExpiry malformedInputEnum = "prespayee_expiry"
	FieldTransactionType malformedInputEnum = "transaction_type"
)

// ErrMalformedPayload is triggered before sending the request
// to MAIB EComm, if an error was encountered in payload input.
type ErrMalformedPayload struct {
	// Which field is malformed
	Field malformedInputEnum

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
	errMalformedBillerClientID = ErrMalformedPayload{
		Field:       FieldBillerClientId,
		Description: "either 0 or more than 49 characters",
	}
	errMalformedPerspayeeExpiry = ErrMalformedPayload{
		Field:       FieldPerspayeeExpiry,
		Description: "not 4 digits in MMYY format",
	}
	errMalformedTransactionType = ErrMalformedPayload{
		Field:       FieldTransactionType,
		Description: "not SMS or DMS",
	}
)
