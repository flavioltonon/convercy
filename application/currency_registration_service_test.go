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
	"convercy/domain/valueobject"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CurrencyRegistrationServiceTestSuite struct {
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

		invalid struct {
			currencyCode valueobject.CurrencyCode
		}
	}

	registeredCurrencies *aggregate.RegisteredCurrencies
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

	s.registeredCurrencies = aggregate.NewRegisteredCurrencies(s.clientID, s.currencies.brl.currency, s.currencies.usd.currency)
}

func TestCurrencyRegistrationServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyRegistrationServiceTestSuite))
}

func (s *CurrencyRegistrationServiceTestSuite) TestCurrencyRegistrationService_RegisterCurrency() {
	type fields struct {
		currenciesRepository           func() repositories.CurrenciesRepository
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
				currenciesRepository: func() repositories.CurrenciesRepository {
					return repositoriesMocks.NewCurrenciesRepository(s.T())
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
			name: "When the CurrenciesRepository fails to fetch currencies, an error should be returned",
			fields: fields{
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(nil, errors.New("some error"))
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
			name: "When the parsed CurrencyCode is not valid, an error should be returned",
			fields: fields{
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
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
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
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
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
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
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.usd.currencyCode}, nil)
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
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.usd.currencyCode}, nil)
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
				currenciesRepository: func() repositories.CurrenciesRepository {
					mock := repositoriesMocks.NewCurrenciesRepository(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.usd.currencyCode}, nil)
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
				currenciesRepository:           tt.fields.currenciesRepository(),
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

func (s *CurrencyRegistrationServiceTestSuite) TestCurrencyRegistrationService_ListRegisteredCurrencies() {
	type fields struct {
		registeredCurrenciesRepository func() repositories.RegisteredCurrenciesRepository
	}

	tests := []struct {
		name    string
		fields  fields
		want    dto.ListRegisteredCurrenciesResponse
		wantErr bool
	}{
		{
			name: "When the search for registered currencies fails (except for domain.ErrNotFound errors), an error should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(nil, errors.New("some error"))
					return mock
				},
			},
			wantErr: true,
		},
		{
			name: "When everything works as intended and everything is valid, no errors should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(s.registeredCurrencies, nil)
					return mock
				},
			},
			want: dto.ListRegisteredCurrenciesResponse{
				{
					ID:   s.currencies.brl.currencyID.String(),
					Code: s.currencies.brl.currencyCode.String(),
				},
				{
					ID:   s.currencies.usd.currencyID.String(),
					Code: s.currencies.usd.currencyCode.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "When the search for registered currencies fails with a domain.ErrRegisteredCurrenciesNotFound, no errors should be returned",
			fields: fields{
				registeredCurrenciesRepository: func() repositories.RegisteredCurrenciesRepository {
					mock := repositoriesMocks.NewRegisteredCurrenciesRepository(s.T())
					mock.On("GetRegisteredCurrencies").Return(nil, domain.ErrRegisteredCurrenciesNotFound())
					return mock
				},
			},
			want:    dto.ListRegisteredCurrenciesResponse{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			service := &CurrencyRegistrationService{
				registeredCurrenciesRepository: tt.fields.registeredCurrenciesRepository(),
			}
			got, err := service.ListRegisteredCurrencies()
			s.ElementsMatch(tt.want, got)
			s.Equal(tt.wantErr, err != nil)
		})
	}
}
