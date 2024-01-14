package requests

import (
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm/internal/validators"
	"github.com/NikSays/go-maib-ecomm/types"
)

const executeRecurringCommand = "e"

// ExecuteRecurring executes a recurring transaction (-e) after it was created with
// [RegisterRecurring] (-z/-d/-p). It should not be finalized with
// [TransactionStatus] (-c) or [ExecuteDMS] (-t).
type ExecuteRecurring struct {
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

	// Identifier of the recurring payment.
	BillerClientID string `url:"biller_client_id"`
}

// ExecuteRecurringResult contains data returned on execution of a recurring transaction,
// if no error is encountered.
type ExecuteRecurringResult struct {
	// ID of the executed transaction. 28 symbols in base64.
	TransactionID string `mapstructure:"TRANSACTION_ID"`

	// Transaction result status.
	Result types.ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// Retrieval reference number returned from Card Suite FO.
	RRN int `mapstructure:"RRN"`

	// Approval Code returned from Card Suite FO (max 6 characters).
	ApprovalCode string `mapstructure:"APPROVAL_CODE"`
}

func (payload ExecuteRecurring) Values() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	v.Set("command", executeRecurringCommand)
	return v, nil
}

func (payload ExecuteRecurring) Validate() error {
	return validators.Validate(
		validators.WithAmount(payload.Amount, true),
		validators.WithCurrency(payload.Currency),
		validators.WithClientIPAddress(payload.ClientIPAddress),
		validators.WithDescription(payload.Description),
		validators.WithBillerClientID(payload.BillerClientID, true),
	)
}
