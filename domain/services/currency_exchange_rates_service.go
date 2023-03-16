package services

import (
	"convercy/domain"
	"convercy/domain/entity"
	"convercy/domain/usecases"
	"convercy/domain/valueobject"
)

// CurrencyExchangeRatesService is an implementation of usecases.CurrencyExchangeRatesService interface
type CurrencyExchangeRatesService struct {
	core usecases.ExchangeRatesService
}

func NewCurrencyExchangeRatesService(core usecases.ExchangeRatesService) *CurrencyExchangeRatesService {
	return &CurrencyExchangeRatesService{
		core: core,
	}
}

func (s *CurrencyExchangeRatesService) ListCurrencyExchangeRates(currency *entity.Currency) (valueobject.ExchangeRates, error) {
	if err := currency.Validate(); err != nil {
		return nil, err
	}

	exchangeRates, err := s.core.ListExchangeRates()
	if err != nil {
		return nil, err
	}

	baseExchangeRate, err := exchangeRates.FindByBaseCurrencyCode(currency.Code())
	if err != nil {
		return nil, err
	}

	relativeExchangeRates := make(valueobject.ExchangeRates, 0, len(exchangeRates))

	for _, targetExchangeRate := range exchangeRates {
		relativeExchangeRate, err := s.combineExchangeRates(baseExchangeRate, targetExchangeRate)
		if err != nil {
			return nil, err
		}

		relativeExchangeRates = append(relativeExchangeRates, relativeExchangeRate)
	}

	return relativeExchangeRates, nil
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
