package maib

import (
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
)

const dmsAuthCommand = "a"

// DMSAuthPayload contains data required to register an DMS authorization.
type DMSAuthPayload struct {
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

// DMSAuthResult contains data returned on registration of an DMS transaction,
// if no error is encountered.
type DMSAuthResult struct {
	// ID of the created transaction. 28 symbols in base64.
	TransactionId string `mapstructure:"TRANSACTION_ID"`
}

// AuthorizeDMS creates a new DMS transaction (-a). The resulting transaction should be
// confirmed with [ECommClient.FetchStatus] (-c), and executed with [ECommClient.ExecuteDMS] (-t).
func (c *ECommClient) AuthorizeDMS(payload DMSAuthPayload) (*DMSAuthResult, error) {
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
	res, err := c.send(dmsAuthCommand, payloadValues.Encode())
	if err != nil {
		return nil, err
	}
	result := &DMSAuthResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
