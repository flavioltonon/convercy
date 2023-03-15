package usecases

import "convercy/domain/valueobject"

type CurrencyConversionService interface {
	ConvertCurrency(amount valueobject.CurrencyAmount, code valueobject.CurrencyCode, exchangeRate valueobject.ExchangeRate) (valueobject.CurrencyAmount, error)
}
