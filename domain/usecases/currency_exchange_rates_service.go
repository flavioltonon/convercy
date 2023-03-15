package usecases

import (
	"convercy/domain/entity"
	"convercy/domain/valueobject"
)

//go:generate mockery --name CurrencyExchangeRatesService
type CurrencyExchangeRatesService interface {
	ListCurrencyExchangeRates(currency *entity.Currency) (valueobject.ExchangeRates, error)
}
