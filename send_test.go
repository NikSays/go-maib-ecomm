package maib

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm/requests"
	"github.com/NikSays/go-maib-ecomm/types"
)

const (
	caPath = "testdata/certs/ca.crt"

	serverCertPath = "testdata/certs/server.crt"
	serverKeyPath  = "testdata/certs/server.key"
)

func loadCerts() (caPool *x509.CertPool, serverCert tls.Certificate, err error) {
	// Read CA certificate
	caCert, err := os.ReadFile(caPath)
	if err != nil {
		return nil, tls.Certificate{}, fmt.Errorf("read CA: %w", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Read server's HTTPS certificate
	serverCert, err = tls.LoadX509KeyPair(serverCertPath, serverKeyPath)
	if err != nil {
		return nil, tls.Certificate{}, fmt.Errorf("load server certificate: %w", err)
	}

	return caCertPool, serverCert, nil
}

func createServer(caPool *x509.CertPool, serverCert tls.Certificate, handler http.HandlerFunc) *httptest.Server {
	tlsConfig := &tls.Config{
		// HTTPS certificate
		Certificates: []tls.Certificate{serverCert},
		// mTLS Client verification
		ClientCAs:  caPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	// Create a mTLS Server instance
	server := httptest.NewUnstartedServer(handler)
	server.TLS = tlsConfig
	return server
}

func createTrustingClient(endpointURL string, caPool *x509.CertPool) (*Client, error) {
	// Create client
	client, err := NewClient(Config{
		PFXPath:                 clientCertPath,
		Passphrase:              clientCertPass,
		MerchantHandlerEndpoint: endpointURL,
	})
	if err != nil {
		return nil, err
	}

	// Trust the local CA
	client.httpClient.Transport.(*http.Transport).TLSClientConfig.RootCAs = caPool

	return client, nil
}

func TestClient_Send_InvalidRequest(t *testing.T) {
	client := Client{}
	res, err := client.Send(requests.ExecuteDMS{
		ClientIPAddress: "invalid",
	})

	assert.Nil(t, res)
	assert.ErrorAs(t, err, &types.ValidationError{})
}

func TestClient_Send_InvalidEndpoint(t *testing.T) {
	client := Client{
		merchantHandlerEndpoint: ":",
	}
	res, err := client.Send(requests.CloseDay{})

	var urlErr *url.Error
	assert.Nil(t, res)
	assert.ErrorAs(t, err, &urlErr)
}
func TestClient_Send_WithCerts(t *testing.T) {
	caPool, serverCert, err := loadCerts()
	assert.Nil(t, err)

	t.Run("OK", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "b", request.FormValue("command"))
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		_, err = client.Send(requests.CloseDay{})
		assert.Nil(t, err)
	})

	t.Run("TLS fail", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, &x509.CertPool{})
		assert.Nil(t, err)

		_, err = client.Send(requests.CloseDay{})
		var certErr *tls.CertificateVerificationError
		assert.ErrorAs(t, err, &certErr)
	})

	t.Run("Bad status", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusInternalServerError)
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		_, err = client.Send(requests.CloseDay{})
		assert.ErrorAs(t, err, &types.ECommError{})
	})

	t.Run("Error response", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("error: ecommerce has encountered an unknown error"))
			assert.Nil(t, err)
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		_, err = client.Send(requests.CloseDay{})
		assert.ErrorAs(t, err, &types.ECommError{})
	})

	t.Run("Malformed response", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("welcome"))
			assert.Nil(t, err)
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		_, err = client.Send(requests.CloseDay{})
		assert.ErrorAs(t, err, &types.ParseError{})
	})
}

func TestParseBody_OK(t *testing.T) {
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

	// Build a MAIB EComm response to parse into map
	var body string
	for _, f := range textFields {
		body += fmt.Sprintf("%s: %s\n", f, textValue)
	}
	for _, f := range intFields {
		body += fmt.Sprintf("%s: %d\n", f, intValue)
	}

	parsed, err := parseBody(body)
	assert.Nil(t, err)

	// Verify that all fields have the correct value
	for _, f := range textFields {
		assert.Equal(t, textValue, parsed[f])
	}
	for _, f := range intFields {
		assert.Equal(t, intValue, parsed[f])
	}

	// Try to decode into all results
	// Prevents datatype inconsistencies
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
	_, err = requests.DecodeResponse[requests.ExecuteOneClickResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.RegisterOneClickResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.RegisterTransactionResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.ReverseTransactionResult](parsed)
	assert.Nil(t, err)
	_, err = requests.DecodeResponse[requests.TransactionStatusResult](parsed)
	assert.Nil(t, err)
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
