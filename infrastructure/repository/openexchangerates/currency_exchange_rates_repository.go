package openexchangerates

import (
	"convercy/domain/aggregate"
	"convercy/domain/valueobject"
)

// CurrencyExchangeRatesRepository is an implementation of usecases.CurrencyExchangeRatesRepository interface
type CurrencyExchangeRatesRepository struct {
	baseCurrencyCode valueobject.CurrencyCode
	client           *Client
}

func NewCurrencyExchangeRatesRepository(client *Client) *CurrencyExchangeRatesRepository {
	// A free OpenExchangeRates account does not allow changes in the latest exchange rates base currency, which defaults to USD
	baseCurrencyCode, _ := valueobject.NewCurrencyCode("USD")

	return &CurrencyExchangeRatesRepository{
		baseCurrencyCode: baseCurrencyCode,
		client:           client,
	}
}

func (s *CurrencyExchangeRatesRepository) GetCurrencyExchangeRates(code valueobject.CurrencyCode) (*aggregate.CurrencyExchangeRates, error) {
	baseCurrencyExchangeRates, err := s.getBaseCurrencyExchangeRates()
	if err != nil {
		return nil, err
	}

	if s.baseCurrencyCode.Equal(code) {
		return baseCurrencyExchangeRates, nil
	}

	targetExchangeRates := baseCurrencyExchangeRates.ExchangeRates()

	// baseExchangeRate is the exchange rate of the base currency in terms of the input currency (e.g. 1 USD = 5 BRL)
	baseExchangeRate, err := targetExchangeRates.FindByTargetCurrencyCode(code)
	if err != nil {
		return nil, err
	}

	// currencyBaseExchangeRate is the exchange rate of the input currency in terms of the base currency (e.g. 1 BRL = 0.2 USD)
	currencyBaseExchangeRate := baseExchangeRate.Inverse()

	relativeExchangeRates := make(valueobject.ExchangeRates, 0, len(targetExchangeRates))

	for _, targetExchangeRate := range targetExchangeRates {
		relativeExchangeRate, err := valueobject.NewExchangeRate(
			currencyBaseExchangeRate.Rate().Value()*targetExchangeRate.Rate().Value(),
			currencyBaseExchangeRate.Unit().BaseCurrencyCode().String(),
			targetExchangeRate.Unit().TargetCurrencyCode().String(),
		)
		if err != nil {
			return nil, err
		}

		relativeExchangeRates = append(relativeExchangeRates, relativeExchangeRate)
	}

	return aggregate.NewCurrencyExchangeRates(code, relativeExchangeRates...)
}

func (s *CurrencyExchangeRatesRepository) getBaseCurrencyExchangeRates() (*aggregate.CurrencyExchangeRates, error) {
	response, err := s.client.GetLatestExchangeRates()
	if err != nil {
		return nil, err
	}

	exchangeRates := make(valueobject.ExchangeRates, 0, len(response.Rates))

	for code, rate := range response.Rates {
		exchangeRate, err := valueobject.NewExchangeRate(rate, response.Base, code)
		if err != nil {
			return nil, err
		}

		exchangeRates = append(exchangeRates, exchangeRate)
	}

	return aggregate.NewCurrencyExchangeRates(s.baseCurrencyCode, exchangeRates...)
}
