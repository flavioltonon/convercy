package usecases

import "convercy/domain/valueobject"

// ListExchangeRates returns a list of valueobject.ExchangeRates of different currencies using USD as the default target currency (e.g: 1 BRL = 0.1936 USD)
//
//go:generate mockery --name ExchangeRatesService
type ExchangeRatesService interface {
	ListExchangeRates() (valueobject.ExchangeRates, error)
}
