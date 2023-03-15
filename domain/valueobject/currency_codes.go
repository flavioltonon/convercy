package valueobject

type CurrencyCodes []CurrencyCode

func (v CurrencyCodes) Contains(code CurrencyCode) bool {
	for _, currencyCode := range v {
		if currencyCode.Equal(code) {
			return true
		}
	}

	return false
}
