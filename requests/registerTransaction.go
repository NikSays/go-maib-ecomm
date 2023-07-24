package requests

import (
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/google/go-querystring/query"
	"net/url"
)

type registerTransactionTypesEnum string

// Possible types for transaction
const (
	RegisterTransactionSMS registerTransactionTypesEnum = "v"
	RegisterTransactionDMS registerTransactionTypesEnum = "a"
)

// RegisterTransaction creates a new SMS (-v) or DMS (-a) transaction.
//
// For SMS transactions:
// The resulting transaction should be confirmed with [ECommClient.TransactionStatus] (-c).
//
// For DMS transactions:
// The resulting transaction should be confirmed with [ECommClient.TransactionStatus] (-c),
// and executed with [ECommClient.ExecuteDMS] (-t).
type RegisterTransaction struct {
	// Transaction type. Can be SMS (-v) or DMS (-a)
	TransactionType registerTransactionTypesEnum `url:"-"`

	// Transaction payment amount. Positive integer with last 2 digits being the cents.
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
}

// RegisterTransactionResult contains data returned on execution of a transaction registration request,
// if no error is encountered.
type RegisterTransactionResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionID string `mapstructure:"TRANSACTION_ID"`
}

func (payload RegisterTransaction) Encode() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	if len(payload.TransactionType) == 0 {
		payload.TransactionType = RegisterTransactionSMS
	}
	setCommand(&v, payload.TransactionType)
	return v, nil
}
