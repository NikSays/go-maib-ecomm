package maib

import (
	"github.com/mitchellh/mapstructure"
)

const closeDayCommand = "b"

// CloseDayResult contains data returned on closing of the business day,
// if no error is encountered.
type CloseDayResult struct {
	// Transaction result status.
	Result resultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// Number of credit transactions (fld_074, max 10 digits).
	// Available only if resultCode begins with 5.
	CreditTransactionNumber int `mapstructure:"fld_074"`
	// Number of credit reversals (fld_075, max 10 digits).
	// Available only if resultCode begins with 5.
	CreditReversalNumber int `mapstructure:"fld_075"`
	// Number of debit transactions (fld_076, max 10 digits).
	// Available only if resultCode begins with 5.
	DebitTransactionNumber int `mapstructure:"fld_076"`
	// Number of debit reversals (fld_077, max 10 digits).
	// Available only if resultCode begins with 5.
	DebitReversalNumber int `mapstructure:"fld_077"`

	// Total amount of credit transactions (fld_086, max 16 digits).
	// Available only if resultCode begins with 5.
	CreditTransactionAmount int `mapstructure:"fld_086"`
	// Total amount of credit reversals (fld_087, max 16 digits).
	// Available only if resultCode begins with 5.
	CreditReversalAmount int `mapstructure:"fld_087"`
	// Total amount of debit transactions (fld_088, max 16 digits).
	// Available only if resultCode begins with 5.
	DebitTransactionAmount int `mapstructure:"fld_088"`
	// Total amount of debit reversals (fld_089, max 16 digits).
	// Available only if resultCode begins with 5.
	DebitReversalAmount int `mapstructure:"fld_089"`
}

// CloseDay closes the business day (-b).
// This procedure must be initiated once a day. Recommended time is 23:59:00.
func (c *ECommClient) CloseDay() (*CloseDayResult, error) {
	res, err := c.send(closeDayCommand, "")
	if err != nil {
		return nil, err
	}
	result := &CloseDayResult{}
	err = mapstructure.Decode(&res, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
