package types

// Currency is an alias type for int. Valid values are 3 digit ISO4217 codes. The most common codes are exported by
// this package.
type Currency int

const (
	// CurrencyMDL is ISO4217 code for Moldovan Lei.
	CurrencyMDL Currency = 498

	// CurrencyEUR is ISO4217 code for Euro.
	CurrencyEUR Currency = 978

	// CurrencyUSD is ISO4217 code for United States Dollar.
	CurrencyUSD Currency = 840
)

// Language is an alias type for string. Valid values are language identifiers, that the merchant has sent to
// MAIB. The default identifiers are exported by this package.
type Language string

const (
	// LanguageRomanian language identifier.
	LanguageRomanian Language = "ro"

	// LanguageRussian language identifier.
	LanguageRussian Language = "ru"

	// LanguageEnglish language identifier.
	LanguageEnglish Language = "en"
)
