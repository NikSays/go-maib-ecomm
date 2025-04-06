package requests

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm"
	"github.com/NikSays/go-maib-ecomm/internal/validators"
)

const executeOneClickCommand = "f"

// ExecuteOneClick executes a oneClick transaction (-f) after it was created
// with [RegisterOneClick] (-z/-p with oneclick=Y). It should be finalized with
// [TransactionStatus] (-c).
type ExecuteOneClick struct {
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

	// Identifier of the oneClick payment.
	BillerClientID string `url:"biller_client_id"`
}

// ExecuteOneClickResult contains the response to a ExecuteOneClick request.
type ExecuteOneClickResult struct {
	// ID of the executed transaction. 28 symbols in base64.
	TransactionID string `mapstructure:"TRANSACTION_ID"`

	// Transaction result status.
	Result maib.ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// Retrieval reference number returned from Card Suite FO.
	RRN int `mapstructure:"RRN"`

	// Approval Code returned from Card Suite FO (max 6 characters).
	ApprovalCode string `mapstructure:"APPROVAL_CODE"`
}

func (payload ExecuteOneClick) Values() (url.Values, error) {
	err := validators.Validate(
		validators.WithAmount(payload.Amount, true),
		validators.WithCurrency(payload.Currency),
		validators.WithClientIPAddress(payload.ClientIPAddress),
		validators.WithDescription(payload.Description),
		validators.WithBillerClientID(payload.BillerClientID, true),
	)
	if err != nil {
		return nil, fmt.Errorf("validate request: %w", err)
	}

	v, err := query.Values(payload)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}

	v.Set("oneclick", "Y")
	v.Set("command", executeOneClickCommand)
	return v, nil
}
