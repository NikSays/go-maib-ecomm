package maib

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Request is a payload that can be sent to the ECommerce system.
type Request interface {
	// Values validates the request and returns the payload as a URL value map, that
	// can be encoded into a querystring to be sent to the ECommerce system.
	Values() (url.Values, error)
}

// Send validates a [Request] and sends it to the ECommerce system. The value
// returned on success can be parsed into a result struct using
// requests.DecodeResponse.
//
// The request is cancelled when the context is done.
func (c *Client) Send(ctx context.Context, req Request) (map[string]any, error) {
	reqURL, err := url.Parse(c.merchantHandlerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}

	queryValues, err := req.Values()
	if err != nil {
		return nil, fmt.Errorf("get request values: %w", err)
	}
	reqURL.RawQuery = queryValues.Encode()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	res, err := c.httpClient.Do(httpReq)
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

// ECommError is returned when the ECommerce system responds with a non-200
// status, or when the response body starts with "error:".
type ECommError struct {
	// HTTP status code.
	Code int

	// Response body.
	Body string
}

func (e *ECommError) Error() string {
	return fmt.Sprintf("maib ecomm returned %d: %s", e.Code, e.Body)
}
