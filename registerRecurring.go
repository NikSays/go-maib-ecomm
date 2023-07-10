package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const recurringSMSCommand = "z"
const recurringDMSCommand = "d"
const recurringNoPaymentCommand = "p"

// RegisterRecurringPayload contains data required to register or update a recurring transaction.
type RegisterRecurringPayload struct {
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

	// Language in which the bank payment page will be displayed.
	// One of: LanguageRomanian, LanguageRussian, LanguageEnglish.
	Language languageEnum `url:"language"`

	// Identifier of the recurring payment. If not specified,
	// resulting TRANSACTION_ID will be used as the recurring payment ID.
	BillerClientID string `url:"biller_client_id"`

	// Validity limit of the regular payment in the format "MMYY".
	PerspayeeExpiry string `url:"perspayee_expiry"`
}

// RegisterRecurringResult contains data returned after registration of a recurring transaction,
// if no error is encountered.
type RegisterRecurringResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionId string `mapstructure:"TRANSACTION_ID"`
}

// RegisterRecurringTransaction creates a new recurring transaction.
//
// If payload.Amount != 0, it executes the first payment with an [SMS] (-z) or a [DMS] (-d) transaction.
//
// If payload.Amount == 0, it does not execute the first payment (-p). (transactionType is ignored).
//
// For SMS transactions:
// The resulting transaction should be confirmed with [ECommClient.TransactionStatus] (-c).
//
// For DMS transactions:
// The resulting transaction should be confirmed with [ECommClient.TransactionStatus] (-c),
// and executed with [ECommClient.ExecuteDMS] (-t).
func (c *ECommClient) RegisterRecurringTransaction(transactionType transactionTypeEnum, payload RegisterRecurringPayload, updateExisting bool) (*RegisterRecurringResult, error) {
	// Validate transaction type
	var command string
	switch transactionType {
	case SMS:
		command = recurringSMSCommand
	case DMS:
		command = recurringDMSCommand
	default:
		return nil, errMalformedTransactionType
	}

	// Set NoPayment if amount is 0
	if payload.Amount == 0 {
		command = recurringNoPaymentCommand
	} else if !isValidAmount(payload.Amount) {
		return nil, errMalformedAmount
	}

	// Validate payload
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
	if !isValidBillerClientID(payload.BillerClientID) {
		return nil, errMalformedBillerClientID
	}
	if !isValidPerspayeeExpiry(payload.PerspayeeExpiry) {
		return nil, errMalformedPerspayeeExpiry
	}

	payloadValues, err := query.Values(payload)
	if err != nil {
		return nil, err
	}
	// Create or update recurring transaction
	if updateExisting {
		payloadValues.Set("perspayee_overwrite", "1")
	} else {
		payloadValues.Set("perspayee_gen", "1")
	}

	// Send command
	res, err := c.send(command, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &RegisterRecurringResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
