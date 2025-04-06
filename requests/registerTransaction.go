package requests

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm"
	"github.com/NikSays/go-maib-ecomm/internal/validators"
)

// RegisterTransactionType holds possible types for recurring transaction.
type RegisterTransactionType int

const (
	// RegisterTransactionSMS is the Single Messaging System transaction type (-v).
	// Such a transaction is executed immediately and should be confirmed with
	// TransactionStatus (-c).
	//
	// This is the default transaction type.
	RegisterTransactionSMS RegisterTransactionType = iota

	// RegisterTransactionDMS is the Double Messaging System transaction type (-a).
	// This transaction should be confirmed with TransactionStatus (-c), and
	// executed with ExecuteDMS (-t).
	RegisterTransactionDMS
)

// String converts RegisterTransactionType into the ECommerce command. Returns
// an empty string for unknown values.
func (t RegisterTransactionType) String() string {
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
type RegisterTransaction struct {
	// Transaction type. Can be SMS (-v) or DMS (-a).
	// Default is SMS.
	TransactionType RegisterTransactionType `url:"-"`

	// Transaction payment amount. Positive integer with last 2 digits being the cents.
	//
	// Example: if Amount:199 and Currency:CurrencyUSD, $1.99 will be requested from
	// the client's card.
	Amount int `url:"amount"`

	// Transaction currency in ISO4217 3 digit format.
	Currency maib.Currency `url:"currency"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIPAddress string `url:"client_ip_addr"`

	// Transaction details. Optional.
	Description string `url:"description,omitempty"`

	// Language in which the bank payment page will be displayed.
	Language maib.Language `url:"language"`
}

// RegisterTransactionResult contains the response to a RegisterTransaction
// request.
type RegisterTransactionResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionID string `mapstructure:"TRANSACTION_ID"`
}

func (payload RegisterTransaction) Values() (url.Values, error) {
	err := validators.Validate(
		validators.WithTransactionType(payload.TransactionType.String()),
		validators.WithAmount(payload.Amount, true),
		validators.WithCurrency(payload.Currency),
		validators.WithClientIPAddress(payload.ClientIPAddress),
		validators.WithDescription(payload.Description),
		validators.WithLanguage(payload.Language),
	)
	if err != nil {
		return nil, fmt.Errorf("validate request: %w", err)
	}

	v, err := query.Values(payload)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}

	v.Set("command", payload.TransactionType.String())
	return v, nil
}
