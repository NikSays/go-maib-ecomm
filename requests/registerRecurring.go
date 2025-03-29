package requests

import (
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm"
	"github.com/NikSays/go-maib-ecomm/internal/validators"
)

// RegisterRecurringType holds possible types for recurring transaction.
type RegisterRecurringType int

const (
	// RegisterRecurringSMS is a recurring transaction type which is initialized with an SMS transaction (-z).
	// The resulting transaction should be confirmed with TransactionStatus (-c).
	//
	// This is the default transaction type.
	RegisterRecurringSMS RegisterRecurringType = iota

	// RegisterRecurringDMS is a recurring transaction type which is initialized with a DMS transaction (-d).
	// The resulting transaction should be confirmed with TransactionStatus (-c), and executed with ExecuteDMS (-t).
	RegisterRecurringDMS

	// RegisterRecurringWithoutPayment is a recurring transaction type which is initialized without a transaction (-p).
	RegisterRecurringWithoutPayment
)

// String converts RegisterRecurringType into the ECommerce command.
// Returns an empty string for unknown values.
func (t RegisterRecurringType) String() string {
	switch t {
	case RegisterRecurringSMS:
		return "z"
	case RegisterRecurringDMS:
		return "d"
	case RegisterRecurringWithoutPayment:
		return "p"
	default:
		return ""
	}
}

// RegisterRecurring creates a new recurring transaction.
type RegisterRecurring struct {
	// Transaction type used for registration. Can be SMS (-z), DMS (-d), or without first payment (-p).
	// Default is SMS.
	TransactionType RegisterRecurringType `url:"-"`

	// Transaction payment amount. Positive integer with last 2 digits being the cents.
	// Ignored for registration without first payment.
	//
	// Example: if Amount:199 and Currency:CurrencyUSD, $1.99 will be requested from the client's card.
	Amount uint `url:"amount"`

	// Transaction currency in ISO4217 3 digit format.
	Currency maib.Currency `url:"currency"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIPAddress string `url:"client_ip_addr"`

	// Transaction details. Optional.
	Description string `url:"description,omitempty"`

	// Language in which the bank payment page will be displayed.
	Language maib.Language `url:"language"`

	// Identifier of the recurring payment. If not specified,
	// resulting TRANSACTION_ID will be used as the recurring payment ID.
	BillerClientID string `url:"biller_client_id"`

	// Validity limit of the regular payment in the format "MMYY".
	PerspayeeExpiry string `url:"perspayee_expiry"`

	// Whether the recurring transaction with a given BillerClientID should be updated.
	// This way, same BillerClientID may be used when customer changes payment information.
	OverwriteExisting bool `url:"-"`

	// If true, there will be a checkbox on the client handler. The card will be saved
	// only if the checkbox is checked. The field TransactionStatusResult.RecurringPaymentID
	// will be set only if the card is saved.
	AskSaveCardData bool `url:"-"`
}

// RegisterRecurringResult contains the response to a RegisterRecurring request.
type RegisterRecurringResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionID string `mapstructure:"TRANSACTION_ID"`
}

func (payload RegisterRecurring) Values() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}

	if payload.AskSaveCardData {
		v.Set("ask_save_card_data", "True")
	}

	if payload.OverwriteExisting {
		v.Set("perspayee_overwrite", "1")
	} else {
		v.Set("perspayee_gen", "1")
	}

	// Amount not needed for -p
	if payload.TransactionType == RegisterRecurringWithoutPayment {
		v.Del("amount")
	}
	v.Set("command", payload.TransactionType.String())
	return v, nil
}

func (payload RegisterRecurring) Validate() error {
	isAmountRequired := true
	if payload.TransactionType == RegisterRecurringWithoutPayment {
		isAmountRequired = false
	}
	return validators.Validate(
		validators.WithTransactionType(payload.TransactionType.String()),
		validators.WithAmount(payload.Amount, isAmountRequired),
		validators.WithCurrency(payload.Currency),
		validators.WithClientIPAddress(payload.ClientIPAddress),
		validators.WithDescription(payload.Description),
		validators.WithLanguage(payload.Language),
		validators.WithBillerClientID(payload.BillerClientID, false),
		validators.WithPerspayeeExpiry(payload.PerspayeeExpiry),
	)
}
