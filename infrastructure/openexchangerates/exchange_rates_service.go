package openexchangerates

import (
	"convercy/domain/valueobject"
)

type ExchangeRatesService struct {
	client *Client
}

func NewExchangeRatesService(client *Client) *ExchangeRatesService {
	return &ExchangeRatesService{client: client}
}

// ListExchangeRates returns a list of valueobject.ExchangeRates of different currencies using USD as the default target currency (e.g: 1 BRL = 0.1936 USD)
func (s *ExchangeRatesService) ListExchangeRates() (valueobject.ExchangeRates, error) {
	response, err := s.client.GetLatestExchangeRates()
	if err != nil {
		return nil, err
	}

	exchangeRates := make(valueobject.ExchangeRates, 0, len(response.Rates))

	for code, rate := range response.Rates {
		exchangeRate, err := valueobject.NewExchangeRate(1/rate, code, response.Base)
		if err != nil {
			return nil, err
		}

		exchangeRates = append(exchangeRates, exchangeRate)
	}

	return exchangeRates, nil
}
