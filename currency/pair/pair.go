package pair

import (
	"strings"
)

// // CurrencyItem is an exported string with methods to manipulate the data instead
// // of using array/slice access modifiers
// type CurrencyItem string

// Lower converts the CurrencyItem object c to lowercase
func LowerCurrencyItem(c string) string {
	return strings.ToLower(c)
}

// Upper converts the CurrencyItem object c to uppercase
func UpperCurrencyItem(c string) string {
	return strings.ToUpper(c)
}

// // CurrencyPair holds currency pair information
// type CurrencyPair struct {
// 	Delimiter      string       `json:"delimiter"`
// 	FirstCurrency  CurrencyItem `json:"first_currency"`
// 	SecondCurrency CurrencyItem `json:"second_currency"`
// }

// // GetFirstCurrencyItem returns the first currency item
// func (c *CurrencyPair) GetFirstCurrencyItem() CurrencyItem {
// 	return CurrencyItem(c.FirstCurrency)
// }

// // GetSecondCurrencyItem returns the second currency item
// func (c *CurrencyPair) GetSecondCurrencyItem() CurrencyItem {
// 	return CurrencyItem(c.SecondCurrency)
// }

// Pair returns a currency pair string
func (c *CurrencyPair) Pair() string {
	return c.FirstCurrency + c.Delimiter + c.SecondCurrency
}

// Display formats and returns the currency based on user preferences,
// overriding the default Pair() display
func (c *CurrencyPair) Display(delimiter string, uppercase bool) (pair string) {

	if delimiter != "" {
		pair = c.FirstCurrency + delimiter + c.SecondCurrency
	} else {
		pair = c.FirstCurrency + c.SecondCurrency
	}

	if uppercase {
		return UpperCurrencyItem(pair)
	}
	return LowerCurrencyItem(pair)
}

// Equal compares two currency pairs and returns whether or not they are equal
func (c *CurrencyPair) Equal(p CurrencyPair) bool {
	if UpperCurrencyItem(c.FirstCurrency) == UpperCurrencyItem(p.FirstCurrency) &&
		UpperCurrencyItem(c.SecondCurrency) == UpperCurrencyItem(p.SecondCurrency) ||
		UpperCurrencyItem(c.FirstCurrency) == UpperCurrencyItem(p.SecondCurrency) &&
			UpperCurrencyItem(c.SecondCurrency) == UpperCurrencyItem(p.FirstCurrency) {
		return true
	}
	return false
}

// NewCurrencyPairDelimiter splits the desired currency string at delimeter,
// the returns a CurrencyPair struct
func NewCurrencyPairDelimiter(currency, delimiter string) CurrencyPair {
	result := strings.Split(currency, delimiter)
	return CurrencyPair{
		Delimiter:      delimiter,
		FirstCurrency:  result[0],
		SecondCurrency: result[1],
	}
}

// NewCurrencyPair returns a CurrencyPair without a delimiter
func NewCurrencyPair(firstCurrency, secondCurrency string) CurrencyPair {
	return CurrencyPair{
		FirstCurrency:  firstCurrency,
		SecondCurrency: secondCurrency,
	}
}

// NewCurrencyPairFromIndex returns a CurrencyPair via a currency string and
// specific index
func NewCurrencyPairFromIndex(currency, index string) CurrencyPair {
	i := strings.Index(currency, index)
	if i == 0 {
		return NewCurrencyPair(currency[0:len(index)], currency[len(index):])
	}
	return NewCurrencyPair(currency[0:i], currency[i:])
}

// NewCurrencyPairFromString converts currency string into a new CurrencyPair
// with or without delimeter
func NewCurrencyPairFromString(currency string) CurrencyPair {
	delimiters := []string{"_", "-"}
	var delimiter string
	for _, x := range delimiters {
		if strings.Contains(currency, x) {
			delimiter = x
			return NewCurrencyPairDelimiter(currency, delimiter)
		}
	}
	return NewCurrencyPair(currency[0:3], currency[3:])
}

// Contains checks to see if a specified pair exists inside a currency pair
// array
func Contains(pairs []CurrencyPair, p CurrencyPair) bool {
	for x := range pairs {
		if pairs[x].Equal(p) {
			return true
		}
	}
	return false
}

// FormatPairs formats a string array to a list of currency pairs with the
// supplied currency pair format
func FormatPairs(pairs []string, delimiter, index string) []CurrencyPair {
	var result []CurrencyPair
	for x := range pairs {
		if pairs[x] == "" {
			continue
		}
		var p CurrencyPair
		if delimiter != "" {
			p = NewCurrencyPairDelimiter(pairs[x], delimiter)
		} else {
			if index != "" {
				p = NewCurrencyPairFromIndex(pairs[x], index)
			} else {
				p = NewCurrencyPair(pairs[x][0:3], pairs[x][3:])
			}
		}
		result = append(result, p)
	}
	return result
}

// CopyPairFormat copies the pair format from a list of pairs once matched
func CopyPairFormat(p CurrencyPair, pairs []CurrencyPair) CurrencyPair {
	for x := range pairs {
		if p.Equal(pairs[x]) {
			return pairs[x]
		}
	}
	return CurrencyPair{}
}
