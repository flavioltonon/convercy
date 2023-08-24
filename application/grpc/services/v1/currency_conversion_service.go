package v1

import (
	"context"

	"convercy/application"
	"convercy/application/dto"
	protov1 "convercy/application/grpc/proto/v1"
)

type CurrencyConversionService struct {
	service application.CurrencyConversionService
}

func (s *CurrencyConversionService) ConvertCurrency(ctx context.Context, r *protov1.ConvertCurrencyRequest) (*protov1.ConvertCurrencyResponse, error) {
	request := dto.ConvertCurrencyRequest{
		Amount: r.GetAmount(),
		Code:   r.GetCode(),
	}

	serviceResponse, err := s.service.ConvertCurrency(request)
	if err != nil {
		return nil, err
	}

	response := &protov1.ConvertCurrencyResponse{
		Values: make(map[string]float64, len(serviceResponse)),
	}

	for k, v := range serviceResponse {
		response.Values[k] = v
	}

	return response, nil
}
