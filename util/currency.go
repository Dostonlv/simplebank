package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	UZS = "UZS"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, UZS:
		return true
	}
	return false
}
