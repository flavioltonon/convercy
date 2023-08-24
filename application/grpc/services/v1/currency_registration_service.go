package v1

import (
	"context"

	"convercy/application"
	"convercy/application/dto"
	protov1 "convercy/application/grpc/proto/v1"

	"github.com/golang/protobuf/ptypes/empty"
)

type CurrencyRegistrationService struct {
	service application.CurrencyRegistrationService
}

func (s *CurrencyRegistrationService) RegisterCurrency(ctx context.Context, r *protov1.RegisterCurrencyRequest) (*protov1.RegisterCurrencyResponse, error) {
	request := dto.RegisterCurrencyRequest{
		Code: r.GetCode(),
	}

	response, err := s.service.RegisterCurrency(request)
	if err != nil {
		return nil, err
	}

	return &protov1.RegisterCurrencyResponse{
		Value: &protov1.Currency{
			Id:   response.ID,
			Code: request.Code,
		},
	}, nil
}

func (s *CurrencyRegistrationService) UnregisterCurrency(ctx context.Context, r *protov1.UnregisterCurrencyRequest) (*empty.Empty, error) {
	request := dto.UnregisterCurrencyRequest{
		CurrencyID: r.GetCurrencyId(),
	}

	if err := s.service.UnregisterCurrency(request); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *CurrencyRegistrationService) ListRegisteredCurrencies(ctx context.Context, _ *empty.Empty) (*protov1.ListRegisteredCurrenciesResponse, error) {
	registeredCurrencies, err := s.service.ListRegisteredCurrencies()
	if err != nil {
		return nil, err
	}

	response := &protov1.ListRegisteredCurrenciesResponse{
		Values: make([]*protov1.Currency, 0, len(registeredCurrencies)),
	}

	for _, currency := range registeredCurrencies {
		response.Values = append(response.Values, &protov1.Currency{
			Id:   currency.ID,
			Code: currency.Code,
		})
	}

	return response, nil
}
