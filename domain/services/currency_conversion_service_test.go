package services

import (
	"testing"

	"convercy/domain/valueobject"

	"github.com/stretchr/testify/suite"
)

type CurrencyConversionServiceTestSuite struct {
	suite.Suite

	currencies struct {
		amount          valueobject.CurrencyAmount
		convertedAmount valueobject.CurrencyAmount

		brl struct {
			currencyCode valueobject.CurrencyCode
		}

		usd struct {
			currencyCode valueobject.CurrencyCode
		}
	}

	exchangeRates struct {
		brlUsd valueobject.ExchangeRate
		usdBrl valueobject.ExchangeRate
	}
}

func (s *CurrencyConversionServiceTestSuite) SetupSuite() {
	// Currencies
	s.currencies.amount, _ = valueobject.NewCurrencyAmount(10)
	s.currencies.convertedAmount, _ = valueobject.NewCurrencyAmount(1.93603)
	s.currencies.brl.currencyCode, _ = valueobject.NewCurrencyCode("BRL")
	s.currencies.usd.currencyCode, _ = valueobject.NewCurrencyCode("USD")

	// Exchange rates
	s.exchangeRates.brlUsd, _ = valueobject.NewExchangeRate(0.193603, "BRL", "USD")
	s.exchangeRates.usdBrl, _ = valueobject.NewExchangeRate(5.1652, "USD", "BRL")

}

func TestCurrencyConversionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyConversionServiceTestSuite))
}

func (s *CurrencyConversionServiceTestSuite) TestCurrencyConversionService_ConvertCurrency() {
	type args struct {
		amount       valueobject.CurrencyAmount
		code         valueobject.CurrencyCode
		exchangeRate valueobject.ExchangeRate
	}

	tests := []struct {
		name    string
		args    args
		want    valueobject.CurrencyAmount
		wantErr bool
	}{
		{
			name: "When the input amount is not valid, an error should be returned",
			args: args{
				amount: valueobject.CurrencyAmount{},
			},
			wantErr: true,
		},
		{
			name: "When the input code is not valid, an error should be returned",
			args: args{
				amount: s.currencies.amount,
				code:   valueobject.CurrencyCode{},
			},
			wantErr: true,
		},
		{
			name: "When the input exchange rate is not valid, an error should be returned",
			args: args{
				amount:       s.currencies.amount,
				code:         s.currencies.brl.currencyCode,
				exchangeRate: valueobject.ExchangeRate{},
			},
			wantErr: true,
		},
		{
			name: "When the base currency code of the exchange rate is not the same as the input code, an error should be returned",
			args: args{
				amount:       s.currencies.amount,
				code:         s.currencies.brl.currencyCode,
				exchangeRate: s.exchangeRates.usdBrl,
			},
			wantErr: true,
		},
		{
			name: "When everything is OK, a new currency amount should be returned",
			args: args{
				amount:       s.currencies.amount,
				code:         s.currencies.brl.currencyCode,
				exchangeRate: s.exchangeRates.brlUsd,
			},
			want: s.currencies.convertedAmount,
		},
	}

	service := &CurrencyConversionService{}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			if got, err := service.ConvertCurrency(tt.args.amount, tt.args.code, tt.args.exchangeRate); tt.wantErr {
				s.Error(err)
			} else {
				s.Equal(tt.want, got)
				s.NoError(err)
			}
		})
	}
}
