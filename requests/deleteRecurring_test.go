package requests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm/v2"
)

func TestDeleteRecurring(t *testing.T) {
	cases := []struct {
		name               string
		payload            DeleteRecurring
		expectedErrorField maib.PayloadField
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
			expectedErrorField: maib.FieldBillerClientID,
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
