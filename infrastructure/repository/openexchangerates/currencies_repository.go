package openexchangerates

import "convercy/domain/valueobject"

// CurrenciesRepository is an implementation of usecases.CurrenciesRepository interface
type CurrenciesRepository struct {
	client *Client
}

func NewCurrenciesRepository(client *Client) *CurrenciesRepository {
	return &CurrenciesRepository{client: client}
}

func (s *CurrenciesRepository) ListCurrencyCodes() (valueobject.CurrencyCodes, error) {
	response, err := s.client.GetCurrencies()
	if err != nil {
		return nil, err
	}

	currencyCodes := make(valueobject.CurrencyCodes, 0, len(response))

	for code := range response {
		currencyCode, err := valueobject.NewCurrencyCode(code)
		if err != nil {
			return nil, err
		}

		currencyCodes = append(currencyCodes, currencyCode)
	}

	return currencyCodes, nil
}
