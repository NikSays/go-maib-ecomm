package maib

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Send validates a [Request], and sends it to the ECommerce system.
// The value returned on success can be parsed into a result struct using requests.DecodeResponse
func (c *Client) Send(req Request) (map[string]any, error) {
	// Validate request
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate request: %w", err)
	}

	// Send request
	reqURL, err := url.Parse(c.merchantHandlerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}

	queryValues, err := req.Values()
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}
	reqURL.RawQuery = queryValues.Encode()
	res, err := c.httpClient.Post(reqURL.String(), "", nil)
	if err != nil {
		return nil, fmt.Errorf("send request to MAIB EComm: %w", err)
	}

	// Read body
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	body := string(bodyBytes)

	// Catch error
	if res.StatusCode != http.StatusOK || strings.HasPrefix(body, "error") {
		return nil, &ECommError{
			Code: res.StatusCode,
			Body: body,
		}
	}

	result, err := parseBody(body)
	if err != nil {
		return nil, &ParseError{
			Err:  err,
			Body: body,
		}
	}

	return result, nil
}

// ECommError is returned when the ECommerce system responds with
// a non-200 status, or when the response body starts with "error:".
type ECommError struct {
	// HTTP status code
	Code int

	// Response body
	Body string
}

func (e *ECommError) Error() string {
	return fmt.Sprintf("maib ecomm returned %d: %s", e.Code, e.Body)
}
