// Package validators provides functions to validate input without boilerplate.
package validators

import (
	"encoding/base64"
	"net"
	"strconv"

	"github.com/NikSays/go-maib-ecomm/v2"
)

// FieldValidator is the function used as argument to [Validate].
type FieldValidator func() error

// Validate runs the argument functions until one of them returns an error. Use
// any of the `With...` functions as arguments.
func Validate(validators ...FieldValidator) error {
	for _, v := range validators {
		err := v()
		if err != nil {
			return err
		}
	}
	return nil
}

// WithTransactionType verifies that transactionType is exactly one character,
// and not the default empty value.
func WithTransactionType(transactionType string) FieldValidator {
	return func() error {
		if len(transactionType) != 1 {
			return &maib.ValidationError{
				Field:       maib.FieldCommand,
				Description: "not 1 character",
			}
		}
		return nil
	}
}

// WithTransactionID verifies that transactionID is 28 base64 characters.
func WithTransactionID(transactionID string) FieldValidator {
	return func() error {
		if len(transactionID) != 28 {
			return &maib.ValidationError{
				Field:       maib.FieldTransactionID,
				Description: "not 28 characters",
			}
		}
		if _, err := base64.StdEncoding.DecodeString(transactionID); err != nil {
			return &maib.ValidationError{
				Field:       maib.FieldTransactionID,
				Description: "not in base64",
			}
		}
		return nil
	}

}

// WithAmount verifies that amount is not negative; at most 12 digits;
// not 0, if required.
func WithAmount(amount int, required bool) FieldValidator {
	return func() error {
		if amount < 0 {
			return &maib.ValidationError{
				Field:       maib.FieldAmount,
				Description: "negative number",
			}
		} else if amount > 999999999999 {
			return &maib.ValidationError{
				Field:       maib.FieldAmount,
				Description: "more than 12 digits",
			}
		} else if required && amount == 0 {
			return &maib.ValidationError{
				Field:       maib.FieldAmount,
				Description: "not a positive number",
			}
		}
		return nil
	}
}

// WithCurrency verifies that currency is a 3 digit non-negative integer.
func WithCurrency(currency maib.Currency) FieldValidator {
	return func() error {
		if currency < 0 || currency > 999 {
			return &maib.ValidationError{
				Field:       maib.FieldCurrency,
				Description: "invalid ISO 4217 3-number code",
			}
		}
		return nil
	}
}

// WithClientIPAddress verifies that address is a valid IP address.
func WithClientIPAddress(address string) FieldValidator {
	return func() error {
		ip := net.ParseIP(address)
		if ip == nil {
			return &maib.ValidationError{
				Field:       maib.FieldClientIPAddress,
				Description: "invalid IP address",
			}
		}
		return nil
	}
}

// WithLanguage verifies that language is a non-empty string, with at most
// 32 characters.
func WithLanguage(language maib.Language) FieldValidator {
	return func() error {
		if len(language) < 1 || len(language) > 32 {
			return &maib.ValidationError{
				Field:       maib.FieldLanguage,
				Description: "not between 1 and 32 characters",
			}
		}
		return nil
	}
}

// WithBillerClientID verifies that billerClientID is at most 49 characters;
// not empty, if required.
func WithBillerClientID(billerClientID string, required bool) FieldValidator {
	return func() error {
		if len(billerClientID) > 49 {
			return &maib.ValidationError{
				Field:       maib.FieldBillerClientID,
				Description: "more than 49 characters",
			}
		} else if required && len(billerClientID) < 1 {
			return &maib.ValidationError{
				Field:       maib.FieldBillerClientID,
				Description: "empty string",
			}
		}
		return nil
	}
}

// WithPerspayeeExpiry verifies that prespayeeExpiry is 4 characters, first 2
// being a number between 1 and 12, second 2 being a non-negative integer.
func WithPerspayeeExpiry(prespayeeExpiry string) FieldValidator {
	return func() error {
		if len(prespayeeExpiry) != 4 {
			return &maib.ValidationError{
				Field:       maib.FieldPerspayeeExpiry,
				Description: "not 4 digits",
			}
		}
		month, err := strconv.Atoi(prespayeeExpiry[0:2])
		if err != nil {
			return &maib.ValidationError{
				Field:       maib.FieldPerspayeeExpiry,
				Description: "not a valid month",
			}
		}
		if month < 1 || month > 12 {
			return &maib.ValidationError{
				Field:       maib.FieldPerspayeeExpiry,
				Description: "not a valid month",
			}
		}
		year, err := strconv.Atoi(prespayeeExpiry[2:4])
		if err != nil {
			return &maib.ValidationError{
				Field:       maib.FieldPerspayeeExpiry,
				Description: "not a valid year",
			}
		}
		if year < 0 {
			return &maib.ValidationError{
				Field:       maib.FieldPerspayeeExpiry,
				Description: "not a valid year",
			}
		}
		return nil
	}
}

// WithDescription verifies that description is at most 125 characters.
func WithDescription(description string) FieldValidator {
	return func() error {
		if len(description) > 125 {
			return &maib.ValidationError{
				Field:       maib.FieldDescription,
				Description: "more than 125 characters",
			}
		}
		return nil
	}
}
