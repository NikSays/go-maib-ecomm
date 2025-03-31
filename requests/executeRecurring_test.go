package requests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm"
)

func TestExecuteRecurring(t *testing.T) {
	cases := []struct {
		name               string
		payload            ExecuteRecurring
		expectedErrorField maib.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK",
			payload: ExecuteRecurring{
				BillerClientID:  "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedEncoded: "amount=1234&biller_client_id=abcdefghijklmnopqrstuvwxyz1%3D&client_ip_addr=127.0.0.1&command=e&currency=498&description=Description",
		},
		{
			name: "BillerClientID invalid",
			payload: ExecuteRecurring{
				BillerClientID:  "",
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldBillerClientID,
		},
		{
			name: "Amount invalid",
			payload: ExecuteRecurring{
				BillerClientID:  "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          0,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldAmount,
		},
		{
			name: "Currency invalid",
			payload: ExecuteRecurring{
				BillerClientID:  "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          1234,
				Currency:        1000,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldCurrency,
		},
		{
			name: "ClientIPAddress invalid",
			payload: ExecuteRecurring{
				BillerClientID:  "abcdefghijklmnopqrstuvwxyz1=",
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "927.0.0.1",
				Description:     "Description",
			},
			expectedErrorField: maib.FieldClientIPAddress,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			val, err := c.payload.Values()
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
				assert.Equal(t, c.expectedEncoded, val.Encode())
			} else {
				valErr := &maib.ValidationError{}
				isValErr := errors.As(err, &valErr)
				assert.True(t, isValErr)
				assert.Equal(t, c.expectedErrorField, valErr.Field)
			}
		})
	}
}
