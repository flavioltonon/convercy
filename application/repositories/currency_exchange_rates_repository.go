package repositories

import (
	"convercy/domain/aggregate"
	"convercy/domain/valueobject"
)

//go:generate mockery --name CurrencyExchangeRatesRepository
type CurrencyExchangeRatesRepository interface {
	GetCurrencyExchangeRates(code valueobject.CurrencyCode) (*aggregate.CurrencyExchangeRates, error)
}
