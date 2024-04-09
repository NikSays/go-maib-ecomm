package requests

import (
	"net/url"

	"github.com/NikSays/go-maib-ecomm/types"
)

const closeDayCommand = "b"

// CloseDay closes the business day (-b).
// This procedure must be initiated once a day. Recommended time is 23:59:00.
type CloseDay struct{}

// CloseDayResult contains the response to a CloseDay request.
type CloseDayResult struct {
	// Transaction result status.
	Result types.ResultEnum `mapstructure:"RESULT"`

	// Transaction result code returned from Card Suite FO (3 digits).
	ResultCode int `mapstructure:"RESULT_CODE"`

	// Number of credit transactions (FLD_074, max 10 digits).
	// Available only if resultCode begins with 5.
	CreditTransactionNumber int `mapstructure:"FLD_074"`
	// Number of credit reversals (FLD_075, max 10 digits).
	// Available only if resultCode begins with 5.
	CreditReversalNumber int `mapstructure:"FLD_075"`
	// Number of debit transactions (FLD_076, max 10 digits).
	// Available only if resultCode begins with 5.
	DebitTransactionNumber int `mapstructure:"FLD_076"`
	// Number of debit reversals (FLD_077, max 10 digits).
	// Available only if resultCode begins with 5.
	DebitReversalNumber int `mapstructure:"FLD_077"`

	// Total amount of credit transactions (FLD_086, max 16 digits).
	// Available only if resultCode begins with 5.
	CreditTransactionAmount int `mapstructure:"FLD_086"`
	// Total amount of credit reversals (FLD_087, max 16 digits).
	// Available only if resultCode begins with 5.
	CreditReversalAmount int `mapstructure:"FLD_087"`
	// Total amount of debit transactions (FLD_088, max 16 digits).
	// Available only if resultCode begins with 5.
	DebitTransactionAmount int `mapstructure:"FLD_088"`
	// Total amount of debit reversals (FLD_089, max 16 digits).
	// Available only if resultCode begins with 5.
	DebitReversalAmount int `mapstructure:"FLD_089"`
}

func (CloseDay) Values() (url.Values, error) {
	v := url.Values{}
	v.Set("command", closeDayCommand)
	return v, nil
}

func (CloseDay) Validate() error {
	return nil
}
