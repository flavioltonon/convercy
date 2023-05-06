package repositories

import (
	"convercy/domain/aggregate"
	"convercy/domain/valueobject"
)

//go:generate mockery --name CurrencyExchangeRatesCache
type CurrencyExchangeRatesCache interface {
	GetCurrencyExchangeRates(code valueobject.CurrencyCode) (*aggregate.CurrencyExchangeRates, error)
	SaveCurrencyExchangeRates(currencyExchangeRates *aggregate.CurrencyExchangeRates) error
}
