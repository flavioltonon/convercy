package services

import (
	"errors"
	"testing"

	"convercy/domain/entity"
	"convercy/domain/usecases"
	domainUsecasesMocks "convercy/domain/usecases/mocks"
	"convercy/domain/valueobject"

	"github.com/stretchr/testify/suite"
)

type CurrencyExchangeRatesServiceTestSuite struct {
	suite.Suite

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
	}

	exchangeRates struct {
		brlUsd valueobject.ExchangeRate
		brlGbp valueobject.ExchangeRate
		brlBrl valueobject.ExchangeRate
		usdBrl valueobject.ExchangeRate
		usdGbp valueobject.ExchangeRate
		usdUsd valueobject.ExchangeRate
		gbpBrl valueobject.ExchangeRate
		gbpUsd valueobject.ExchangeRate
		gbpGbp valueobject.ExchangeRate
	}
}

func (s *CurrencyExchangeRatesServiceTestSuite) SetupSuite() {
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

	// Exchange rates
	s.exchangeRates.brlUsd, _ = valueobject.NewExchangeRate(0.193603, "BRL", "USD")
	s.exchangeRates.brlBrl, _ = valueobject.NewExchangeRate(1, "BRL", "BRL")
	s.exchangeRates.brlGbp, _ = valueobject.NewExchangeRate(0.1006901537373359, "BRL", "GBP")
	s.exchangeRates.usdBrl, _ = valueobject.NewExchangeRate(5.1652, "USD", "BRL")
	s.exchangeRates.usdGbp, _ = valueobject.NewExchangeRate(0.838732, "USD", "GBP")
	s.exchangeRates.usdUsd, _ = valueobject.NewExchangeRate(1, "USD", "USD")
	s.exchangeRates.gbpBrl, _ = valueobject.NewExchangeRate(9.931458, "GBP", "BRL")
	s.exchangeRates.gbpUsd, _ = valueobject.NewExchangeRate(1.92276, "GBP", "USD")
	s.exchangeRates.gbpGbp, _ = valueobject.NewExchangeRate(1, "GBP", "GBP")
}

func TestCurrencyExchangeRatesServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyExchangeRatesServiceTestSuite))
}

func (s *CurrencyExchangeRatesServiceTestSuite) TestCurrencyExchangeRatesService_ListCurrencyExchangeRates() {
	type fields struct {
		core func() usecases.ExchangeRatesService
	}

	type args struct {
		currency *entity.Currency
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    valueobject.ExchangeRates
		wantErr bool
	}{
		{
			name: "When the input currency is not valid, an error should be returned",
			fields: fields{
				core: func() usecases.ExchangeRatesService {
					return domainUsecasesMocks.NewExchangeRatesService(s.T())
				},
			},
			args: args{
				currency: &entity.Currency{},
			},
			wantErr: true,
		},
		{
			name: "When ExchangeRatesService.ListExchangeRates fails to list exchange rates, an error should be returned",
			fields: fields{
				core: func() usecases.ExchangeRatesService {
					mock := domainUsecasesMocks.NewExchangeRatesService(s.T())
					mock.On("ListExchangeRates").Return(nil, errors.New("some error"))
					return mock
				},
			},
			args: args{
				currency: s.currencies.brl.currency,
			},
			wantErr: true,
		},
		{
			name: "When ExchangeRatesService.ListExchangeRates fails to list exchange rates, an error should be returned",
			fields: fields{
				core: func() usecases.ExchangeRatesService {
					mock := domainUsecasesMocks.NewExchangeRatesService(s.T())
					mock.On("ListExchangeRates").Return(nil, errors.New("some error"))
					return mock
				},
			},
			args: args{
				currency: s.currencies.brl.currency,
			},
			wantErr: true,
		},
		{
			name: "When the exchange rates list does not contain any rates with a base code as the input currency, an error should be returned",
			fields: fields{
				core: func() usecases.ExchangeRatesService {
					mock := domainUsecasesMocks.NewExchangeRatesService(s.T())
					mock.On("ListExchangeRates").Return(valueobject.ExchangeRates{s.exchangeRates.brlUsd}, nil)
					return mock
				},
			},
			args: args{
				currency: s.currencies.usd.currency,
			},
			wantErr: true,
		},
		{
			name: "When the exchange rates list contain a rate with a target currency different than the others, an error should be returned",
			fields: fields{
				core: func() usecases.ExchangeRatesService {
					mock := domainUsecasesMocks.NewExchangeRatesService(s.T())
					mock.On("ListExchangeRates").Return(valueobject.ExchangeRates{s.exchangeRates.brlUsd, s.exchangeRates.usdBrl}, nil)
					return mock
				},
			},
			args: args{
				currency: s.currencies.brl.currency,
			},
			wantErr: true,
		},
		{
			name: "When everything works as intended and everything is valid, a list of exchange rates based on the input currency should be returned",
			fields: fields{
				core: func() usecases.ExchangeRatesService {
					mock := domainUsecasesMocks.NewExchangeRatesService(s.T())
					mock.On("ListExchangeRates").Return(valueobject.ExchangeRates{s.exchangeRates.brlUsd, s.exchangeRates.gbpUsd, s.exchangeRates.usdUsd}, nil)
					return mock
				},
			},
			args: args{
				currency: s.currencies.brl.currency,
			},
			want: valueobject.ExchangeRates{
				s.exchangeRates.brlBrl,
				s.exchangeRates.brlUsd,
				s.exchangeRates.brlGbp,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			service := &CurrencyExchangeRatesService{
				core: tt.fields.core(),
			}
			if got, err := service.ListCurrencyExchangeRates(tt.args.currency); tt.wantErr {
				s.Error(err)
			} else {
				s.ElementsMatch(tt.want, got)
				s.NoError(err)
			}
		})
	}
}
