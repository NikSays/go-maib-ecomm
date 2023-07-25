package requests

import (
	"github.com/NikSays/go-maib-ecomm/internal/validators"
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/google/go-querystring/query"
	"net/url"
)

const reverseTransactionCommand = "r"

// ReverseTransaction reverses transaction and returns all or some funds to the client (-r).
type ReverseTransaction struct {
	// ID of the transaction. 28 symbols in base64.
	TransactionID string `url:"trans_id"`

	// Reversal amount. Positive integer with last 2 digits being the cents.
	//
	// For DMS authorizations only full amount can be reversed, i.e., the reversal and authorization
	// amounts have to match. In other cases, a partial reversal is also available.
	Amount uint `url:"amount"`

	// A flag indicating that a transaction is being reversed because of suspected fraud.
	// If this parameter is used, only full reversals are allowed.
	SuspectedFraud bool `url:"-"`
}

// ReverseTransactionResult contains data returned on reversal of a transaction,
// if no error is encountered.
type ReverseTransactionResult struct {
	// Transaction result status.
	Result types.ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`
}

func (payload ReverseTransaction) Encode() (url.Values, error) {
	v, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	if payload.SuspectedFraud {
		v.Set("suspected_fraud", "yes")
	}
	v.Set("command", reverseTransactionCommand)
	return v, nil
}

func (payload ReverseTransaction) Validate() error {
	return validators.Validate(
		validators.WithTransactionID(payload.TransactionID),
		validators.WithAmount(payload.Amount, true),
	)
}
