package services

import (
	"convercy/domain"
	"convercy/domain/usecases"
	"convercy/domain/valueobject"
)

type CurrencyCodeValidationService struct {
	currenciesService usecases.CurrenciesService
}

func NewCurrencyCodeValidationService(currenciesService usecases.CurrenciesService) *CurrencyCodeValidationService {
	return &CurrencyCodeValidationService{currenciesService: currenciesService}
}

func (s *CurrencyCodeValidationService) ValidateCurrencyCode(code valueobject.CurrencyCode) error {
	if err := code.Validate(); err != nil {
		return err
	}

	codes, err := s.currenciesService.ListCurrencyCodes()
	if err != nil {
		return err
	}

	if !codes.Contains(code) {
		return domain.ErrInvalidCurrencyCode(domain.ErrCurrencyCodeNotFound())
	}

	return nil
}
