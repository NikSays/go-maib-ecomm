package requests

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm"
	"github.com/NikSays/go-maib-ecomm/internal/validators"
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
	Result maib.ResultEnum `mapstructure:"RESULT"`
}

func (payload DeleteRecurring) Values() (url.Values, error) {
	err := validators.Validate(
		validators.WithBillerClientID(payload.BillerClientID, true),
	)
	if err != nil {
		return nil, fmt.Errorf("validate request: %w", err)
	}

	v, err := query.Values(payload)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}

	v.Set("command", deleteRecurringCommand)
	return v, nil
}
