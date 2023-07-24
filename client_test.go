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
	newTr, err := requests.DecodeResponse[requests.RegisterTransactionResult](res)
	res, err = client.Send(requests.TransactionStatus{
		TransactionID:   newTr.TransactionID,
		ClientIPAddress: "127.0.0.1",
	})
	if err != nil {
		panic(err)
	}
	status, err := requests.DecodeResponse[requests.TransactionStatusResult](res)
	if err != nil {
		panic(err)
	}
	fmt.Println(status.Result)
}
