package requests

import (
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/google/go-querystring/query"
	"net/url"
)

type registerRecurringTypesEnum string

// Possible types for recurring transaction
const (
	RegisterRecurringSMS            registerRecurringTypesEnum = "z"
	RegisterRecurringDMS            registerRecurringTypesEnum = "d"
	RegisterRecurringWithoutPayment registerRecurringTypesEnum = "p"
)

// RegisterRecurring creates a new recurring transaction.
//
// For SMS transactions:
// The resulting transaction should be confirmed with [TransactionStatus] (-c).
//
// For DMS transactions:
// The resulting transaction should be confirmed with [TransactionStatus] (-c),
// and executed with [ExecuteDMS] (-t).
type RegisterRecurring struct {
	// Transaction type used for registration. Can be SMS (-z), DMS (-d), or without first payment (-p).
	TransactionType registerRecurringTypesEnum `url:"-"`

	// Transaction payment amount. Positive integer with last 2 digits being the cents.
	// Ignored for registration without first payment.
	//
	// Example: if Amount:199 and Currency:CurrencyUSD, $1.99 will be requested from the client's card.
	Amount uint `url:"amount"`

	// Transaction currency in ISO4217 3 digit format.
	Currency types.Currency `url:"currency"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIPAddress string `url:"client_ip_addr"`

	// Transaction details. Optional.
	Description string `url:"description,omitempty"`

	// Language in which the bank payment page will be displayed
	Language types.Language `url:"language"`

	// Identifier of the recurring payment. If not specified,
	// resulting TRANSACTION_ID will be used as the recurring payment ID.
	BillerClientID string `url:"biller_client_id"`

	// Validity limit of the regular payment in the format "MMYY".
	PerspayeeExpiry string `url:"perspayee_expiry"`

	// Whether the recurring transaction with a given BillerClientID should be updated.
	// This way, same BillerClientID may be used when customer changes payment information.
	OverwriteExisting bool `url:"-"`
}

// RegisterRecurringResult contains data returned after registration of a recurring transaction,
// if no error is encountered.
type RegisterRecurringResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionID string `mapstructure:"TRANSACTION_ID"`
}

func (payload RegisterRecurring) Encode() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	if payload.OverwriteExisting {
		v.Set("perspayee_overwrite", "1")
	} else {
		v.Set("perspayee_gen", "1")
	}
	// Default transaction type
	if len(payload.TransactionType) == 0 {
		payload.TransactionType = RegisterRecurringSMS
	}
	// Amount not needed for -p
	if payload.TransactionType == RegisterRecurringWithoutPayment {
		v.Del("amount")
	}
	setCommand(&v, payload.TransactionType)
	return v, nil
}
