package maib

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	caPath = "testdata/certs/ca.crt"

	serverCertPath = "testdata/certs/server.crt"
	serverKeyPath  = "testdata/certs/server.key"

	testCommand = "q"
)

var ctx = context.Background()

type testRequest struct {
	isValid bool
}

func (t testRequest) Values() (url.Values, error) {
	if t.isValid {
		return map[string][]string{"command": {testCommand}}, nil
	} else {
		return nil, &ValidationError{
			Field:       FieldClientIPAddress,
			Description: "invalid request",
		}
	}
}

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
	res, err := client.Send(ctx, testRequest{false})

	assert.Nil(t, res)
	assert.ErrorAs(t, err, new(*ValidationError))
}

func TestClient_Send_InvalidEndpoint(t *testing.T) {
	client := Client{
		merchantHandlerEndpoint: ":",
	}
	res, err := client.Send(ctx, testRequest{true})

	var urlErr *url.Error
	assert.Nil(t, res)
	assert.ErrorAs(t, err, &urlErr)
}

func TestClient_Send_InvalidContext(t *testing.T) {
	client := Client{}
	res, err := client.Send(nil, testRequest{true})

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestClient_Send_WithCerts(t *testing.T) {
	caPool, serverCert, err := loadCerts()
	assert.Nil(t, err)

	t.Run("OK", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, testCommand, request.FormValue("command"))
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		_, err = client.Send(ctx, testRequest{true})
		assert.Nil(t, err)
	})

	t.Run("TLS fail", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, &x509.CertPool{})
		assert.Nil(t, err)

		_, err = client.Send(ctx, testRequest{true})
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

		_, err = client.Send(ctx, testRequest{true})
		assert.ErrorAs(t, err, new(*ECommError))
	})

	t.Run("Error response", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("error: ecommerce has encountered an unknown error"))
			assert.Nil(t, err)
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		_, err = client.Send(ctx, testRequest{true})
		assert.ErrorAs(t, err, new(*ECommError))
	})

	t.Run("Malformed response", func(t *testing.T) {
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("welcome"))
			assert.Nil(t, err)
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		_, err = client.Send(ctx, testRequest{true})
		assert.ErrorAs(t, err, new(*ParseError))
	})

	t.Run("Timeout", func(t *testing.T) {
		const timeout = 100 * time.Millisecond
		server := createServer(caPool, serverCert, func(writer http.ResponseWriter, request *http.Request) {
			time.Sleep(2 * timeout)
		})
		server.StartTLS()
		client, err := createTrustingClient(server.URL, caPool)
		assert.Nil(t, err)

		timeoutCtx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		_, err = client.Send(timeoutCtx, testRequest{true})
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}
