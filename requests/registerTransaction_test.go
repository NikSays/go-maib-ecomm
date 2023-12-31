package requests

import (
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRegisterTransaction(t *testing.T) {
	cases := []struct {
		name               string
		payload            RegisterTransaction
		expectedErrorField types.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK SMS",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        types.LanguageEnglish,
			},
			expectedEncoded: "amount=1234&client_ip_addr=127.0.0.1&command=v&currency=498&description=Description&language=en",
		},
		{
			name: "OK DMS",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionDMS,
				Amount:          1234,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        types.LanguageEnglish,
			},
			expectedEncoded: "amount=1234&client_ip_addr=127.0.0.1&command=a&currency=498&description=Description&language=en",
		},
		{
			name: "TransactionType invalid",
			payload: RegisterTransaction{
				TransactionType: -9,
				Amount:          1234,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        types.LanguageEnglish,
			},
			expectedErrorField: types.FieldCommand,
		},
		{
			name: "Amount invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          0,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        types.LanguageEnglish,
			},
			expectedErrorField: types.FieldAmount,
		},
		{
			name: "Currency invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        1000,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        types.LanguageEnglish,
			},
			expectedErrorField: types.FieldCurrency,
		},
		{
			name: "ClientIPAddress invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "927.0.0.1",
				Description:     "Description",
				Language:        types.LanguageEnglish,
			},
			expectedErrorField: types.FieldClientIPAddress,
		},
		{
			name: "Description invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     strings.Repeat("-", 130),
				Language:        types.LanguageEnglish,
			},
			expectedErrorField: types.FieldDescription,
		},
		{
			name: "Description encoding",
			payload: RegisterTransaction{
				Amount:          1234,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "this=should&not=be&injected",
				Language:        types.LanguageEnglish,
			},
			expectedEncoded: "amount=1234&client_ip_addr=127.0.0.1&command=v&currency=498&description=this%3Dshould%26not%3Dbe%26injected&language=en",
		},
		{
			name: "Language invalid",
			payload: RegisterTransaction{
				TransactionType: RegisterTransactionSMS,
				Amount:          1234,
				Currency:        types.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        "",
			},
			expectedErrorField: types.FieldLanguage,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.payload.Validate()
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
				enc, err := c.payload.Encode()
				assert.Nil(t, err)
				assert.Equal(t, c.expectedEncoded, enc.Encode())
			} else {
				assert.Equal(t, c.expectedErrorField, err.(types.ErrMalformedPayload).Field)
			}
		})
	}
}
