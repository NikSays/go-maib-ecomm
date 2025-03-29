package validators

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm/types"
)

const validTransactionID = "abcdefghijklmnopqrstuvwxyz1="

func TestWithTransactionType(t *testing.T) {
	cases := []struct {
		name               string
		transactionType    string
		expectedErrorField types.PayloadField
	}{
		{
			name:               "OK",
			transactionType:    "a",
			expectedErrorField: "",
		},
		{
			name:               "Invalid type",
			transactionType:    "",
			expectedErrorField: types.FieldCommand,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithTransactionType(c.transactionType))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
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
		expectedErrorField types.PayloadField
	}{
		{
			name:               "OK",
			transactionID:      validTransactionID,
			expectedErrorField: "",
		},
		{
			name:               "Invalid length",
			transactionID:      "",
			expectedErrorField: types.FieldTransactionID,
		},
		{
			name:               "Not base64",
			transactionID:      invalidTransactionID,
			expectedErrorField: types.FieldTransactionID,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithTransactionID(c.transactionID))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
			}
		})
	}
}

func TestWithAmount(t *testing.T) {
	cases := []struct {
		name               string
		amount             uint
		required           bool
		expectedErrorField types.PayloadField
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
			name:               "Too small",
			amount:             0,
			required:           true,
			expectedErrorField: types.FieldAmount,
		},
		{
			name:               "Too big",
			amount:             1000000000000,
			required:           true,
			expectedErrorField: types.FieldAmount,
		},
		{
			name:               "Too big (not required)",
			amount:             1000000000000,
			required:           false,
			expectedErrorField: types.FieldAmount,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithAmount(c.amount, c.required))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
			}
		})
	}
}

func TestWithCurrency(t *testing.T) {
	cases := []struct {
		name               string
		currency           types.Currency
		expectedErrorField types.PayloadField
	}{
		{
			name:               "OK",
			currency:           types.CurrencyMDL,
			expectedErrorField: "",
		},
		{
			name:               "Negative",
			currency:           -1,
			expectedErrorField: types.FieldCurrency,
		},
		{
			name:               "too big",
			currency:           1000,
			expectedErrorField: types.FieldCurrency,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithCurrency(c.currency))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
			}
		})
	}
}

func TestWithClientIPAddress(t *testing.T) {
	cases := []struct {
		name               string
		clientIPAddress    string
		expectedErrorField types.PayloadField
	}{
		{
			name:               "OK",
			clientIPAddress:    "127.0.0.1",
			expectedErrorField: "",
		},
		{
			name:               "Invalid IP",
			clientIPAddress:    "927.0.0.1",
			expectedErrorField: types.FieldClientIPAddress,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithClientIPAddress(c.clientIPAddress))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
			}
		})
	}
}

func TestWithLanguage(t *testing.T) {
	const languageMaxLength = 32
	tooLongLanguage := strings.Repeat("-", languageMaxLength+1)
	cases := []struct {
		name               string
		language           types.Language
		expectedErrorField types.PayloadField
	}{
		{
			name:               "OK",
			language:           types.LanguageEnglish,
			expectedErrorField: "",
		},
		{
			name:               "Too short",
			language:           "",
			expectedErrorField: types.FieldLanguage,
		},
		{
			name:               "too long",
			language:           types.Language(tooLongLanguage),
			expectedErrorField: types.FieldLanguage,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithLanguage(c.language))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
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
		expectedErrorField types.PayloadField
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
			expectedErrorField: types.FieldBillerClientID,
		},
		{
			name:               "Too long",
			billerClientID:     tooLongBillerClientID,
			required:           true,
			expectedErrorField: types.FieldBillerClientID,
		},
		{
			name:               "Too long (not required)",
			billerClientID:     tooLongBillerClientID,
			required:           false,
			expectedErrorField: types.FieldBillerClientID,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithBillerClientID(c.billerClientID, c.required))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
			}
		})
	}
}

func TestWithPerspayeeExpiry(t *testing.T) {
	cases := []struct {
		name               string
		perspayeeExpiry    string
		expectedErrorField types.PayloadField
	}{
		{
			name:               "OK",
			perspayeeExpiry:    "0624",
			expectedErrorField: "",
		},
		{
			name:               "Invalid length",
			perspayeeExpiry:    "0",
			expectedErrorField: types.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (not int)",
			perspayeeExpiry:    "aa01",
			expectedErrorField: types.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (not positive)",
			perspayeeExpiry:    "-201",
			expectedErrorField: types.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (zero)",
			perspayeeExpiry:    "0001",
			expectedErrorField: types.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid month (more than 12)",
			perspayeeExpiry:    "1301",
			expectedErrorField: types.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid year (not int)",
			perspayeeExpiry:    "12aa",
			expectedErrorField: types.FieldPerspayeeExpiry,
		},
		{
			name:               "Invalid year (not positive)",
			perspayeeExpiry:    "12-1",
			expectedErrorField: types.FieldPerspayeeExpiry,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithPerspayeeExpiry(c.perspayeeExpiry))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
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
		expectedErrorField types.PayloadField
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
			expectedErrorField: types.FieldDescription,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(WithDescription(c.description))
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*types.ValidationError).Field)
			}
		})
	}
}

func Example() {
	err := Validate(
		WithTransactionType("a"),
		WithTransactionID(validTransactionID),
		WithAmount(1000, true),
		WithCurrency(types.CurrencyMDL),
		WithClientIPAddress("127.0.0.1"),
		WithLanguage(types.LanguageEnglish),
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
