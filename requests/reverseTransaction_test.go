package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm/types"
)

func TestReverseTransaction(t *testing.T) {
	cases := []struct {
		name               string
		payload            ReverseTransaction
		expectedErrorField types.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK",
			payload: ReverseTransaction{
				TransactionID:  "abcdefghijklmnopqrstuvwxyz1=",
				Amount:         1234,
				SuspectedFraud: false,
			},
			expectedEncoded: "amount=1234&command=r&trans_id=abcdefghijklmnopqrstuvwxyz1%3D",
		},
		{
			name: "OK fraud",
			payload: ReverseTransaction{
				TransactionID:  "abcdefghijklmnopqrstuvwxyz1=",
				Amount:         1234,
				SuspectedFraud: true,
			},
			expectedEncoded: "amount=1234&command=r&suspected_fraud=yes&trans_id=abcdefghijklmnopqrstuvwxyz1%3D",
		},
		{
			name: "TransactionID invalid",
			payload: ReverseTransaction{
				TransactionID:  "",
				Amount:         1234,
				SuspectedFraud: false,
			},
			expectedErrorField: types.FieldTransactionID,
		},
		{
			name: "Amount invalid",
			payload: ReverseTransaction{
				TransactionID:  "abcdefghijklmnopqrstuvwxyz1=",
				Amount:         0,
				SuspectedFraud: false,
			},
			expectedErrorField: types.FieldAmount,
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
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
			}
		})
	}
}
