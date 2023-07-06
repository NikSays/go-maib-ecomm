package maib

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func parseField(key string, value string) (any, error) {
	switch key {
	case
		"RESULT_CODE", "RRN", "APPROVAL_CODE",
		"fld_074", "fld_075", "fld_076", "fld_077",
		"fld_086", "fld_087", "fld_088", "fld_089":

		parsed, err := strconv.Atoi(value)
		return parsed, err
	}
	return value, nil
}

// send makes a POST request to the MAIB EComm servers.
func (c *ECommClient) send(command string, payload string) (map[string]any, error) {
	// Make request
	url := fmt.Sprintf("%s?command=%s&%s", c.merchantHandlerEndpoint, command, payload)
	res, err := c.httpClient.Post(url, "", nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't complete request to MAIB EComm: %w", err)
	}
	// Read body
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read response body: %w", err)
	}
	body := string(bodyBytes)

	// Catch error
	if strings.HasPrefix(body, "error") {
		return nil, fmt.Errorf("%w: %s", ErrMAIB, body)
	}

	// Parse response
	result := make(map[string]any, 0)
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("%w: wrong line format: \"%s\"", ErrParse, line)
		}
		key, value := parts[0], parts[1]
		parsedValue, err := parseField(key, value)
		if err != nil {
			return nil, fmt.Errorf("%w: wrong value type in \"%s\": %w", ErrParse, line, err)
		}
		result[key] = parsedValue
	}

	return result, nil
}
