package util

const (
	USD = "USD"
	EUR = "EUR"
	KZ = "KZ"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, KZ, CAD:
		return true
	}
	return false
}
