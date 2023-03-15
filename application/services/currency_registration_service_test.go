package services

import (
	"convercy/application/dto"
	"convercy/application/repositories"
	repositoriesMocks "convercy/application/repositories/mocks"
	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/domain/entity"
	"convercy/domain/usecases"
	domainUsecasesMocks "convercy/domain/usecases/mocks"
	"convercy/domain/valueobject"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CurrencyRegistrationServiceTestSuite struct {
	suite.Suite

	clientID valueobject.ClientID

	currencyID valueobject.CurrencyID

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

		invalid struct {
			currencyCode valueobject.CurrencyCode
		}
	}
}

func (s *CurrencyRegistrationServiceTestSuite) SetupSuite() {
	s.clientID = valueobject.GenerateClientID()

	// Currencies
	s.currencies.brl.currencyID = valueobject.GenerateCurrencyID()
	s.currencies.brl.currencyCode, _ = valueobject.NewCurrencyCode("BRL")
	s.currencies.brl.currency, _ = entity.NewCurrency(s.currencies.brl.currencyID, s.currencies.brl.currencyCode)
	s.currencies.usd.currencyID = valueobject.GenerateCurrencyID()
	s.currencies.usd.currencyCode, _ = valueobject.NewCurrencyCode("USD")
	s.currencies.usd.currency, _ = entity.NewCurrency(s.currencies.usd.currencyID, s.currencies.usd.currencyCode)
	s.currencies.invalid.currencyCode, _ = valueobject.NewCurrencyCode("FOO")
}

func TestCurrencyRegistrationServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyRegistrationServiceTestSuite))
}

func (s *CurrencyRegistrationServiceTestSuite) TestCurrencyRegistrationService_RegisterCurrency() {
	type fields struct {
		currencyCodeValidationService  func() usecases.CurrencyCodeValidationService
		registeredCurrenciesRepository func() repositories.RegisteredCurrenciesRepository
	}

	type args struct {
		request dto.RegisterCurrencyRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.RegisterCurrencyResponse
		wantErr bool
	}{
		{
			name: "When the input request.Code is not valid, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					return domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
				},
			},
			args: args{
				request: dto.RegisterCurrencyRequest{
					Code: "",
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
			},
			args: args{
				request: dto.RegisterCurrencyRequest{
					Code: "FOO",
				},
			},
			wantErr: true,
		},
		{
			name: "When the search for registered currencies fails (except for domain.ErrNotFound errors), an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.brl.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(nil, errors.New("some error"))
					return mock
				},
			},
			args: args{
				request: dto.RegisterCurrencyRequest{
					Code: "BRL",
				},
			},
			wantErr: true,
		},
		{
			name: "When the parsed currency code has already been registered, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.brl.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency), nil)
					return mock
				},
			},
			args: args{
				request: dto.RegisterCurrencyRequest{
					Code: "BRL",
				},
			},
			wantErr: true,
		},
		{
			name: "When the persistence of the updated registered currencies fail, an error should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.usd.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					registeredCurrencies := aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency)
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(registeredCurrencies, nil)
					mock.On("SaveRegisteredCurrencies", registeredCurrencies).Return(errors.New("some error"))
					return mock
				},
			},
			args: args{
				request: dto.RegisterCurrencyRequest{
					Code: "USD",
				},
			},
			wantErr: true,
		},
		{
			name: "When everything works as intended and everything is valid (WOW), no errors should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.usd.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					registeredCurrencies := aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency)
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(registeredCurrencies, nil)
					mock.On("SaveRegisteredCurrencies", registeredCurrencies).Return(nil)
					return mock
				},
			},
			args: args{
				request: dto.RegisterCurrencyRequest{
					Code: "USD",
				},
			},
			wantErr: false,
		},
		{
			name: "When no RegisteredCurrencies are found, an empty one should be created and no errors should be returned",
			fields: fields{
				currencyCodeValidationService: func() usecases.CurrencyCodeValidationService {
					mock := domainUsecasesMocks.NewCurrencyCodeValidationService(s.T())
					mock.On("ValidateCurrencyCode", s.currencies.usd.currencyCode).Return(nil)
					return mock
				},
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					repository := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					repository.On("GetRegisteredCurrencies").Return(nil, domain.ErrRegisteredCurrenciesNotFound())
					repository.On("SaveRegisteredCurrencies", mock.AnythingOfType("*aggregate.RegisteredCurrencies")).Return(nil)
					return repository
				},
			},
			args: args{
				request: dto.RegisterCurrencyRequest{
					Code: "USD",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			service := &CurrencyRegistrationService{
				currencyCodeValidator:          tt.fields.currencyCodeValidationService(),
				registeredCurrenciesRepository: tt.fields.registeredCurrenciesRepository(),
			}
			_, err := service.RegisterCurrency(tt.args.request)
			s.Equal(tt.wantErr, err != nil)
		})
	}
}

func (s *CurrencyRegistrationServiceTestSuite) TestCurrencyRegistrationService_UnregisterCurrency() {
	type fields struct {
		registeredCurrenciesRepository func() repositories.RegisteredCurrenciesRepository
	}

	type args struct {
		request dto.UnregisterCurrencyRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "When the input request.CurrencyID is not valid, an error should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					return repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
				},
			},
			args: args{
				request: dto.UnregisterCurrencyRequest{
					CurrencyID: "",
				},
			},
			wantErr: true,
		},
		{
			name: "When the search for registered currencies fails, an error should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(nil, domain.ErrRegisteredCurrenciesNotFound())
					return mock
				},
			},
			args: args{
				request: dto.UnregisterCurrencyRequest{
					CurrencyID: s.currencies.brl.currencyID.String(),
				},
			},
			wantErr: true,
		},
		{
			name: "When the currency ID provided is not related to any registered currencies, an error should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency), nil)
					return mock
				},
			},
			args: args{
				request: dto.UnregisterCurrencyRequest{
					CurrencyID: s.currencies.usd.currencyID.String(),
				},
			},
			wantErr: true,
		},
		{
			name: "When the persistence of the updated registered currencies fail, an error should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					registeredCurrencies := aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency)
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(registeredCurrencies, nil)
					mock.On("SaveRegisteredCurrencies", registeredCurrencies).Return(errors.New("some error"))
					return mock
				},
			},
			args: args{
				request: dto.UnregisterCurrencyRequest{
					CurrencyID: s.currencies.brl.currencyID.String(),
				},
			},
			wantErr: true,
		},
		{
			name: "When everything works as intended and everything is valid, no errors should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					registeredCurrencies := aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency)
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(registeredCurrencies, nil)
					mock.On("SaveRegisteredCurrencies", registeredCurrencies).Return(nil)
					return mock
				},
			},
			args: args{
				request: dto.UnregisterCurrencyRequest{
					CurrencyID: s.currencies.brl.currencyID.String(),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			service := &CurrencyRegistrationService{
				registeredCurrenciesRepository: tt.fields.registeredCurrenciesRepository(),
			}
			err := service.UnregisterCurrency(tt.args.request)
			s.Equal(tt.wantErr, err != nil)
		})
	}
}

