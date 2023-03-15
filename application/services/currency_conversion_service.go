package services

import (
	"math"

	"convercy/application/dto"
	"convercy/application/repositories"
	"convercy/domain/usecases"
	"convercy/domain/valueobject"
)

type CurrencyConversionService struct {
	currencyCodeValidationService  usecases.CurrencyCodeValidationService
	currencyConversionService      usecases.CurrencyConversionService
	registeredCurrenciesRepository repositories.RegisteredCurrenciesRepository
	currencyExchangeRatesService   usecases.CurrencyExchangeRatesService
}

func NewCurrencyConversionService(
	currencyCodeValidationService usecases.CurrencyCodeValidationService,
	currencyConversionService usecases.CurrencyConversionService,
	registeredCurrenciesRepository repositories.RegisteredCurrenciesRepository,
	currencyExchangeRatesService usecases.CurrencyExchangeRatesService,
) *CurrencyConversionService {
	return &CurrencyConversionService{
		currencyCodeValidationService:  currencyCodeValidationService,
		currencyConversionService:      currencyConversionService,
		registeredCurrenciesRepository: registeredCurrenciesRepository,
		currencyExchangeRatesService:   currencyExchangeRatesService,
	}
}

func (s *CurrencyConversionService) ConvertCurrency(request dto.ConvertCurrencyRequest) (dto.ConvertCurrencyResponse, error) {
	amount, err := valueobject.NewCurrencyAmount(request.Amount)
	if err != nil {
		return nil, err
	}

	code, err := valueobject.NewCurrencyCode(request.Code)
	if err != nil {
		return nil, err
	}

	if err := s.currencyCodeValidationService.ValidateCurrencyCode(code); err != nil {
		return nil, err
	}

	registeredCurrencies, err := s.registeredCurrenciesRepository.GetRegisteredCurrencies()
	if err != nil {
		return nil, err
	}

	baseCurrency, err := registeredCurrencies.FindCurrencyByCode(code)
	if err != nil {
		return nil, err
	}

	exchangeRates, err := s.currencyExchangeRatesService.ListCurrencyExchangeRates(baseCurrency)
	if err != nil {
		return nil, err
	}

	response := make(dto.ConvertCurrencyResponse, len(exchangeRates))

	for _, exchangeRate := range exchangeRates {
		targetCurrencyCode := exchangeRate.Unit().TargetCurrencyCode()

		if targetCurrencyCode.Equal(code) {
			continue
		}

		if !registeredCurrencies.HasCurrencyWithCode(exchangeRate.Unit().TargetCurrencyCode()) {
			continue
		}

		convertedAmount, err := s.currencyConversionService.ConvertCurrency(amount, code, exchangeRate)
		if err != nil {
			return nil, err
		}

		response[exchangeRate.Unit().TargetCurrencyCode().String()] = math.Round(convertedAmount.Value()*100) / 100
	}

	return response, nil
}
