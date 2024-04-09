package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm/types"
)

func TestTransactionStatus(t *testing.T) {
	cases := []struct {
		name               string
		payload            TransactionStatus
		expectedErrorField types.PayloadField
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
			expectedErrorField: types.FieldTransactionID,
		},
		{
			name: "ClientIPAddress invalid",
			payload: TransactionStatus{
				TransactionID:   "abcdefghijklmnopqrstuvwxyz1=",
				ClientIPAddress: "927.0.0.1",
			},
			expectedErrorField: types.FieldClientIPAddress,
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
				assert.Equal(t, c.expectedErrorField, err.(types.ValidationError).Field)
			}
		})
	}
}
