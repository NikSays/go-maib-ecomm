package requests

import (
	"github.com/NikSays/go-maib-ecomm/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteRecurring(t *testing.T) {
	cases := []struct {
		name               string
		payload            DeleteRecurring
		expectedErrorField types.PayloadField
		expectedEncoded    string
	}{
		{
			name:            "OK",
			payload:         DeleteRecurring{BillerClientID: "abcABC123="},
			expectedEncoded: "biller_client_id=abcABC123%3D&command=x",
		},
		{
			name:               "BillerClientID invalid",
			payload:            DeleteRecurring{BillerClientID: ""},
			expectedErrorField: types.FieldBillerClientID,
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
