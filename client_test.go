package maib

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"software.sslmate.com/src/go-pkcs12"
)

const (
	clientCertPath = "testdata/certs/client.pfx"
	clientCertPass = "password"
)

func TestNewClient_OK(t *testing.T) {
	_, err := NewClient(Config{
		PFXPath:                 clientCertPath,
		Passphrase:              clientCertPass,
		MerchantHandlerEndpoint: "",
	})

	assert.Nil(t, err)
}

func TestNewClient_InvalidPath(t *testing.T) {
	_, err := NewClient(Config{
		PFXPath:                 clientCertPath + "wrong",
		Passphrase:              clientCertPass,
		MerchantHandlerEndpoint: "",
	})

	assert.ErrorIs(t, err, os.ErrNotExist)
}

func TestNewClient_InvalidPass(t *testing.T) {
	_, err := NewClient(Config{
		PFXPath:                 clientCertPath,
		Passphrase:              clientCertPass + "wrong",
		MerchantHandlerEndpoint: "",
	})

	assert.ErrorIs(t, err, pkcs12.ErrIncorrectPassword)
}
func TestNewClient_InvalidEndpoint(t *testing.T) {
	_, err := NewClient(Config{
		PFXPath:                 clientCertPath,
		Passphrase:              clientCertPass,
		MerchantHandlerEndpoint: ":",
	})

	var urlErr *url.Error
	assert.ErrorAs(t, err, &urlErr)
}
