package maib

// TransactionTypeEnum holds EComm transaction types
type TransactionTypeEnum uint8

const (
	// SMS (Single Message System) transaction.
	// A single transaction that allows you to transfer money from the client's account.
	SMS TransactionTypeEnum = iota

	// DMS (Dual Message System) transaction.
	// Debiting from the client's card occurs in two stages:
	// 1: Authorization (the funds on the client's card are blocked);
	// 2: Execution.
	DMS
)

// CurrencyEnum holds ISO4217 codes for currencies.
type CurrencyEnum uint16

const (
	// CurrencyMDL is ISO4217 code for Moldovan Lei.
	CurrencyMDL CurrencyEnum = 498

	// CurrencyEUR is ISO4217 code for Euro.
	CurrencyEUR CurrencyEnum = 978

	// CurrencyUSD is ISO4217 code for United States Dollar.
	CurrencyUSD CurrencyEnum = 840
)

// LanguageEnum holds 2 letter language identifiers.
type LanguageEnum string

const (
	// LanguageRomanian language identifier.
	LanguageRomanian LanguageEnum = "ro"

	// LanguageRussian language identifier.
	LanguageRussian LanguageEnum = "ru"

	// LanguageEnglish language identifier.
	LanguageEnglish LanguageEnum = "en"
)

// ResultEnum holds possible values for RESULT field in response from MAIB EComm system.
type ResultEnum string

const (
	// ResultOk - the transaction is successfully completed.
	ResultOk ResultEnum = "OK"

	// ResultFailed - the transaction has failed.
	ResultFailed ResultEnum = "FAILED"

	// ResultCreated - the transaction is just registered in the system.
	ResultCreated ResultEnum = "CREATED"

	// ResultPending - the transaction is not completed yet.
	ResultPending ResultEnum = "PENDING"

	// ResultDeclined - the transaction is declined by EComm.
	ResultDeclined ResultEnum = "DECLINED"

	// ResultReversed - the transaction is reversed.
	ResultReversed ResultEnum = "REVERSED"

	// ResultAutoReversed - the transaction is reversed by autoreversal.
	ResultAutoReversed ResultEnum = "AUTOREVERSED"

	// ResultTimeout - the transaction was timed out.
	ResultTimeout ResultEnum = "TIMEOUT"
)

// ResultPSEnum holds possible values for RESULT_PS field in response from MAIB EComm system.
type ResultPSEnum string

const (
	// ResultPSActive - the transaction was registered and payment is not completed yet.
	ResultPSActive ResultPSEnum = "ACTIVE"

	// ResultPSFinished - payment was completed successfully.
	ResultPSFinished ResultPSEnum = "FINISHED"

	// ResultPSCancelled - payment was cancelled.
	ResultPSCancelled ResultPSEnum = "CANCELLED"

	// ResultPSReturned - payment was returned.
	ResultPSReturned ResultPSEnum = "RETURNED"
)
