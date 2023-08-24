package application

import (
	"errors"
	"testing"

	"convercy/application/dto"
	"convercy/application/repositories"
	repositoriesMocks "convercy/application/repositories/mocks"
	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/domain/entity"
	"convercy/domain/services"
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
		usdUsd valueobject.ExchangeRate
		gbpBrl valueobject.ExchangeRate
		gbpUsd valueobject.ExchangeRate
	}

	currencyExchangeRates struct {
		usd *aggregate.CurrencyExchangeRates
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
	s.exchangeRates.usdUsd, _ = valueobject.NewExchangeRate(1, "USD", "USD")
	s.exchangeRates.gbpBrl, _ = valueobject.NewExchangeRate(9.931458, "GBP", "BRL")
	s.exchangeRates.gbpUsd, _ = valueobject.NewExchangeRate(1.92276, "GBP", "USD")

	s.currencyExchangeRates.usd, _ = aggregate.NewCurrencyExchangeRates(
		s.currencies.usd.currencyCode,
		s.exchangeRates.usdBrl, s.exchangeRates.usdUsd)

	s.registeredCurrencies = aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency, s.currencies.usd.currency)
}

func TestCurrencyConversionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyConversionServiceTestSuite))
}

func (s *CurrencyConversionServiceTestSuite) TestCurrencyConversionService_ConvertCurrency() {
	type fields struct {
		baselineCurrencyCode            func() valueobject.CurrencyCode
		currenciesRepository            func() repositories.CurrenciesRepository
		currencyConversionService       func() *services.CurrencyConversionService
		currencyExchangeRatesCache      func() repositories.CurrencyExchangeRatesCache
		currencyExchangeRatesRepository func() repositories.CurrencyExchangeRatesRepository
		registeredCurrenciesRepository  func() repositories.RegisteredCurrenciesRepository
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
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					return repositoriesMocks.NewCurrenciesRepository(s.T())
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					return repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					return repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
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
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					return repositoriesMocks.NewCurrenciesRepository(s.T())
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					return repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					return repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
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
			name: "When the CurrenciesRepository fails to fetch currencies, an error should be returned",
			fields: fields{
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(nil, errors.New("some error"))
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					return repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					return repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
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
			name: "When the parsed CurrencyCode is not valid, an error should be returned",
			fields: fields{
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					return repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					return repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
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
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					return repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					return repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(nil, domain.ErrRegisteredCurrenciesNotFound())
					return mock
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
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.gbp.currencyCode}, nil)
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					return repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					return repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
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
			name: "When ExchangeRatesService.ListExchangeRates fails to list exchange rates for a currency, an error should be returned",
			fields: fields{
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					mock := repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
					mock.On("GetCurrencyExchangeRates", s.currencies.usd.currencyCode).Return(nil, errors.New("some error"))
					return mock
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					mock := repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
					mock.On("GetCurrencyExchangeRates", s.currencies.usd.currencyCode).Return(nil, domain.ErrExchangeRateNotFound())
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
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
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					mock := repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
					mock.On("GetCurrencyExchangeRates", s.currencies.usd.currencyCode).Return(nil, domain.ErrCurrencyExchangeRatesNotFound())
					return mock
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					mock := repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
					mock.On("GetCurrencyExchangeRates", s.currencies.usd.currencyCode).Return(nil, domain.ErrExchangeRateNotFound())
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
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
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					m := repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
					m.On("GetCurrencyExchangeRates", s.currencies.usd.currencyCode).Return(nil, domain.ErrCurrencyExchangeRatesNotFound()).Once()
					m.On("SaveCurrencyExchangeRates", s.currencyExchangeRates.usd).Return(nil).Once()
					return m
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					mock := repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
					mock.On("GetCurrencyExchangeRates", s.currencies.usd.currencyCode).Return(s.currencyExchangeRates.usd, nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
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
		{
			name: `When everything works as intended and everything is valid after a CurrencyExchangeRatesCache cache hit, a ConvertCurrencyResponse should be returned
				with a set of exchange rates for the input currency`,
			fields: fields{
				baselineCurrencyCode: func() valueobject.CurrencyCode {
					return s.currencies.usd.currencyCode
				},
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
				currencyConversionService: func() *services.CurrencyConversionService {
					return services.NewCurrencyConversionService()
				},
				currencyExchangeRatesCache: func() repositories.CurrencyExchangeRatesCache {
					m := repositoriesMocks.NewCurrencyExchangeRatesCache(s.T())
					m.On("GetCurrencyExchangeRates", s.currencies.usd.currencyCode).Return(s.currencyExchangeRates.usd, nil)
					return m
				},
				currencyExchangeRatesRepository: func() repositories.CurrencyExchangeRatesRepository {
					mock := repositoriesMocks.NewCurrencyExchangeRatesRepository(s.T())
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
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
				baselineCurrencyCode:            tt.fields.baselineCurrencyCode(),
				currenciesRepository:            tt.fields.currenciesRepository(),
				currencyConversionService:       tt.fields.currencyConversionService(),
				currencyExchangeRatesCache:      tt.fields.currencyExchangeRatesCache(),
				currencyExchangeRatesRepository: tt.fields.currencyExchangeRatesRepository(),
				registeredCurrenciesRepository:  tt.fields.registeredCurrenciesRepository(),
			}
			if got, err := service.ConvertCurrency(tt.args.request); tt.wantErr {
				s.Error(err)
			} else {
				s.Equal(tt.want, got)
				s.NoError(err)
			}
		})
	}
}
