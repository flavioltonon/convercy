package services

import (
	"convercy/domain"
	"convercy/domain/valueobject"
)

type CurrencyConversionService struct{}

func NewCurrencyConversionService() *CurrencyConversionService {
	return new(CurrencyConversionService)
}

func (s *CurrencyConversionService) ConvertCurrency(
	amount valueobject.CurrencyAmount,
	code valueobject.CurrencyCode,
	exchangeRate valueobject.ExchangeRate,
) (valueobject.CurrencyAmount, error) {
	if !exchangeRate.Unit().BaseCurrencyCode().Equal(code) {
		return valueobject.CurrencyAmount{}, domain.ErrUnexpectedExchangeRateBaseCurrencyCode(
			code.String(),
			exchangeRate.Unit().BaseCurrencyCode().String(),
		)
	}

	return valueobject.NewCurrencyAmount(amount.Value() * exchangeRate.Rate().Value())
}
