package requests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm"
)

func TestExecuteDMS(t *testing.T) {
	cases := []struct {
		name               string
		payload            ExecuteDMS
		expectedErrorField maib.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK",
			payload: ExecuteDMS{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedEncoded: "amount=1234&client_ip_addr=127.0.0.1&command=t&currency=498&description=Description&trans_id=abcdefghijklmnopqrstuvwxyz1%3D",
		},
		{
			name: "TransactionID invalid",
			payload: ExecuteDMS{
				TransactionID:   "",
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldTransactionID,
		},
		{
			name: "Amount invalid",
			payload: ExecuteDMS{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          0,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldAmount,
		},
		{
			name: "Currency invalid",
			payload: ExecuteDMS{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          1234,
				Currency:        1000,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldCurrency,
		},
		{
			name: "ClientIPAddress invalid",
			payload: ExecuteDMS{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "927.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldClientIPAddress,
		},
		{
			name: "Description invalid",
			payload: ExecuteDMS{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     strings.Repeat("-", 130),
			},
			expectedErrorField: maib.FieldDescription,
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
