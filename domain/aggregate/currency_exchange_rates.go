package aggregate

import (
	"convercy/domain"
	"convercy/domain/valueobject"
)

type CurrencyExchangeRates struct {
	baseCurrencyCode valueobject.CurrencyCode
	exchangeRates    valueobject.ExchangeRates
}

func NewCurrencyExchangeRates(baseCurrencyCode valueobject.CurrencyCode, exchangeRates ...valueobject.ExchangeRate) (*CurrencyExchangeRates, error) {
	currencyExchangeRates := &CurrencyExchangeRates{
		baseCurrencyCode: baseCurrencyCode,
		exchangeRates:    make([]valueobject.ExchangeRate, 0, len(exchangeRates)),
	}

	for _, exchangeRate := range exchangeRates {
		if err := currencyExchangeRates.AddExchangeRate(exchangeRate); err != nil {
			return nil, err
		}
	}

	return currencyExchangeRates, nil
}

func (a *CurrencyExchangeRates) CurrencyCode() valueobject.CurrencyCode {
	return a.baseCurrencyCode
}

func (a *CurrencyExchangeRates) ExchangeRates() valueobject.ExchangeRates {
	return a.exchangeRates
}

func (a *CurrencyExchangeRates) AddExchangeRate(newExchangeRate valueobject.ExchangeRate) error {
	if baseCurrencyCode := newExchangeRate.Unit().BaseCurrencyCode(); !baseCurrencyCode.Equal(a.baseCurrencyCode) {
		return domain.ErrUnexpectedExchangeRateBaseCurrencyCode(
			a.baseCurrencyCode.String(), // expected
			baseCurrencyCode.String(),   // actual
		)
	}

	// The currency exchange rates should contain only a single occurrence per unit
	for _, exchangeRate := range a.exchangeRates {
		if newExchangeRate.Unit().Equal(exchangeRate.Unit()) {
			return domain.ErrExchangeRateAlreadyExists()
		}
	}

	a.exchangeRates = append(a.exchangeRates, newExchangeRate)
	return nil
}
