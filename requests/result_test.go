package requests

import "fmt"

func ExampleDecodeResponse() {
	// Example of a response from the ECommerce system for a RegisterTransaction request
	ecommResponse := map[string]any{
		"TRANSACTION_ID": "abcdefghijklmnopqrstuvwxyz1=",
	}

	result, err := DecodeResponse[RegisterTransactionResult](ecommResponse)
	if err != nil {
		panic(err)
	}

	fmt.Print(result.TransactionID)

	// Output: abcdefghijklmnopqrstuvwxyz1=
}
