package maib

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

// Request is a payload that can be sent to the ECommerce system.
type Request interface {
	// Values returns the payload as a URL value map,
	// that can be encoded into a querystring to be sent to the ECommerce system.
	Values() (url.Values, error)

	// Validate goes through the fields of the payload, and returns an error
	// if any one of them does not fit the requirements.
	Validate() error
}

// Sender allows sending requests to the MAIB ECommerce system.
//
// This interface is makes it easy to substitute [Client]
// with a different behavior. E.g. using a mock for testing.
//
// Send validates the [Request] before sending it, and checks the response
// for errors. The response is then parsed into a map that can be decoded
// into a result struct using requests.DecodeResponse.
type Sender interface {
	Send(req Request) (map[string]any, error)
}

// Client is a [Sender] that uses HTTPS with mutual TLS to communicate
// with the MAIB ECommerce system. It is safe for concurrent use.
//
// Must be initiated with [NewClient].
type Client struct {
	httpClient              *http.Client
	merchantHandlerEndpoint string
}

// Config is the configuration required to set up a [Client].
type Config struct {
	// Path to .pfx certificate issued by MAIB.
	PFXPath string
	// Passphrase to the certificate.
	Passphrase string
	// API communication URL issued by MAIB.
	MerchantHandlerEndpoint string
}

// NewClient reads and parses the PFX certificate file and returns a *[Client]
// that uses the certificate for mutual TLS.
func NewClient(config Config) (*Client, error) {
	// Read pfx certificate
	pfxBytes, err := os.ReadFile(config.PFXPath)
	if err != nil {
		return nil, fmt.Errorf("error reading certificate: %w", err)
	}
	// Decode certificate
	privateKey, certificate, caArray, err := pkcs12.DecodeChain(pfxBytes, config.Passphrase)
	if err != nil {
		return nil, fmt.Errorf("error loading certificate: %w", err)
	}
	// Parse CAs
	caPool := x509.NewCertPool()
	for _, v := range caArray {
		caPool.AddCert(v)
	}

	// Build client
	tlsCertificate := tls.Certificate{
		Certificate: [][]byte{certificate.Raw},
		PrivateKey:  privateKey,
		Leaf:        certificate,
	}
	tlsConfig := &tls.Config{
		ClientCAs:    caPool,
		Certificates: []tls.Certificate{tlsCertificate},
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Parse merchantHandlerEndpoint to check for malformed URL before any actual requests
	_, err = url.Parse(config.MerchantHandlerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing merchant handler endpoint: %w", err)
	}

	return &Client{
		httpClient:              httpClient,
		merchantHandlerEndpoint: config.MerchantHandlerEndpoint,
	}, nil
}
