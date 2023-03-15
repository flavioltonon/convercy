package dto

import "convercy/domain/entity"

type RegisterCurrencyResponse Currency

func BuildRegisterCurrencyResponse(currency *entity.Currency) RegisterCurrencyResponse {
	return RegisterCurrencyResponse{
		ID:   currency.ID().String(),
		Code: currency.Code().String(),
	}
}
