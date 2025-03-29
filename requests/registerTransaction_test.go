package requests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm"
)

func TestRegisterTransaction(t *testing.T) {
	cases := []struct {
		name               string
		payload            RegisterTransaction
		expectedErrorField maib.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK SMS",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        maib.LanguageEnglish,
			},
			expectedEncoded: "amount=1234&client_ip_addr=127.0.0.1&command=v&currency=498&description=Description&language=en",
		},
		{
			name: "OK DMS",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionDMS,
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        maib.LanguageEnglish,
			},
			expectedEncoded: "amount=1234&client_ip_addr=127.0.0.1&command=a&currency=498&description=Description&language=en",
		},
		{
			name: "TransactionType invalid",
			payload: RegisterTransaction{
				TransactionType: -9,
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        maib.LanguageEnglish,
			},
			expectedErrorField: maib.FieldCommand,
		},
		{
			name: "Amount invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          0,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        maib.LanguageEnglish,
			},
			expectedErrorField: maib.FieldAmount,
		},
		{
			name: "Currency invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        1000,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        maib.LanguageEnglish,
			},
			expectedErrorField: maib.FieldCurrency,
		},
		{
			name: "ClientIPAddress invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "927.0.0.1",
				Description:     "Description",
				Language:        maib.LanguageEnglish,
			},
			expectedErrorField: maib.FieldClientIPAddress,
		},
		{
			name: "Description invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     strings.Repeat("-", 130),
				Language:        maib.LanguageEnglish,
			},
			expectedErrorField: maib.FieldDescription,
		},
		{
			name: "Description encoding",
			payload: RegisterTransaction{
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "this=should&not=be&injected",
				Language:        maib.LanguageEnglish,
			},
			expectedEncoded: "amount=1234&client_ip_addr=127.0.0.1&command=v&currency=498&description=this%3Dshould%26not%3Dbe%26injected&language=en",
		},
		{
			name: "Language invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        "",
			},
			expectedErrorField: maib.FieldLanguage,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.payload.Validate()
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
				val, err := c.payload.Values()
				assert.Nil(t, err)
				assert.Equal(t, c.expectedEncoded, val.Encode())
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}
