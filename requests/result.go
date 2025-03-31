package requests

import (
	"github.com/mitchellh/mapstructure"
)

type resultTypes interface {
	CloseDayResult | DeleteRecurringResult | ExecuteDMSResult | ExecuteOneClickResult |
		ExecuteRecurringResult | RegisterRecurringResult | RegisterOneClickResult |
		RegisterTransactionResult | ReverseTransactionResult | TransactionStatusResult
}

// DecodeResponse is a generic function that parses the map returned from the ECommerce system
// into any Result type. The generic type must be specified.
func DecodeResponse[ResultType resultTypes](maibResponse map[string]any) (result ResultType, err error) {
	err = mapstructure.Decode(maibResponse, &result)
	return
}
