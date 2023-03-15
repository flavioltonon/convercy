package openexchangerates

import "convercy/domain/valueobject"

type CurrenciesService struct {
	client *Client
}

func NewCurrenciesService(client *Client) *CurrenciesService {
	return &CurrenciesService{client: client}
}

func (s *CurrenciesService) ListCurrencyCodes() (valueobject.CurrencyCodes, error) {
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
