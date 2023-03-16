package usecases

import "convercy/domain/valueobject"

// CurrencyExchangeRatesService is a service capable of providing all known exchange rates
//
//go:generate mockery --name ExchangeRatesService
type ExchangeRatesService interface {
	// ListExchangeRates returns a list of valueobject.ExchangeRates of different currencies using USD as the default target currency (e.g: 1 BRL = 0.1936 USD)
	ListExchangeRates() (valueobject.ExchangeRates, error)
}
