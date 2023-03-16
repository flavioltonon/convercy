package usecases

import (
	"convercy/domain/entity"
	"convercy/domain/valueobject"
)

// CurrencyExchangeRatesService is a service capable of providing exchange rates of a given currency in terms of all known currencies
//
//go:generate mockery --name CurrencyExchangeRatesService
type CurrencyExchangeRatesService interface {
	ListCurrencyExchangeRates(currency *entity.Currency) (valueobject.ExchangeRates, error)
}
