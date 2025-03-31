package requests

import (
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm"
	"github.com/NikSays/go-maib-ecomm/internal/validators"
)

const executeDMSCommand = "t"

// ExecuteDMS executes a DMS transaction (-t) after it was created with [RegisterTransaction] (-a),
// and checked with [TransactionStatus] (-c).
type ExecuteDMS struct {
	// ID of the transaction. 28 symbols in base64.
	TransactionID string `url:"trans_id"`

	// Transaction payment amount. Positive integer with last 2 digits being the cents.
	//
	// Example: if Amount:199 and Currency:CurrencyUSD, $1.99 will be requested from the client's card.
	Amount int `url:"amount"`

	// Transaction currency in ISO4217 3 digit format.
	Currency maib.Currency `url:"currency"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIPAddress string `url:"client_ip_addr"`

	// Transaction details. Optional.
	Description string `url:"description,omitempty"`
}

// ExecuteDMSResult contains the response to a ExecuteDMS request.
type ExecuteDMSResult struct {
	// Transaction result status.
	Result maib.ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// Retrieval reference number returned from Card Suite FO.
	RRN int `mapstructure:"RRN"`

	// Approval Code returned from Card Suite FO (max 6 characters).
	ApprovalCode string `mapstructure:"APPROVAL_CODE"`

	// Masked card number.
	CardNumber string `mapstructure:"CARD_NUMBER"`
}

func (payload ExecuteDMS) Values() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	v.Set("command", executeDMSCommand)
	return v, nil
}

func (payload ExecuteDMS) Validate() error {
	return validators.Validate(
		validators.WithTransactionID(payload.TransactionID),
		validators.WithAmount(payload.Amount, true),
		validators.WithCurrency(payload.Currency),
		validators.WithClientIPAddress(payload.ClientIPAddress),
		validators.WithDescription(payload.Description),
	)
}
