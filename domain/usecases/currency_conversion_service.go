package usecases

import "convercy/domain/valueobject"

// CurrencyConversionService is a service capable of converting currency amounts from a currency to another
type CurrencyConversionService interface {
	// ConvertCurrency converts a given currency amount to a different currency with a given exchange rate
	ConvertCurrency(amount valueobject.CurrencyAmount, code valueobject.CurrencyCode, exchangeRate valueobject.ExchangeRate) (valueobject.CurrencyAmount, error)
}
