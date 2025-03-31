package maib

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBody_OK(t *testing.T) {
	const (
		textValue = "TEXT"
		intValue  = "123"
	)

	// Read an example MAIB EComm response to parse into map
	body, err := os.ReadFile("testdata/response.txt")
	assert.Nil(t, err)

	// Parse using values to determine type
	original := make(map[string]any)
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ": ")
		key, value := parts[0], parts[1]
		switch value {
		case textValue:
			original[key] = value
		case intValue:
			original[key], err = strconv.Atoi(value)
			assert.Nil(t, err)
		default:
			t.Error("Unknown value in response.txt")
		}
	}

	// Parse using keys to determine type
	parsed, err := parseBody(string(body))
	assert.Nil(t, err)

	// Verify that all fields have the correct value
	for k := range original {
		assert.Equal(t, original[k], parsed[k])
	}
}

func TestParseBody_MalformedLine(t *testing.T) {
	body := "No colon"
	parsed, err := parseBody(body)

	assert.Nil(t, parsed)
	assert.NotNil(t, err)
}

func TestParseBody_InvalidType(t *testing.T) {
	body := "RESULT_CODE: TEXT"
	parsed, err := parseBody(body)

	assert.Nil(t, parsed)
	assert.NotNil(t, err)
}
