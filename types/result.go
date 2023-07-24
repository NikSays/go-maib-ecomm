package types

// ResultEnum holds possible values for RESULT field in response from MAIB EComm system.
type ResultEnum string

const (
	// ResultOk - the transaction is successfully completed.
	ResultOk ResultEnum = "OK"

	// ResultFailed - the transaction has failed.
	ResultFailed ResultEnum = "FAILED"

	// ResultCreated - the transaction is just registered in the system. Client didn't input card information yet.
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
