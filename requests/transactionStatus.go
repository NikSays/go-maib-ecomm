package requests

import (
	"github.com/NikSays/go-maib-ecomm/internal/validators"
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/google/go-querystring/query"
	"net/url"
)

const transactionStatusCommand = "c"

// TransactionStatus returns the status of a transaction (-c).
type TransactionStatus struct {
	// ID of the transaction. 28 symbols in base64.
	TransactionID string `url:"trans_id"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIPAddress string `url:"client_ip_addr"`
}

// TransactionStatusResult contains data returned by transaction status request (-c),
// if no error is encountered.
type TransactionStatusResult struct {
	// Transaction result status.
	Result types.ResultEnum `mapstructure:"RESULT"`

	// Transaction result, Payment Server interpretation.
	ResultPS types.ResultPSEnum `mapstructure:"RESULT_PS"`

	// Transaction resul code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// 3D Secure status.
	ThreeDSecure string `mapstructure:"3DSECURE"`

	ThreeDSecureReason string `mapstructure:"3DSECURE_REASON"`

	// Retrieval reference number returned from Card Suite FO.
	RRN int `mapstructure:"RRN"`

	// Approval Code returned from Card Suite FO (max 6 digits).
	ApprovalCode int `mapstructure:"APPROVAL_CODE"`

	// Masked card number.
	CardNumber string `mapstructure:"CARD_NUMBER"`

	AAV string `mapstructure:"AAV"`

	// PAR value identifying an account.
	PaymentAccountReference string `mapstructure:"PAYMENT_ACCOUNT_REFERENCE"`

	// Recurring payment identification in Payment Server.
	// Available only if transaction is recurring.
	RecurringPaymentID string `mapstructure:"RECC_PMNT_ID"`

	// Recurring payment expiry date in Payment Server in the form "MMYY".
	// Available only if transaction is recurring.
	RecurringPaymentExpiry string `mapstructure:"RECC_PMNT_EXPIRY"`
}

func (payload TransactionStatus) Encode() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	v.Set("command", transactionStatusCommand)
	return v, nil
}

func (payload TransactionStatus) Validate() error {
	return validators.Validate(
		validators.WithTransactionID(payload.TransactionID),
		validators.WithClientIPAddress(payload.ClientIPAddress),
	)
}
