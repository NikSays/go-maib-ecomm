package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const statusCommand = "c"

// StatusPayload contains data required to fetch transaction status (-c).
type StatusPayload struct {
	// ID of the transaction. 28 symbols in base64.
	TransactionId string `url:"trans_id"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIpAddress string `url:"client_ip_addr"`
}

// StatusResult contains data returned by transaction status request (-c),
// if no error is encountered.
type StatusResult struct {
	// Transaction result status.
	Result ResultEnum `mapstructure:"RESULT"`

	// Transaction result, Payment Server interpretation.
	ResultPs ResultPSEnum `mapstructure:"RESULT_PS"`

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

	// PAR value identifying an account
	PaymentAccountReference string `mapstructure:"PAYMENT_ACCOUNT_REFERENCE"`

	// Recurring payment identification in Payment Server.
	// Available only if transaction is recurring.
	RecurringPaymentId string `mapstructure:"RECC_PMNT_ID"`

	// Recurring payment expiry date in Payment Server in the form "MMYY".
	// Available only if transaction is recurring.
	RecurringPaymentExpiry string `mapstructure:"RECC_PMNT_EXPIRY"`
}

// TransactionStatus returns the status of a transaction (-c).
func (c *ECommClient) TransactionStatus(payload StatusPayload) (*StatusResult, error) {
	// Validate payload
	if !isValidTransactionID(payload.TransactionId) {
		return nil, errMalformedTransactionID
	}
	if !isValidClientIpAddress(payload.ClientIpAddress) {
		return nil, errMalformedClientIP
	}
	// Send command
	payloadValues, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	res, err := c.send(statusCommand, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &StatusResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
