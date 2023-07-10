package maib

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"software.sslmate.com/src/go-pkcs12"
)

// ECommClient is a client that implements methods for interaction with MAIB ECommerce API.
type ECommClient struct {
	httpClient              http.Client
	merchantHandlerEndpoint string
}

// Config is the configuration required to set up a [ECommClient].
type Config struct {
	// Path to .pfx certificate issued by MAIB.
	PFXPath string
	// Passphrase to the certificate.
	Passphrase string
	// API communication URL issued by MAIB.
	MerchantHandlerEndpoint string
}

// NewClient creates a new [ECommClient].
func NewClient(config Config) (*ECommClient, error) {
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
		// Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	client := &ECommClient{
		httpClient:              httpClient,
		merchantHandlerEndpoint: config.MerchantHandlerEndpoint,
	}
	return client, nil
}
