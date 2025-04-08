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

// Client allows sending requests to the MAIB ECommerce system using HTTPS with
// mutual TLS. It is safe for concurrent use.
//
// Client validates the [Request] before sending it, and checks the response for
// errors. The response is then parsed into a map that can be decoded into a
// result struct using requests.DecodeResponse.
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
		return nil, fmt.Errorf("read certificate: %w", err)
	}
	// Decode certificate
	privateKey, certificate, caArray, err := pkcs12.DecodeChain(pfxBytes, config.Passphrase)
	if err != nil {
		return nil, fmt.Errorf("load certificate: %w", err)
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
		MinVersion:   tls.VersionTLS12,
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Parse merchantHandlerEndpoint to check for malformed URL before any actual requests
	_, err = url.Parse(config.MerchantHandlerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("parse merchant handler endpoint: %w", err)
	}

	return &Client{
		httpClient:              httpClient,
		merchantHandlerEndpoint: config.MerchantHandlerEndpoint,
	}, nil
}
