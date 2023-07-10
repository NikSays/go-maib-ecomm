package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const smsCommand = "v"
const dmsCommand = "a"

// RegisterPayload contains data required to register a transaction.
type RegisterPayload struct {
	// Transaction payment amount. Positive integer with last 2 digits being the cents.
	//
	// Example: if Amount:199 and Currency:CurrencyUSD, $1.99 will be requested from the client's card.
	Amount uint `url:"amount"`

	// Transaction currencyEnum.
	// One of: CurrencyMDL, CurrencyEUR, CurrencyUSD.
	Currency currencyEnum `url:"currency"`

	// Client's IP address in quad-dotted notation, like "127.0.0.1".
	ClientIpAddress string `url:"client_ip_addr"`

	// Transaction details. Optional.
	Description string `url:"description,omitempty"`

	// Language in which bank payment page will be displayed.
	// One of: LanguageRomanian, LanguageRussian, LanguageEnglish.
	Language languageEnum `url:"language"`
}

// RegisterResult contains data returned on execution of a transaction registration request,
// if no error is encountered.
type RegisterResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionId string `mapstructure:"TRANSACTION_ID"`
}

// RegisterTransaction creates a new [SMS] (-v) or [DMS] (-a) transaction.
//
// For SMS transactions:
// The resulting transaction should be confirmed with [ECommClient.TransactionStatus] (-c).
//
// For DMS transactions:
// The resulting transaction should be confirmed with [ECommClient.TransactionStatus] (-c),
// and executed with [ECommClient.ExecuteDMS] (-t).
func (c *ECommClient) RegisterTransaction(transactionType transactionTypeEnum, payload RegisterPayload) (*RegisterResult, error) {
	// Validate transaction type
	var command string
	switch transactionType {
	case SMS:
		command = smsCommand
	case DMS:
		command = dmsCommand
	default:
		return nil, errMalformedTransactionType
	}

	// Validate payload
	if !isValidAmount(payload.Amount) {
		return nil, errMalformedAmount
	}
	if !isValidCurrency(uint16(payload.Currency)) {
		return nil, errMalformedCurrency
	}
	if !isValidClientIpAddress(payload.ClientIpAddress) {
		return nil, errMalformedClientIP
	}
	if !isValidDescription(payload.Description) {
		return nil, errMalformedDescription
	}
	if !isValidLanguage(string(payload.Language)) {
		return nil, errMalformedLanguage
	}

	// Send command
	payloadValues, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	res, err := c.send(command, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &RegisterResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
