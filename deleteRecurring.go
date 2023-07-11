package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const deleteRecurringCommand = "x"

// DeleteRecurringPayload contains data required to execute a recurring transaction.
type DeleteRecurringPayload struct {
	// Identifier of the recurring payment.
	BillerClientID string `url:"biller_client_id"`
}

// DeleteRecurringResult contains data returned on execution of a recurring transaction,
// if no error is encountered.
type DeleteRecurringResult struct {
	// Transaction result status.
	Result ResultEnum `mapstructure:"RESULT"`
}

// DeleteRecurring deletes a recurring transaction (-x).
func (c *ECommClient) DeleteRecurring(payload DeleteRecurringPayload) (*DeleteRecurringResult, error) {
	// Validate payload
	if !isValidBillerClientID(payload.BillerClientID) {
		return nil, errMalformedBillerClientID
	}
	// Send command
	payloadValues, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	res, err := c.send(deleteRecurringCommand, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &DeleteRecurringResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
