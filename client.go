package maib

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

// Sender allows sending requests to the MAIB EComm system.
//
// Send validates the [Request] before sending it, and checks the response
// for errors. The response is then parsed into a map that can be decoded
// into a result struct using requests.DecodeResponse.
type Sender interface {
	Send(req Request) (map[string]any, error)
}

// client is the default [Sender].
type client struct {
	httpClient              *http.Client
	merchantHandlerEndpoint string
}

// Config is the configuration required to set up the default [Sender] implementation.
type Config struct {
	// Path to .pfx certificate issued by MAIB.
	PFXPath string
	// Passphrase to the certificate.
	Passphrase string
	// API communication URL issued by MAIB.
	MerchantHandlerEndpoint string
}

// NewClient creates a new instance of the default [Sender] that uses
// HTTPS with mutual TLS to communicate with the MAIB EComm system.
func NewClient(config Config) (Sender, error) {
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
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	return &client{
		httpClient:              httpClient,
		merchantHandlerEndpoint: config.MerchantHandlerEndpoint,
	}, nil
}
