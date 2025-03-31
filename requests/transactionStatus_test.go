package requests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm"
)

func TestTransactionStatus(t *testing.T) {
	cases := []struct {
		name               string
		payload            TransactionStatus
		expectedErrorField maib.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK",
			payload: TransactionStatus{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				ClientIPAddress: "127.0.0.1",
			},
			expectedEncoded: "client_ip_addr=127.0.0.1&command=c&trans_id=abcdefghijklmnopqrstuvwxyz1%3D",
		},

		{
			name: "TransactionID invalid",
			payload: TransactionStatus{
				TransactionID:   "",
				ClientIPAddress: "127.0.0.1",
			},
			expectedErrorField: maib.FieldTransactionID,
		},
		{
			name: "ClientIPAddress invalid",
			payload: TransactionStatus{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				ClientIPAddress: "927.0.0.1",
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
