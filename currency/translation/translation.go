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
	for k, v := range translations {
		if k == currency {
			return v, nil
		}

		if v == currency {
			return k, nil
		}
	}
	return "", errors.New("no translation found for specified currency")
}

// HasTranslation returns whether or not a particular currency has a translation
func HasTranslation(currency string) bool {
	_, err := GetTranslation(currency)
	if err != nil {
		return false
	}
	return true
}
