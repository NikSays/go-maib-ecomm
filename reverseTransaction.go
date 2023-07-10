package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const reverseCommand = "r"

// ReversePayload contains data required to reverse a transaction.
type ReversePayload struct {
	// ID of the transaction. 28 symbols in base64.
	TransactionId string `url:"trans_id"`

	// Reversal amount. Positive integer with last 2 digits being the cents.
	//
	// For DMS authorizations only full amount can be reversed, i.e., the reversal and authorization
	// amounts have to match. In other cases, a partial reversal is also available.
	Amount uint `url:"amount"`

	// A flag indicating that a transaction is being reversed because of suspected fraud.
	// If this parameter is used, only full reversals are allowed.
	SuspectedFraud bool `url:"suspected_fraud,omitempty"`
}

// ReverseResult contains data returned on reversal of a transaction,
// if no error is encountered.
type ReverseResult struct {
	// Transaction result status.
	Result resultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`
}

// ReverseTransaction reverses transaction and returns all or some funds to the client (-r).
func (c *ECommClient) ReverseTransaction(payload ReversePayload) (*ReverseResult, error) {
	// Validate payload
	if !isValidTransactionID(payload.TransactionId) {
		return nil, errMalformedTransactionID
	}
	if !isValidAmount(payload.Amount) {
		return nil, errMalformedAmount
	}
	payloadValues, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	// Parse true as 'yes'
	if payload.SuspectedFraud {
		payloadValues.Set("suspected_fraud", "yes")
	}
	// Send command
	res, err := c.send(reverseCommand, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &ReverseResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
