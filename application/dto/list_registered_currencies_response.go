package dto

import "convercy/domain/entity"

type ListRegisteredCurrenciesResponse []Currency

func BuildListRegisteredCurrenciesResponse(currencies []*entity.Currency) ListRegisteredCurrenciesResponse {
	response := make(ListRegisteredCurrenciesResponse, 0, len(currencies))

	for _, currency := range currencies {
		response = append(response, Currency{
			ID:   currency.ID().String(),
			Code: currency.Code().String(),
		})
	}

	return response
}
