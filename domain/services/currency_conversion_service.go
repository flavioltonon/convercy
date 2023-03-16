package services

import (
	"convercy/domain"
	"convercy/domain/valueobject"
)

// CurrencyConversionService is an implementation of usecases.CurrencyConversionService interface
type CurrencyConversionService struct{}

func NewCurrencyConversionService() *CurrencyConversionService {
	return new(CurrencyConversionService)
}

func (s *CurrencyConversionService) ConvertCurrency(
	amount valueobject.CurrencyAmount,
	code valueobject.CurrencyCode,
	exchangeRate valueobject.ExchangeRate,
) (valueobject.CurrencyAmount, error) {
	if err := amount.Validate(); err != nil {
		return valueobject.CurrencyAmount{}, err
	}

	if err := code.Validate(); err != nil {
		return valueobject.CurrencyAmount{}, err
	}

	if err := exchangeRate.Validate(); err != nil {
		return valueobject.CurrencyAmount{}, err
	}

	if !exchangeRate.Unit().BaseCurrencyCode().Equal(code) {
		return valueobject.CurrencyAmount{}, domain.ErrUnexpectedExchangeRateBaseCurrencyCode(
			code.String(),
			exchangeRate.Unit().BaseCurrencyCode().String(),
		)
	}

	return valueobject.NewCurrencyAmount(amount.Value() * exchangeRate.Rate().Value())
}
