package maib

import (
	"context"
	"errors"
	"fmt"
	"time"

	// Since the real requests package imports the current package, a fake requests
	// package must be used. Otherwise, an import cycle is created.
	requests "github.com/NikSays/go-maib-ecomm/v2/testdata/fakerequests"
)

func Example() {
	// In this example we will
	// * Create a Client
	// * Execute an SMS transaction and decode the response
	// * Check the created transaction's status
	// Errors are ignored for brevity, please handle them in your code.

	// Create new client to send requests to MAIB ECommerce
	client, _ := NewClient(Config{
		PFXPath:                 "cert.pfx",
		Passphrase:              "p4ssphr4s3",
		MerchantHandlerEndpoint: "https://example.org/handler",
	})

	// Execute an SMS transaction (-v) for 10 Euro.
	// Equivalent to this POST request:
	// command=v&amount=1000&currency=978&language=en&client_ip_addr=127.0.0.1&description=10+EUR+will+be+charged
	res, _ := client.Send(requests.RegisterTransaction{
		TransactionType: requests.RegisterTransactionSMS,
		Amount:          1000,
		Currency:        CurrencyEUR,
		ClientIPAddress: "127.0.0.1",
		Description:     "10 EUR will be charged",
		Language:        LanguageEnglish,
	})

	// Decode response map into RegisterTransactionResult struct,
	// to get the ID of the created transaction.
	newTransaction, _ := requests.DecodeResponse[requests.RegisterTransactionResult](res)

	// Send a Transaction Status request with a timeout.
	// Equivalent to this POST request:
	// command=c&trans_id=<TransactionID>&client_ip_addr=127.0.0.1
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	res, _ = client.SendWithContext(ctx, requests.TransactionStatus{
		TransactionID:   newTransaction.TransactionID,
		ClientIPAddress: "127.0.0.1",
	})

	// Decode response map into TransactionStatusResult struct,
	// to get the transaction result.
	status, _ := requests.DecodeResponse[requests.TransactionStatusResult](res)

	// Print the result of the transaction
	fmt.Println(status.Result)
}

func Example_errorHandling() {
	// Send a request with a client.
	client, _ := NewClient(Config{ /* ... */ })
	_, err := client.Send(requests.RegisterTransaction{ /* ... */ })

	// The target of errors.As should be a pointer to a type that implements error.
	// Since the Error method is defined on *ValidationError, the second argument to
	// errors.As should be **ValidationError.
	//
	// The same principle is used for ECommError and ParseError.
	if valErr := (&ValidationError{}); errors.As(err, &valErr) {
		fmt.Printf("Invalid field %s: %s", valErr.Field, valErr.Description)
	}
}
