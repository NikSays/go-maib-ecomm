package requests

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/NikSays/go-maib-ecomm/v2"
	"github.com/NikSays/go-maib-ecomm/v2/internal/validators"
)

const reverseTransactionCommand = "r"

// ReverseTransaction reverses transaction and returns all or some funds to the
// client (-r).
type ReverseTransaction struct {
	// ID of the transaction. 28 symbols in base64.
	TransactionID string `url:"trans_id"`

	// Reversal amount. Positive integer with last 2 digits being the cents.
	//
	// For DMS authorizations only full amount can be reversed, i.e., the reversal
	// and authorization amounts have to match. In other cases, a partial reversal
	// is also available.
	Amount int `url:"amount"`

	// A flag indicating that a transaction is being reversed because of suspected
	// fraud. If this parameter is used, only full reversals are allowed.
	SuspectedFraud bool `url:"-"`
}

// ReverseTransactionResult contains the response to a ReverseTransaction
// request.
type ReverseTransactionResult struct {
	// Transaction result status.
	Result maib.ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`
}

func (payload ReverseTransaction) Values() (url.Values, error) {
	err := validators.Validate(
		validators.WithTransactionID(payload.TransactionID),
		validators.WithAmount(payload.Amount, true),
	)
	if err != nil {
		return nil, fmt.Errorf("validate request: %w", err)
	}

	v, err := query.Values(payload)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}

	if payload.SuspectedFraud {
		v.Set("suspected_fraud", "yes")
	}
	v.Set("command", reverseTransactionCommand)
	return v, nil
}
