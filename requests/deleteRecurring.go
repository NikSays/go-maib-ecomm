package requests

import (
	"github.com/NikSays/go-maib-ecomm/internal/validators"
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/google/go-querystring/query"
	"net/url"
)

const deleteRecurringCommand = "x"

// DeleteRecurring deletes a recurring transaction (-x).
type DeleteRecurring struct {
	// Identifier of the recurring payment.
	BillerClientID string `url:"biller_client_id"`
}

// DeleteRecurringResult contains data returned on execution of a recurring transaction,
// if no error is encountered.
type DeleteRecurringResult struct {
	// Transaction result status.
	Result types.ResultEnum `mapstructure:"RESULT"`
}

func (payload DeleteRecurring) Encode() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	setCommand(&v, deleteRecurringCommand)
	return v, nil
}

func (payload DeleteRecurring) Validate() error {
	return validators.Validate(
		validators.WithBillerClientID(payload.BillerClientID, true),
	)
}
