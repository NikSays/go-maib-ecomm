package requests

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleDecodeResponse() {
	// Example of a response from the ECommerce system for a RegisterTransaction request
	ecommResponse := map[string]any{
		"TRANSACTION_ID": "abcdefghijklmnopqrstuvwxyz1=",
	}

	result, err := DecodeResponse[RegisterTransactionResult](ecommResponse)
	if err != nil {
		panic(err)
	}

	fmt.Print(result.TransactionID)

	// Output: abcdefghijklmnopqrstuvwxyz1=
}

func TestDecodeResponse(t *testing.T) {
	const (
		textValue = "TEXT"
		intValue  = "123"
	)

	// Read an example MAIB EComm response to parse into map
	body, err := os.ReadFile("../testdata/response.txt")
	assert.Nil(t, err)

	// Parse using values as guide
	parsed := make(map[string]any)
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ": ")
		key, value := parts[0], parts[1]
		switch value {
		case textValue:
			parsed[key] = value
		case intValue:
			parsed[key], err = strconv.Atoi(value)
			assert.Nil(t, err)
		default:
			t.Error("Unknown value in response.txt")
		}
	}

	// Try to decode into all results
	// Prevents datatype inconsistencies
	_, err = DecodeResponse[CloseDayResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[DeleteRecurringResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[ExecuteDMSResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[ExecuteRecurringResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[RegisterRecurringResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[ExecuteOneClickResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[RegisterOneClickResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[RegisterTransactionResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[ReverseTransactionResult](parsed)
	assert.Nil(t, err)
	_, err = DecodeResponse[TransactionStatusResult](parsed)
	assert.Nil(t, err)
}
