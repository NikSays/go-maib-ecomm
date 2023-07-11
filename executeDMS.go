package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const executeDmsCommand = "t"

// ExecuteDMSPayload contains data required to execute an DMS transaction.
type ExecuteDMSPayload struct {
	// ID of the transaction. 28 symbols in base64
	TransactionId string `url:"trans_id"`

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
}

// ExecuteDMSResult contains data returned on execution of an DMS transaction,
// if no error is encountered.
type ExecuteDMSResult struct {
	// Transaction result status.
	Result ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// Retrieval reference number returned from Card Suite FO.
	RRN int `mapstructure:"RRN"`

	// Approval Code returned from Card Suite FO (max 6 digits).
	ApprovalCode int `mapstructure:"APPROVAL_CODE"`

	// Masked card number.
	CardNumber string `mapstructure:"CARD_NUMBER"`
}

// ExecuteDMS executes a DMS transaction (-t) after it was created with [ECommClient.RegisterTransaction] (-a),
// and checked with [ECommClient.TransactionStatus] (-c).
func (c *ECommClient) ExecuteDMS(payload ExecuteDMSPayload) (*ExecuteDMSResult, error) {
	// Validate payload
	if !isValidTransactionID(payload.TransactionId) {
		return nil, errMalformedTransactionID
	}
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
	// Send command
	payloadValues, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	res, err := c.send(executeDmsCommand, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &ExecuteDMSResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
