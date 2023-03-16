package services

import (
	"testing"

	"convercy/application/dto"
	"convercy/application/repositories"
	repositoriesMocks "convercy/application/repositories/mocks"
	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/domain/entity"
	"convercy/domain/services"
	"convercy/domain/usecases"
	domainUsecasesMocks "convercy/domain/usecases/mocks"
	"convercy/domain/valueobject"

	"github.com/stretchr/testify/suite"
)

type CurrencyConversionServiceTestSuite struct {
	suite.Suite

	clientID valueobject.ClientID

	currencies struct {
		brl struct {
			currencyID   valueobject.CurrencyID
			currencyCode valueobject.CurrencyCode
			currency     *entity.Currency
		}

		usd struct {
			currencyID   valueobject.CurrencyID
			currencyCode valueobject.CurrencyCode
			currency     *entity.Currency
		}

		gbp struct {
			currencyID   valueobject.CurrencyID
			currencyCode valueobject.CurrencyCode
			currency     *entity.Currency
		}

		invalid struct {
			currencyCode valueobject.CurrencyCode
		}
	}

	exchangeRates struct {
		brlUsd valueobject.ExchangeRate
		brlGbp valueobject.ExchangeRate
		usdBrl valueobject.ExchangeRate
		usdGbp valueobject.ExchangeRate
		gbpBrl valueobject.ExchangeRate
		gbpUsd valueobject.ExchangeRate
	}

	registeredCurrencies *aggregate.RegisteredCurrencies
}

func (s *CurrencyConversionServiceTestSuite) SetupSuite() {
	s.clientID = valueobject.GenerateClientID()

	// Currencies
	s.currencies.brl.currencyID = valueobject.GenerateCurrencyID()
	s.currencies.brl.currencyCode, _ = valueobject.NewCurrencyCode("BRL")
	s.currencies.brl.currency, _ = entity.NewCurrency(s.currencies.brl.currencyID, s.currencies.brl.currencyCode)
	s.currencies.usd.currencyID = valueobject.GenerateCurrencyID()
	s.currencies.usd.currencyCode, _ = valueobject.NewCurrencyCode("USD")
	s.currencies.usd.currency, _ = entity.NewCurrency(s.currencies.usd.currencyID, s.currencies.usd.currencyCode)
	s.currencies.gbp.currencyID = valueobject.GenerateCurrencyID()
	s.currencies.gbp.currencyCode, _ = valueobject.NewCurrencyCode("GBP")
	s.currencies.gbp.currency, _ = entity.NewCurrency(s.currencies.gbp.currencyID, s.currencies.gbp.currencyCode)
	s.currencies.invalid.currencyCode, _ = valueobject.NewCurrencyCode("FOO")

	// Exchange rates
	s.exchangeRates.brlUsd, _ = valueobject.NewExchangeRate(0.193603, "BRL", "USD")
	s.exchangeRates.brlGbp, _ = valueobject.NewExchangeRate(0.10069, "BRL", "GBP")
	s.exchangeRates.usdBrl, _ = valueobject.NewExchangeRate(5.1652, "USD", "BRL")
	s.exchangeRates.usdGbp, _ = valueobject.NewExchangeRate(0.838732, "USD", "GBP")
	s.exchangeRates.gbpBrl, _ = valueobject.NewExchangeRate(9.931458, "GBP", "BRL")
	s.exchangeRates.gbpUsd, _ = valueobject.NewExchangeRate(1.92276, "GBP", "USD")

	s.registeredCurrencies = aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency, s.currencies.usd.currency)
}

func TestCurrencyConversionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyConversionServiceTestSuite))
}

func (s *CurrencyConversionServiceTestSuite) TestCurrencyConversionService_ConvertCurrency() {
	type fields struct {
		currencyCodeValidationService  func() usecases.CurrencyCodeValidationService
		currencyConversionService      func() usecases.CurrencyConversionService
		registeredCurrenciesRepository func() repositories.RegisteredCurrenciesRepository
		currencyExchangeRatesService   func() usecases.CurrencyExchangeRatesService
	}

	type args struct {
		request dto.ConvertCurrencyRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.ConvertCurrencyResponse
		wantErr bool
	}{
		{
			name: "When the input request.Amount is not valid, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					return domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					return domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: -1,
				},
			},
			wantErr: true,
		},
		{
			name: "When the input request.Code is not valid, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					return domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					return domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: 10,
					Code:   "",
				},
			},
			wantErr: true,
		},
		{
			name: "When the parsed CurrencyCode is not valid, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.invalid.currencyCode).Return(domain.ErrCurrencyCodeNotFound())
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					return domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: 10,
					Code:   "FOO",
				},
			},
			wantErr: true,
		},
		{
			name: "When no currencies have been registered yet, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.brl.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(nil, domain.ErrRegisteredCurrenciesNotFound())
					return mock
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					return domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: 10,
					Code:   "BRL",
				},
			},
			wantErr: true,
		},
		{
			name: "When the parsed currency code has not been registered yet, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.gbp.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					return domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: 10,
					Code:   "GBP",
				},
			},
			wantErr: true,
		},
		{
			name: "When CurrencyExchangeRatesService.ListCurrencyExchangeRates fails to list exchange rates for a currency, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.brl.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					mock := domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
					mock.On("ListCurrencyExchangeRates", s.currencies.brl.currency).Return(nil, domain.ErrExchangeRateNotFound())
					return mock
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: 10,
					Code:   "BRL",
				},
			},
			wantErr: true,
		},
		{
			name: "When CurrencyConversionService.ConvertCurrency fails to convert an amount of currency, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.brl.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					mock := domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
					mock.On("ListCurrencyExchangeRates", s.currencies.brl.currency).Return(nil, domain.ErrExchangeRateNotFound())
					return mock
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: 10,
					Code:   "BRL",
				},
			},
			wantErr: true,
		},
		{
			name: `When everything works as intended and everything is valid (WOW), a ConvertCurrencyResponse should be returned with a set of 
			exchange rates for the input currency`,
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.brl.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
				},
				currencyExchangeRatesService: func() usecases.CurrencyExchangeRatesService {
					mock := domainUsecasesMocks.NewCurrencyExchangeRatesService(s.T())
					mock.On("ListCurrencyExchangeRates", s.currencies.brl.currency).Return(valueobject.ExchangeRates{
						s.exchangeRates.brlUsd,
					}, nil)
					return mock
				},
				currencyConversionService: func() usecases.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
			},
			args: args{
				request: dto.ConvertCurrencyRequest{
					Amount: 10,
					Code:   "BRL",
				},
			},
			want: dto.ConvertCurrencyResponse{
				"USD": 1.94,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			service := &CurrencyConversionService{
				currencyCodeValidationService:  tt.fields.currencyCodeValidationService(),
				currencyConversionService:      tt.fields.currencyConversionService(),
				registeredCurrenciesRepository: tt.fields.registeredCurrenciesRepository(),
				currencyExchangeRatesService:   tt.fields.currencyExchangeRatesService(),
			}
			got, err := service.ConvertCurrency(tt.args.request)
			s.Equal(tt.want, got)
			s.Equal(tt.wantErr, err != nil)
		})
	}
}
