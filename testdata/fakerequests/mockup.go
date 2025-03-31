// Package fakerequests is a mockup of the requests package.
// Without it the main package example would lead to an import cycle,
// because the real "requests" imports "maib", and the example imports "requests"
package fakerequests

import (
	"net/url"
)

func DecodeResponse[T any](any) (a T, b error) {
	return
}

const RegisterTransactionSMS = iota

type RegisterTransaction struct {
	TransactionType int
	Amount          int
	Currency        any
	ClientIPAddress string
	Description     string
	Language        any
}

type RegisterTransactionResult struct {
	TransactionID string
}

func (payload RegisterTransaction) Values() (a url.Values, b error) {
	return
}

type TransactionStatus struct {
	TransactionID   string
	ClientIPAddress string
}
type TransactionStatusResult struct {
	Result string
}

func (payload TransactionStatus) Values() (a url.Values, b error) {
	return
}
