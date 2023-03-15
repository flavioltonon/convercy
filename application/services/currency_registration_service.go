package services

import (
	"errors"

	"convercy/application/dto"
	"convercy/application/repositories"
	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/domain/entity"
	"convercy/domain/usecases"
	"convercy/domain/valueobject"
)

type CurrencyRegistrationService struct {
	currencyCodeValidator          usecases.CurrencyCodeValidationService
	registeredCurrenciesRepository repositories.RegisteredCurrenciesRepository
}

func NewCurrencyRegistrationService(currencyCodeValidator usecases.CurrencyCodeValidationService, registeredCurrenciesRepository repositories.RegisteredCurrenciesRepository) *CurrencyRegistrationService {
	return &CurrencyRegistrationService{
		currencyCodeValidator:          currencyCodeValidator,
		registeredCurrenciesRepository: registeredCurrenciesRepository,
	}
}

// RegisterCurrency registers a new currency
func (s *CurrencyRegistrationService) RegisterCurrency(request dto.RegisterCurrencyRequest) (dto.RegisterCurrencyResponse, error) {
	code, err := valueobject.NewCurrencyCode(request.Code)
	if err != nil {
		return dto.RegisterCurrencyResponse{}, err
	}

	if err := s.currencyCodeValidator.ValidateCurrencyCode(code); err != nil {
		return dto.RegisterCurrencyResponse{}, err
	}

	// At this point, NewCurrency can never return an error because all other data has already been validated or is being generated systemically
	currency, _ := entity.NewCurrency(valueobject.GenerateCurrencyID(), code)

	registeredCurrencies, err := s.registeredCurrenciesRepository.GetRegisteredCurrencies()
	if err != nil {
		switch {
		case errors.As(err, new(domain.ErrNotFound)):
			registeredCurrencies = aggregate.NewRegisteredCurrencies(valueobject.GenerateClientID())
		default:
			return dto.RegisterCurrencyResponse{}, err
		}
	}

	if err := registeredCurrencies.RegisterCurrency(currency); err != nil {
		return dto.RegisterCurrencyResponse{}, err
	}

	if err := s.registeredCurrenciesRepository.SaveRegisteredCurrencies(registeredCurrencies); err != nil {
		return dto.RegisterCurrencyResponse{}, err
	}

	return dto.BuildRegisterCurrencyResponse(currency), nil
}

// UnregisterCurrency unregisters a registered currency with a given
func (s *CurrencyRegistrationService) UnregisterCurrency(request dto.UnregisterCurrencyRequest) error {
	currencyID, err := valueobject.NewCurrencyID(request.CurrencyID)
	if err != nil {
		return err
	}

	registeredCurrencies, err := s.registeredCurrenciesRepository.GetRegisteredCurrencies()
	if err != nil {
		return err
	}

	if err := registeredCurrencies.UnregisterCurrency(currencyID); err != nil {
		return err
	}

	if err := s.registeredCurrenciesRepository.SaveRegisteredCurrencies(registeredCurrencies); err != nil {
		return err
	}

	return nil
}
