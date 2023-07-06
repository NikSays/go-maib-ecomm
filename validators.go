package maib

import (
	"encoding/base64"
	"net"
	"strconv"
)

func isValidTransactionID(transactionID string) bool {
	if len(transactionID) != 28 {
		return false
	}
	if _, err := base64.StdEncoding.DecodeString(transactionID); err != nil {
		return false
	}
	return true
}

func isValidAmount(amount uint) bool {
	if amount < 1 || amount >= 1000000000000 {
		return false
	}
	return true
}

func isValidCurrency(currency uint16) bool {
	if currency > 999 {
		return false
	}
	return true
}

func isValidClientIpAddress(clientIpAddress string) bool {
	ip := net.ParseIP(clientIpAddress)
	if ip != nil {
		return true
	}
	return false
}
func isValidLanguage(language string) bool {
	if len(language) < 1 || len(language) > 32 {
		return false
	}
	return true
}

func isValidBillerClientID(billerClientID string) bool {
	if len(billerClientID) < 1 && len(billerClientID) > 49 {
		return false
	}
	return true
}

func isValidPerspayeeExpiry(perspayeeExpiry string) bool {
	if len(perspayeeExpiry) != 4 {
		return false
	}
	month, err := strconv.Atoi(perspayeeExpiry[0:2])
	if err != nil {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	year, err := strconv.Atoi(perspayeeExpiry[2:4])
	if err != nil {
		return false
	}
	if year < 0 {
		return false
	}
	return true
}

func isValidDescription(description string) bool {
	if len(description) > 125 {
		return false
	}
	return true
}
