package services

import (
	"math"

	"convercy/application/dto"
	"convercy/application/repositories"
	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/domain/entity"
	"convercy/domain/services"
	"convercy/domain/valueobject"
)

type CurrencyConversionService struct {
	baselineCurrencyCode            valueobject.CurrencyCode
	currenciesRepository            repositories.CurrenciesRepository
	currencyConversionService       *services.CurrencyConversionService
	currencyExchangeRatesCache      repositories.CurrencyExchangeRatesCache
	currencyExchangeRatesRepository repositories.CurrencyExchangeRatesRepository
	registeredCurrenciesRepository  repositories.RegisteredCurrenciesRepository
}

func NewCurrencyConversionService(
	baselineCurrencyCode valueobject.CurrencyCode,
	currenciesRepository repositories.CurrenciesRepository,
	currencyConversionService *services.CurrencyConversionService,
	currencyExchangeRatesCache repositories.CurrencyExchangeRatesCache,
	currencyExchangeRatesRepository repositories.CurrencyExchangeRatesRepository,
	registeredCurrenciesRepository repositories.RegisteredCurrenciesRepository,
) *CurrencyConversionService {
	return &CurrencyConversionService{
		baselineCurrencyCode:            baselineCurrencyCode,
		currenciesRepository:            currenciesRepository,
		currencyConversionService:       currencyConversionService,
		currencyExchangeRatesCache:      currencyExchangeRatesCache,
		currencyExchangeRatesRepository: currencyExchangeRatesRepository,
		registeredCurrenciesRepository:  registeredCurrenciesRepository,
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

	allCurrencyCodes, err := s.currenciesRepository.ListCurrencyCodes()
	if err != nil {
		return nil, err
	}

	if !allCurrencyCodes.Contains(code) {
		return nil, domain.ErrCurrencyCodeNotFound()
	}

	registeredCurrencies, err := s.registeredCurrenciesRepository.GetRegisteredCurrencies()
	if err != nil {
		return nil, err
	}

	baseCurrency, err := registeredCurrencies.FindCurrencyByCode(code)
	if err != nil {
		return nil, err
	}

	currencyExchangeRates, err := s.getCurrencyExchangeRates(baseCurrency)
	if err != nil {
		return nil, err
	}

	response := make(dto.ConvertCurrencyResponse, len(currencyExchangeRates.ExchangeRates()))

	for _, exchangeRate := range currencyExchangeRates.ExchangeRates() {
		baselineCurrencyCode := exchangeRate.Unit().TargetCurrencyCode()

		if baselineCurrencyCode.Equal(code) {
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

func (s *CurrencyConversionService) getBaselineExchangeRates() (valueobject.ExchangeRates, error) {
	if currencyExchangeRates, err := s.currencyExchangeRatesCache.GetCurrencyExchangeRates(s.baselineCurrencyCode); err == nil {
		return currencyExchangeRates.ExchangeRates(), nil
	}

	currencyExchangeRates, err := s.currencyExchangeRatesRepository.GetCurrencyExchangeRates(s.baselineCurrencyCode)
	if err != nil {
		return nil, err
	}

	if err := s.currencyExchangeRatesCache.SaveCurrencyExchangeRates(currencyExchangeRates); err != nil {
		return nil, err
	}

	return currencyExchangeRates.ExchangeRates(), nil
}

func (s *CurrencyConversionService) getCurrencyExchangeRates(currency *entity.Currency) (*aggregate.CurrencyExchangeRates, error) {
	// Get all baseline exchange rates, which are exchange rates of all currencies in terms of the service's baseline currency code (e.g. USD)
	baselineExchangeRates, err := s.getBaselineExchangeRates()
	if err != nil {
		return nil, err
	}

	// inverseCurrencyExchangeRate is the value of the baseline currency in terms of the input currency
	inverseCurrencyExchangeRate, err := baselineExchangeRates.FindByTargetCurrencyCode(currency.Code())
	if err != nil {
		return nil, err
	}

	// currencyExchangeRate is the value of the input currency in terms of the baseline currency
	currencyExchangeRate := inverseCurrencyExchangeRate.Inverse()

	relativeExchangeRates := make(valueobject.ExchangeRates, 0, len(baselineExchangeRates))

	for _, targetExchangeRate := range baselineExchangeRates {
		if !currencyExchangeRate.Unit().TargetCurrencyCode().Equal(targetExchangeRate.Unit().BaseCurrencyCode()) {
			return nil, domain.ErrIncompatibleExchangeRates()
		}

		relativeExchangeRate, err := valueobject.NewExchangeRate(
			currencyExchangeRate.Rate().Value()*targetExchangeRate.Rate().Value(),
			currencyExchangeRate.Unit().BaseCurrencyCode().String(),
			targetExchangeRate.Unit().TargetCurrencyCode().String(),
		)
		if err != nil {
			return nil, err
		}

		relativeExchangeRates = append(relativeExchangeRates, relativeExchangeRate)
	}

	currencyExchangeRates, err := aggregate.NewCurrencyExchangeRates(currency.Code(), relativeExchangeRates...)
	if err != nil {
		return nil, err
	}

	return currencyExchangeRates, nil
}
