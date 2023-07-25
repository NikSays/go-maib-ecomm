package requests

import (
	"github.com/mitchellh/mapstructure"
)

type resultTypes interface {
	CloseDayResult | DeleteRecurringResult | ExecuteDMSResult |
		ExecuteRecurringResult | RegisterRecurringResult | RegisterTransactionResult |
		ReverseTransactionResult | TransactionStatusResult
}

// DecodeResponse is a generic function that parses the map returned from the MAIB EComm server
// into any Result type. The type must be specified in [brackets]
//
// Example:
//
//	DecodeResponse[CloseDayResult](someResponse)
func DecodeResponse[ResultType resultTypes](maibResponse map[string]any) (result ResultType, err error) {
	err = mapstructure.Decode(maibResponse, &result)
	return result, err
}