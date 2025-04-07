package validators

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm/v2"
)

const validTransactionID = "abcdefghijklmnopqrstuvwxyz1="

func TestWithTransactionType(t *testing.T) {
	cases := []struct {
		name               string
		transactionType    string
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			transactionType:    "a",
			expectedErrorField: "",
		},
		{
			name:               "Invalid type",
			transactionType:    "",
			expectedErrorField: maib.FieldCommand,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithTransactionType(c.transactionType))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithTransactionID(t *testing.T) {
	const transactionIDLength = 28
	var invalidTransactionID = strings.Repeat("$", transactionIDLength)

	cases := []struct {
		name               string
		transactionID      string
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			transactionID:      validTransactionID,
			expectedErrorField: "",
		},
		{
			name:               "Invalid length",
			transactionID:      "",
			expectedErrorField: maib.FieldTransactionID,
		},
		{
			name:               "Not base64",
			transactionID:      invalidTransactionID,
			expectedErrorField: maib.FieldTransactionID,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithTransactionID(c.transactionID))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithAmount(t *testing.T) {
	cases := []struct {
		name               string
		amount             int
		required           bool
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			amount:             100,
			required:           true,
			expectedErrorField: "",
		},
		{
			name:               "OK (not required)",
			amount:             0,
			required:           false,
			expectedErrorField: "",
		},
		{
			name:               "Negative",
			amount:             -1,
			required:           false,
			expectedErrorField: maib.FieldAmount,
		},
		{
			name:               "Zero but required",
			amount:             0,
			required:           true,
			expectedErrorField: maib.FieldAmount,
		},
		{
			name:               "Too big",
			amount:             1000000000000,
			required:           true,
			expectedErrorField: maib.FieldAmount,
		},
		{
			name:               "Too big (not required)",
			amount:             1000000000000,
			required:           false,
			expectedErrorField: maib.FieldAmount,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithAmount(c.amount, c.required))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithCurrency(t *testing.T) {
	cases := []struct {
		name               string
		currency           maib.Currency
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			currency:           maib.CurrencyMDL,
			expectedErrorField: "",
		},
		{
			name:               "Negative",
			currency:           -1,
			expectedErrorField: maib.FieldCurrency,
		},
		{
			name:               "too big",
			currency:           1000,
			expectedErrorField: maib.FieldCurrency,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithCurrency(c.currency))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithClientIPAddress(t *testing.T) {
	cases := []struct {
		name               string
		clientIPAddress    string
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			clientIPAddress:    "127.0.0.1",
			expectedErrorField: "",
		},
		{
			name:               "Invalid IP",
			clientIPAddress:    "927.0.0.1",
			expectedErrorField: maib.FieldClientIPAddress,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithClientIPAddress(c.clientIPAddress))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithLanguage(t *testing.T) {
	const languageMaxLength = 32
	tooLongLanguage := strings.Repeat("-", languageMaxLength+1)
	cases := []struct {
		name               string
		language           maib.Language
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			language:           maib.LanguageEnglish,
			expectedErrorField: "",
		},
		{
			name:               "Too short",
			language:           "",
			expectedErrorField: maib.FieldLanguage,
		},
		{
			name:               "too long",
			language:           maib.Language(tooLongLanguage),
			expectedErrorField: maib.FieldLanguage,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithLanguage(c.language))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithBillerClientID(t *testing.T) {
	const billerClientIDMaxLength = 49
	tooLongBillerClientID := strings.Repeat("-", billerClientIDMaxLength+1)
	cases := []struct {
		name               string
		billerClientID     string
		required           bool
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			billerClientID:     "abc",
			required:           true,
			expectedErrorField: "",
		},
		{
			name:               "OK (not required)",
			billerClientID:     "",
			required:           false,
			expectedErrorField: "",
		},
		{
			name:               "Too short",
			billerClientID:     "",
			required:           true,
			expectedErrorField: maib.FieldBillerClientID,
		},
		{
			name:               "Too long",
			billerClientID:     tooLongBillerClientID,
			required:           true,
			expectedErrorField: maib.FieldBillerClientID,
		},
		{
			name:               "Too long (not required)",
			billerClientID:     tooLongBillerClientID,
			required:           false,
			expectedErrorField: maib.FieldBillerClientID,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithBillerClientID(c.billerClientID, c.required))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithPerspayeeExpiry(t *testing.T) {
	cases := []struct {
		name               string
		perspayeeExpiry    string
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			perspayeeExpiry:    "0624",
			expectedErrorField: "",
		},
		{
			name:               "Invalid length",
			perspayeeExpiry:    "0",
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (not int)",
			perspayeeExpiry:    "aa01",
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (not positive)",
			perspayeeExpiry:    "-201",
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (zero)",
			perspayeeExpiry:    "0001",
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (more than 12)",
			perspayeeExpiry:    "1301",
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid year (not int)",
			perspayeeExpiry:    "12aa",
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid year (not positive)",
			perspayeeExpiry:    "12-1",
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithPerspayeeExpiry(c.perspayeeExpiry))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func TestWithDescription(t *testing.T) {
	const descriptionMaxLength = 125
	tooLongDescription := strings.Repeat("-", descriptionMaxLength+1)
	cases := []struct {
		name               string
		description        string
		expectedErrorField maib.PayloadField
	}{
		{
			name:               "OK",
			description:        "description",
			expectedErrorField: "",
		},
		{
			name:               "OK (empty)",
			description:        "",
			expectedErrorField: "",
		},
		{
			name:               "Too long",
			description:        tooLongDescription,
			expectedErrorField: maib.FieldDescription,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithDescription(c.description))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}

func Example() {
	err := Validate(
		WithTransactionType("a"),
		WithTransactionID(validTransactionID),
		WithAmount(1000, true),
		WithCurrency(maib.CurrencyMDL),
		WithClientIPAddress("127.0.0.1"),
		WithLanguage(maib.LanguageEnglish),
		WithBillerClientID("biller", true),
		WithPerspayeeExpiry("1224"),
		WithDescription("description"))
	if err != nil {
		// Validation failed
	}
}

func BenchmarkValidateAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Example()
	}
}
