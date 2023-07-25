package validators

import (
	"encoding/base64"
	"github.com/NikSays/go-maib-ecomm/types"
	"net"
	"strconv"
)

type FieldValidator func() error

func Validate(validators ...FieldValidator) error {
	for _, v := range validators {
		err := v()
		if err != nil {
			return err
		}
	}
	return nil
}

func WithTransactionType(transactionType string) FieldValidator {
	return func() error {
		if len(transactionType) != 1 {
			return types.ErrMalformedPayload{
				Field:       types.FieldTransactionType,
				Description: "not 1 character",
			}
		}
		return nil
	}
}

func WithTransactionID(transactionID string) FieldValidator {
	return func() error {
		if len(transactionID) != 28 {
			return types.ErrMalformedPayload{
				Field:       types.FieldTransactionID,
				Description: "not 28 characters",
			}
		}
		if _, err := base64.StdEncoding.DecodeString(transactionID); err != nil {
			return types.ErrMalformedPayload{
				Field:       types.FieldTransactionID,
				Description: "not in base64",
			}
		}
		return nil
	}

}

func WithAmount(amount uint, required bool) FieldValidator {
	return func() error {
		if amount > 999999999999 {
			return types.ErrMalformedPayload{
				Field:       types.FieldAmount,
				Description: "more than 12 digits",
			}
		} else if required && amount <= 0 {
			return types.ErrMalformedPayload{
				Field:       types.FieldAmount,
				Description: "not a positive number",
			}
		}
		return nil
	}
}

func WithCurrency(currency types.Currency) FieldValidator {
	return func() error {
		if currency < 0 || currency > 999 {
			return types.ErrMalformedPayload{
				Field:       types.FieldCurrency,
				Description: "invalid ISO 4217 3-number code",
			}
		}
		return nil
	}
}

func WithClientIPAddress(address string) FieldValidator {
	return func() error {
		ip := net.ParseIP(address)
		if ip == nil {
			return types.ErrMalformedPayload{
				Field:       types.FieldClientIPAddress,
				Description: "invalid IP address",
			}
		}
		return nil
	}
}

func WithLanguage(language types.Language) FieldValidator {
	return func() error {
		if len(language) < 1 || len(language) > 32 {
			return types.ErrMalformedPayload{
				Field:       types.FieldLanguage,
				Description: "not between 1 and 32 characters",
			}
		}
		return nil
	}
}

func WithBillerClientID(billerClientID string, required bool) FieldValidator {
	return func() error {
		if len(billerClientID) > 49 {
			return types.ErrMalformedPayload{
				Field:       types.FieldBillerClientID,
				Description: "more than 49 characters",
			}
		} else if required && len(billerClientID) < 1 {
			return types.ErrMalformedPayload{
				Field:       types.FieldBillerClientID,
				Description: "empty string",
			}
		}
		return nil
	}
}

func WithPerspayeeExpiry(prespayeeExpiry string) FieldValidator {
	return func() error {
		if len(prespayeeExpiry) != 4 {
			return types.ErrMalformedPayload{
				Field:       types.FieldPerspayeeExpiry,
				Description: "not 4 digits",
			}
		}
		month, err := strconv.Atoi(prespayeeExpiry[0:2])
		if err != nil {
			return types.ErrMalformedPayload{
				Field:       types.FieldPerspayeeExpiry,
				Description: "not a valid month",
			}
		}
		if month < 1 || month > 12 {
			return types.ErrMalformedPayload{
				Field:       types.FieldPerspayeeExpiry,
				Description: "not a valid month",
			}
		}
		year, err := strconv.Atoi(prespayeeExpiry[2:4])
		if err != nil {
			return types.ErrMalformedPayload{
				Field:       types.FieldPerspayeeExpiry,
				Description: "not a valid year",
			}
		}
		if year < 0 {
			return types.ErrMalformedPayload{
				Field:       types.FieldPerspayeeExpiry,
				Description: "not a valid year",
			}
		}
		return nil
	}
}

func WithDescription(description string) FieldValidator {
	return func() error {
		if len(description) > 125 {
			return types.ErrMalformedPayload{
				Field:       types.FieldDescription,
				Description: "more than 125 characters",
			}
		}
		return nil
	}
}
