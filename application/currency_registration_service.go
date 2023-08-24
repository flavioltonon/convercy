package application

import (
	"errors"

	"convercy/application/dto"
	"convercy/application/repositories"
	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/domain/entity"
	"convercy/domain/valueobject"
)

type CurrencyRegistrationService struct {
	currenciesRepository           repositories.CurrenciesRepository
	registeredCurrenciesRepository repositories.RegisteredCurrenciesRepository
}

func NewCurrencyRegistrationService(
	currenciesRepository repositories.CurrenciesRepository,
	registeredCurrenciesRepository repositories.RegisteredCurrenciesRepository,
) *CurrencyRegistrationService {
	return &CurrencyRegistrationService{
		currenciesRepository:           currenciesRepository,
		registeredCurrenciesRepository: registeredCurrenciesRepository,
	}
}

// RegisterCurrency registers a new currency
func (s *CurrencyRegistrationService) RegisterCurrency(request dto.RegisterCurrencyRequest) (dto.RegisterCurrencyResponse, error) {
	code, err := valueobject.NewCurrencyCode(request.Code)
	if err != nil {
		return dto.RegisterCurrencyResponse{}, err
	}

	allCurrencyCodes, err := s.currenciesRepository.ListCurrencyCodes()
	if err != nil {
		return dto.RegisterCurrencyResponse{}, err
	}

	if !allCurrencyCodes.Contains(code) {
		return dto.RegisterCurrencyResponse{}, domain.ErrCurrencyCodeNotFound()
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

// ListRegisteredCurrencies lists all registered currencies
func (s *CurrencyRegistrationService) ListRegisteredCurrencies() (dto.ListRegisteredCurrenciesResponse, error) {
	registeredCurrencies, err := s.registeredCurrenciesRepository.GetRegisteredCurrencies()
	if err != nil && !errors.As(err, new(domain.ErrNotFound)) {
		return nil, err
	}

	return dto.BuildListRegisteredCurrenciesResponse(registeredCurrencies.Currencies()), nil
}
