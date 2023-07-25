package requests

import (
	"github.com/NikSays/go-maib-ecomm/internal/validators"
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/google/go-querystring/query"
	"net/url"
)

type registerTransactionTypeEnum int

// Possible types for transaction
const (
	RegisterTransactionSMS registerTransactionTypeEnum = iota // default
	RegisterTransactionDMS
)

func (t registerTransactionTypeEnum) String() string {
	switch t {
	case RegisterTransactionSMS:
		return "v"
	case RegisterTransactionDMS:
		return "a"
	default:
		return ""
	}
}

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
	TransactionType registerTransactionTypeEnum `url:"-"`

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
	v.Set("command", payload.TransactionType.String())
	return v, nil
}

func (payload RegisterTransaction) Validate() error {
	return validators.Validate(
		validators.WithTransactionType(payload.TransactionType.String()),
		validators.WithAmount(payload.Amount, true),
		validators.WithCurrency(payload.Currency),
		validators.WithClientIPAddress(payload.ClientIPAddress),
		validators.WithDescription(payload.Description),
		validators.WithLanguage(payload.Language),
	)
}
