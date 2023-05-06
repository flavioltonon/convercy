package services

import (
	"convercy/domain"
	"convercy/domain/entity"
	"convercy/domain/valueobject"
)

// CurrencyExchangeRatesService is an implementation of usecases.CurrencyExchangeRatesService interface
type CurrencyExchangeRatesService struct {
	target valueobject.CurrencyCode
}

func NewCurrencyExchangeRatesService(target valueobject.CurrencyCode) *CurrencyExchangeRatesService {
	return &CurrencyExchangeRatesService{
		target: target,
	}
}

func (s *CurrencyExchangeRatesService) ListCurrencyExchangeRates(currency *entity.Currency) (valueobject.ExchangeRates, error) {
	if err := currency.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}

// combineExchangeRates combines two different exchange rates (e.g. BRL/USD and EUR/USD -> BRL/EUR)
func (s *CurrencyExchangeRatesService) combineExchangeRates(
	base valueobject.ExchangeRate,
	target valueobject.ExchangeRate,
) (valueobject.ExchangeRate, error) {
	if !base.Unit().TargetCurrencyCode().Equal(target.Unit().TargetCurrencyCode()) {
		return valueobject.ExchangeRate{}, domain.ErrIncompatibleExchangeRates()
	}

	return valueobject.NewExchangeRate(
		base.Rate().Value()/target.Rate().Value(),
		base.Unit().BaseCurrencyCode().String(),
		target.Unit().BaseCurrencyCode().String(),
	)
}
