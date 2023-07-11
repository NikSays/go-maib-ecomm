package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const executeRecurringCommand = "e"

// ExecuteRecurringPayload contains data required to execute a recurring transaction.
type ExecuteRecurringPayload struct {
	// Transaction payment amount. Positive integer with last 2 digits being the cents.
	//
	// Example: if Amount:199 and Currency:CurrencyUSD, $1.99 will be requested from the client's card.
	Amount uint `url:"amount"`

	// Transaction CurrencyEnum.
	// One of: CurrencyMDL, CurrencyEUR, CurrencyUSD.
	Currency CurrencyEnum `url:"currency"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIpAddress string `url:"client_ip_addr"`

	// Transaction details. Optional.
	Description string `url:"description,omitempty"`

	// Identifier of the recurring payment.
	BillerClientID string `url:"biller_client_id"`
}

// ExecuteRecurringResult contains data returned on execution of a recurring transaction,
// if no error is encountered.
type ExecuteRecurringResult struct {
	// ID of the executed transaction. 28 symbols in base64.
	TransactionId string `mapstructure:"TRANSACTION_ID"`

	// Transaction result status.
	Result ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// Retrieval reference number returned from Card Suite FO.
	RRN int `mapstructure:"RRN"`

	// Approval Code returned from Card Suite FO (max 6 digits).
	ApprovalCode int `mapstructure:"APPROVAL_CODE"`
}

// ExecuteRecurring executes a recurring transaction (-e) after it was created with
// [ECommClient.RegisterRecurring] (-z/-d/-p). It should not be finalized with
// [ECommClient.TransactionStatus] (-c) or [ECommClient.ExecuteDMS] (-t).
func (c *ECommClient) ExecuteRecurring(payload ExecuteRecurringPayload) (*ExecuteRecurringResult, error) {
	// Validate payload
	if !isValidAmount(payload.Amount) {
		return nil, errMalformedAmount
	}
	if !isValidCurrency(uint16(payload.Currency)) {
		return nil, errMalformedCurrency
	}
	if !isValidClientIpAddress(payload.ClientIpAddress) {
		return nil, errMalformedClientIP
	}
	if !isValidDescription(payload.Description) {
		return nil, errMalformedDescription
	}
	if !isValidBillerClientID(payload.BillerClientID) {
		return nil, errMalformedBillerClientID
	}
	// Send command
	payloadValues, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	res, err := c.send(executeRecurringCommand, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &ExecuteRecurringResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
