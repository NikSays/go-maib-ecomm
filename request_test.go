package maib

import (
	"fmt"
	"github.com/NikSays/go-maib-ecomm/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBody(t *testing.T) {
	const (
		textValue = "TEXT"
		intValue  = 123
	)
	textFields := []string{
		"RESULT", "RESULT_PS", "3DSECURE",
		"CARD_NUMBER", "TRANSACTION_ID",
		"RECC_PMNT_ID", "RECC_PMNT_EXPIRY",
	}

	intFields := []string{
		"RESULT_CODE", "RRN",
		"FLD_074", "FLD_075", "FLD_076", "FLD_077",
		"FLD_086", "FLD_087", "FLD_088", "FLD_089",
	}

	var body string
	for _, f := range textFields {
		body += fmt.Sprintf("%s: %s\n", f, textValue)
	}
	for _, f := range intFields {
		body += fmt.Sprintf("%s: %d\n", f, intValue)
	}

	parsed, err := parseBody(body)

	assert.Nil(t, err)
	for _, f := range textFields {
		assert.Equal(t, textValue, parsed[f])
	}
	for _, f := range intFields {
		assert.Equal(t, intValue, parsed[f])
	}

	_, err = requests.DecodeResponse[requests.CloseDayResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.DeleteRecurringResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.ExecuteDMSResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.ExecuteRecurringResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.RegisterRecurringResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.RegisterTransactionResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.ReverseTransactionResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.TransactionStatusResult](parsed)
	assert.Nil(t, err)
}
