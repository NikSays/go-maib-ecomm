package requests

import (
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm/internal/validators"
	"github.com/NikSays/go-maib-ecomm/types"
)

const deleteRecurringCommand = "x"

// DeleteRecurring deletes a recurring transaction (-x).
type DeleteRecurring struct {
	// Identifier of the recurring payment.
	BillerClientID string `url:"biller_client_id"`
}

// DeleteRecurringResult contains the response to a DeleteRecurring request.
type DeleteRecurringResult struct {
	// Transaction result status.
	Result types.ResultEnum `mapstructure:"RESULT"`
}

func (payload DeleteRecurring) Values() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	v.Set("command", deleteRecurringCommand)
	return v, nil
}

func (payload DeleteRecurring) Validate() error {
	return validators.Validate(
		validators.WithBillerClientID(payload.BillerClientID, true),
	)
}
