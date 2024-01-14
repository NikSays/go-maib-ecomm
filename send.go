package maib

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/NikSays/go-maib-ecomm/types"
)

// Send validates a [Request], and sends it to the MAIB EComm system.
// The value returned on success can be parsed into a result struct using requests.DecodeResponse
func (c *client) Send(req Request) (map[string]any, error) {
	// Validate request
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// Make request
	reqURL, err := url.Parse(c.merchantHandlerEndpoint)
	if err != nil {
		return nil, err
	}

	queryValues, err := req.Values()
	if err != nil {
		return nil, err
	}
	reqURL.RawQuery = queryValues.Encode()
	res, err := c.httpClient.Post(reqURL.String(), "", nil)
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
		return nil, fmt.Errorf("%w: %s", types.ErrMAIB, body)
	}

	result, err := parseBody(body)

	return result, err
}

// parseBody splits each line as "key: value", converting types
func parseBody(body string) (map[string]any, error) {
	result := make(map[string]any)
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("%w: wrong line format: \"%s\"", types.ErrParse, line)
		}

		key, value := parts[0], parts[1]
		parsedValue, err := parseField(key, value)
		if err != nil {
			return nil, fmt.Errorf("%w: wrong value type in \"%s\": %w", types.ErrParse, line, err)
		}

		result[key] = parsedValue
	}
	return result, nil
}

// parseField returns int or string value depending on the key
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
