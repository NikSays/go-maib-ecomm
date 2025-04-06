package maib

import (
	"fmt"
	"strconv"
	"strings"
)

// parseBody splits each line as "key: value", converting types.
func parseBody(body string) (map[string]any, error) {
	result := make(map[string]any)
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("wrong line format: \"%s\"", line)
		}

		key, value := parts[0], parts[1]
		parsedValue, err := parseField(key, value)
		if err != nil {
			return nil, fmt.Errorf("wrong value type in \"%s\": %w", line, err)
		}

		result[key] = parsedValue
	}
	return result, nil
}

// parseField returns int or string value depending on the key.
func parseField(key string, value string) (any, error) {
	switch key {
	// Possible int fields in response
	case
		"RESULT_CODE", "RRN",
		"FLD_074", "FLD_075", "FLD_076", "FLD_077",
		"FLD_086", "FLD_087", "FLD_088", "FLD_089":

		parsed, err := strconv.Atoi(value)
		return parsed, err
	}
	return value, nil
}

// ParseError is returned when the response from the ECommerce system doesn't
// follow "KEY: value" format, or when a field has an unexpected type.
type ParseError struct {
	// Underlying error.
	Err error

	// Response body that couldn't be parsed.
	Body string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse response: %s", e.Err)
}

// Unwrap returns the underlying error, for usage with [errors.As].
func (e *ParseError) Unwrap() error {
	return e.Err
}
