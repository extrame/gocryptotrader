package translation

import (
	"errors"
)

var translations = map[string]string{
	"BTC":  "XBT",
	"ETH":  "XETH",
	"DOGE": "XDG",
	"USD":  "USDT",
}

// GetTranslation returns similar strings for a particular currency
func GetTranslation(currency string) (string, error) {
	result, ok := translations[currency]
	if !ok {
		return "", errors.New("no translation found for specified currency")
	}

	return result, nil
}

// HasTranslation returns whether or not a particular currency has a translation
func HasTranslation(currency string) bool {
	_, ok := translations[currency]
	if !ok {
		return false
	}
	return true
}
