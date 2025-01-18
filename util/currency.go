package util

// constants for all supported currencies
const (
	MYR = "MYR"
	WON = "WON"
	EUR = "EUR"
)

// IsSupportedCurrencies will return true if the currencies is supported
func IsSupportedCurrencies(currency string) bool {
	switch currency {
	case MYR, WON, EUR:
		return true
	}
	return false
}
