package maib

import (
	"fmt"
	"github.com/NikSays/go-maib-ecomm/requests"
	"github.com/NikSays/go-maib-ecomm/types"
)

func Example() {
	// Create new client
	client, err := NewClient(Config{
		PFXPath:                 "cert.pfx",
		Passphrase:              "p4ssphr4s3",
		MerchantHandlerEndpoint: "https://example.org/handler",
	})
	if err != nil {
		panic(err)
	}

	// Send a Transaction Registration request.
	// Equivalent to:
	// command=v&amount=1000&currency=978&language=en&client_ip_addr=127.0.0.1&description=10+EUR+will+be+charged
	res, err := client.Send(requests.RegisterTransaction{
		TransactionType: requests.RegisterTransactionSMS,
		Amount:          1000,
		Currency:        types.CurrencyEUR,
		ClientIPAddress: "127.0.0.1",
		Description:     "10 EUR will be charged",
		Language:        types.LanguageEnglish,
	})
	if err != nil {
		panic(err)
	}

	// Decode response into RegisterTransactionResult struct.
	newTransaction, err := requests.DecodeResponse[requests.RegisterTransactionResult](res)

	// Send a Transaction Status request.
	// Equivalent to:
	// command=c&trans_id=xxxxxxxxxxxxxxxxxxxxxxxxxxx&client_ip_addr=127.0.0.1
	res, err = client.Send(requests.TransactionStatus{
		TransactionID:   newTransaction.TransactionID,
		ClientIPAddress: "127.0.0.1",
	})
	if err != nil {
		panic(err)
	}

	// Decode response into TransactionStatusResult struct.
	status, err := requests.DecodeResponse[requests.TransactionStatusResult](res)
	if err != nil {
		panic(err)
	}

	fmt.Println(status.Result)
}
