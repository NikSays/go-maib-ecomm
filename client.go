package maib

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"software.sslmate.com/src/go-pkcs12"
)

// Sender sends the EComm request, parses the response into a map,
// and catches any errors during request execution.
//
// Useful if you want to substitute [Client] with a mock for testing.
type Sender interface {
	Send(req Request) (map[string]any, error)
}

// Client allows sending requests to MAIB ECommerce.
// It is a [Sender] that uses a http.Client with mutual TLS to communicate with the merchant handler.
type Client struct {
	httpClient              http.Client
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

// NewClient creates a new [Client].
func NewClient(config Config) (*Client, error) {
	// Read pfx certificate
	pfxBytes, err := os.ReadFile(config.PFXPath)
	if err != nil {
		return nil, fmt.Errorf("could not read certificate: %w", err)
	}
	// Decode certificate
	privateKey, certificate, caArray, err := pkcs12.DecodeChain(pfxBytes, config.Passphrase)
	if err != nil {
		return nil, fmt.Errorf("could not load certificate: %w", err)
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
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	client := &Client{
		httpClient:              httpClient,
		merchantHandlerEndpoint: config.MerchantHandlerEndpoint,
	}
	return client, nil
}
