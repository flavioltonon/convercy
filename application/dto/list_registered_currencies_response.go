package dto

import "convercy/domain/aggregate"

type ListRegisteredCurrenciesResponse []Currency

func BuildListRegisteredCurrenciesResponse(registeredCurrencies *aggregate.RegisteredCurrencies) ListRegisteredCurrenciesResponse {
	response := make(ListRegisteredCurrenciesResponse, 0, len(registeredCurrencies.Currencies()))

	for _, currency := range registeredCurrencies.Currencies() {
		response = append(response, Currency{
			ID:   currency.ID().String(),
			Code: currency.Code().String(),
		})
	}

	return response
}
