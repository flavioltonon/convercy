package valueobject

import "convercy/domain"

type ExchangeRates []ExchangeRate

func (v ExchangeRates) FindByBaseCurrencyCode(code CurrencyCode) (ExchangeRate, error) {
	for _, exchangeRate := range v {
		if exchangeRate.unit.BaseCurrencyCode().Equal(code) {
			return exchangeRate, nil
		}
	}

	return ExchangeRate{}, domain.ErrExchangeRateNotFound()
}

func (v ExchangeRates) FindByTargetCurrencyCode(code CurrencyCode) (ExchangeRate, error) {
	for _, exchangeRate := range v {
		if exchangeRate.unit.TargetCurrencyCode().Equal(code) {
			return exchangeRate, nil
		}
	}

	return ExchangeRate{}, domain.ErrExchangeRateNotFound()
}
