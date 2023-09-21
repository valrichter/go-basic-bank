package util

const (
	USD = "USD"
	EUR = "EUR"
	ARS = "ARS"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, ARS:
		return true
	}
	return false
}
