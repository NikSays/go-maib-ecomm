package maib

import "fmt"

// Currency is an alias type for int. Valid values are 3 digit [ISO4217] codes. The most common codes are exported by
// this package.
//
// [ISO4217]: https://www.six-group.com/en/products-services/financial-information/data-standards.html
type Currency int

const (
	// CurrencyMDL is the ISO4217 code for Moldovan Lei.
	CurrencyMDL Currency = 498

	// CurrencyEUR is the ISO4217 code for Euro.
	CurrencyEUR Currency = 978

	// CurrencyUSD is the ISO4217 code for United States Dollar.
	CurrencyUSD Currency = 840
)

// Language is an alias type for string. Valid values are language identifiers that the merchant has sent to
// MAIB. The default identifiers are exported by this package.
type Language string

const (
	LanguageRomanian Language = "ro"
	LanguageRussian  Language = "ru"
	LanguageEnglish  Language = "en"
)

// PayloadField contains the names of the payload fields. Used in [ValidationError].
type PayloadField string

const (
	FieldTransactionID   PayloadField = "trans_id"
	FieldAmount          PayloadField = "amount"
	FieldCurrency        PayloadField = "currency"
	FieldClientIPAddress PayloadField = "client_ip_addr"
	FieldDescription     PayloadField = "description"
	FieldLanguage        PayloadField = "language"
	FieldBillerClientID  PayloadField = "biller_client_id"
	FieldPerspayeeExpiry PayloadField = "prespayee_expiry"
	FieldCommand         PayloadField = "command"
)

// ValidationError is triggered before sending the request to the
// ECommerce system, if the request has failed validation.
type ValidationError struct {
	// Which field is malformed.
	Field PayloadField

	// Human-readable explanation of the requirements.
	Description string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("malformed field %s: %s", e.Field, e.Description)
}
