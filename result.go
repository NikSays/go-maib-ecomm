package maib

// ResultEnum holds the possible values of the RESULT field returned by the ECommerce system.
type ResultEnum string

const (
	// ResultOk - the transaction has successfully completed.
	ResultOk ResultEnum = "OK"

	// ResultFailed - the transaction has failed.
	ResultFailed ResultEnum = "FAILED"

	// ResultCreated - the transaction is just registered in the system. Client didn't input their card information yet.
	ResultCreated ResultEnum = "CREATED"

	// ResultPending - the transaction is not complete yet.
	ResultPending ResultEnum = "PENDING"

	// ResultDeclined - the transaction was declined by EComm.
	ResultDeclined ResultEnum = "DECLINED"

	// ResultReversed - the transaction was reversed.
	ResultReversed ResultEnum = "REVERSED"

	// ResultAutoReversed - the transaction was reversed by autoreversal.
	ResultAutoReversed ResultEnum = "AUTOREVERSED"

	// ResultTimeout - the transaction has timed out.
	ResultTimeout ResultEnum = "TIMEOUT"
)

// ResultPSEnum holds the possible values for the RESULT_PS field returned by the ECommerce system.
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
