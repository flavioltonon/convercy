package services

import (
	"errors"
	"testing"

	"convercy/domain/usecases"
	domainUsecasesMocks "convercy/domain/usecases/mocks"
	"convercy/domain/valueobject"

	"github.com/stretchr/testify/suite"
)

type CurrencyCodeValidationServiceTestSuite struct {
	suite.Suite

	currencies struct {
		brl struct {
			currencyCode valueobject.CurrencyCode
		}

		usd struct {
			currencyCode valueobject.CurrencyCode
		}
	}
}

func (s *CurrencyCodeValidationServiceTestSuite) SetupSuite() {
	s.currencies.brl.currencyCode, _ = valueobject.NewCurrencyCode("BRL")
	s.currencies.usd.currencyCode, _ = valueobject.NewCurrencyCode("USD")
}

func TestCurrencyCodeValidationServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyCodeValidationServiceTestSuite))
}

func (s *CurrencyCodeValidationServiceTestSuite) TestCurrencyCodeValidationService_ValidateCurrencyCode() {
	type fields struct {
		currenciesService func() usecases.CurrenciesService
	}

	type args struct {
		code valueobject.CurrencyCode
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "When the input CurrencyCode is invalid, an error should be returned",
			fields: fields{
				currenciesService: func() usecases.CurrenciesService {
					return domainUsecasesMocks.NewCurrenciesService(s.T())
				},
			},
			args: args{
				code: valueobject.CurrencyCode{},
			},
			wantErr: true,
		},
		{
			name: "When CurrenciesService.ListCurrencyCodes fails, an error should be returned",
			fields: fields{
				currenciesService: func() usecases.CurrenciesService {
					mock := domainUsecasesMocks.NewCurrenciesService(s.T())
					mock.On("ListCurrencyCodes").Return(nil, errors.New("some error"))
					return mock
				},
			},
			args: args{
				code: s.currencies.brl.currencyCode,
			},
			wantErr: true,
		},
		{
			name: "When the list of codes does not contain the input code, an error should be returned",
			fields: fields{
				currenciesService: func() usecases.CurrenciesService {
					mock := domainUsecasesMocks.NewCurrenciesService(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
			},
			args: args{
				code: s.currencies.usd.currencyCode,
			},
			wantErr: true,
		},
		{
			name: "When the list of codes do contain the input code, no errors should be returned",
			fields: fields{
				currenciesService: func() usecases.CurrenciesService {
					mock := domainUsecasesMocks.NewCurrenciesService(s.T())
					mock.On("ListCurrencyCodes").Return(valueobject.CurrencyCodes{s.currencies.brl.currencyCode}, nil)
					return mock
				},
			},
			args: args{
				code: s.currencies.brl.currencyCode,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			service := &CurrencyCodeValidationService{
				currenciesService: tt.fields.currenciesService(),
			}

			if err := service.ValidateCurrencyCode(tt.args.code); tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
			}
		})
	}
}
