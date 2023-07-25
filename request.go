package maib

import (
	"fmt"
	"github.com/NikSays/go-maib-ecomm/types"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// Request is a payload that can be sent to MAIB EComm server.
type Request interface {
	// Encode returns the payload as a URL value map, that can be encoded into necessary querystring
	// to be sent to the EComm server.
	Encode() (url.Values, error)

	// Validate goes through the fields of the payload, and returns an error if any one of them
	// does not fit the requirements.
	Validate() error
}

// Send validates a [Request], and sends it to MAIB EComm servers.
// The value returned on success can be parsed into a result struct using requests.DecodeResponse
func (c ECommClient) Send(req Request) (map[string]any, error) {
	queryValues, err := req.Encode()
	if err != nil {
		return nil, err
	}
	// Validate request
	err = req.Validate()
	if err != nil {
		return nil, err
	}

	// Make request
	urlString := fmt.Sprintf("%s?%s", c.merchantHandlerEndpoint, queryValues.Encode())
	res, err := c.httpClient.Post(urlString, "", nil)
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

	// Parse response
	result := make(map[string]any)
	lines := strings.Split(body, "\n")
	for _, line := range lines {
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
		"RESULT_CODE", "RRN", "APPROVAL_CODE",
		"fld_074", "fld_075", "fld_076", "fld_077",
		"fld_086", "fld_087", "fld_088", "fld_089":

		parsed, err := strconv.Atoi(value)
		return parsed, err
	}
	return value, nil
}
