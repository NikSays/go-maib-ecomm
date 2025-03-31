package requests

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm"
	"github.com/NikSays/go-maib-ecomm/internal/validators"
)

// RegisterOneClickType holds possible types for recurring transaction.
type RegisterOneClickType int

const (
	// RegisterOneClickSMS is a oneClick transaction type which is initialized with an SMS transaction (-z).
	// The resulting transaction should be confirmed with TransactionStatus (-c).
	//
	// This is the default transaction type.
	RegisterOneClickSMS RegisterOneClickType = iota

	// RegisterOneClickWithoutPayment is a oneClick transaction type which is initialized without a transaction (-p).
	RegisterOneClickWithoutPayment
)

// String converts RegisterOneClickType into the ECommerce command.
// Returns an empty string for unknown values.
func (t RegisterOneClickType) String() string {
	switch t {
	case RegisterOneClickSMS:
		return "z"
	case RegisterOneClickWithoutPayment:
		return "p"
	default:
		return ""
	}
}

// RegisterOneClick creates a new oneClick transaction. This allows the merchant
// to save a card, like with recurring transactions. However, oneClick transactions
// cannot be instantly executed. The client must be redirected to the client handler
// to confirm the transaction.
type RegisterOneClick struct {
	// Transaction type used for registration. Can be SMS (-z), or without first payment (-p).
	// Default is SMS.
	TransactionType RegisterOneClickType `url:"-"`

	// Transaction payment amount. Positive integer with last 2 digits being the cents.
	// Ignored for registration without first payment.
	//
	// Example: if Amount:199 and Currency:CurrencyUSD, $1.99 will be requested from the client's card.
	Amount int `url:"amount"`

	// Transaction currency in ISO4217 3 digit format.
	Currency maib.Currency `url:"currency"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIPAddress string `url:"client_ip_addr"`

	// Transaction details. Optional.
	Description string `url:"description,omitempty"`

	// Language in which the bank payment page will be displayed.
	Language maib.Language `url:"language"`

	// Identifier of the oneClick payment. If not specified,
	// resulting TRANSACTION_ID will be used as the oneClick payment ID.
	BillerClientID string `url:"biller_client_id"`

	// Validity limit of the regular payment in the format "MMYY".
	PerspayeeExpiry string `url:"perspayee_expiry"`

	// Whether the oneClick transaction with a given BillerClientID should be updated.
	// This way, same BillerClientID may be used when customer changes payment information.
	OverwriteExisting bool `url:"-"`

	// If true, there will be a checkbox on the client handler. The card will be saved
	// only if the checkbox is checked. The field TransactionStatusResult.RecurringPaymentID
	// will be set only if the card is saved.
	AskSaveCardData bool `url:"-"`
}

// RegisterOneClickResult contains the response to a RegisterOneClick request.
type RegisterOneClickResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionID string `mapstructure:"TRANSACTION_ID"`
}

func (payload RegisterOneClick) Values() (url.Values, error) {
	isAmountRequired := true
	if payload.TransactionType == RegisterOneClickWithoutPayment {
		isAmountRequired = false
	}
	err := validators.Validate(
		validators.WithTransactionType(payload.TransactionType.String()),
		validators.WithAmount(payload.Amount, isAmountRequired),
		validators.WithCurrency(payload.Currency),
		validators.WithClientIPAddress(payload.ClientIPAddress),
		validators.WithDescription(payload.Description),
		validators.WithLanguage(payload.Language),
		validators.WithBillerClientID(payload.BillerClientID, false),
		validators.WithPerspayeeExpiry(payload.PerspayeeExpiry),
	)
	if err != nil {
		return nil, fmt.Errorf("validate request: %w", err)
	}

	v, err := query.Values(payload)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}

	v.Set("oneclick", "Y")

	if payload.AskSaveCardData {
		v.Set("ask_save_card_data", "True")
	}

	if payload.OverwriteExisting {
		v.Set("perspayee_overwrite", "1")
	} else {
		v.Set("perspayee_gen", "1")
	}

	// Amount not needed for -p
	if payload.TransactionType == RegisterOneClickWithoutPayment {
		v.Del("amount")
	}
	v.Set("command", payload.TransactionType.String())
	return v, nil
}
