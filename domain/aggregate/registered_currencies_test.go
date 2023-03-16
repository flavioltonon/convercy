package aggregate

import (
	"testing"

	"convercy/domain/entity"
	"convercy/domain/valueobject"

	"github.com/stretchr/testify/suite"
)

type RegisteredCurrenciesTestSuite struct {
	suite.Suite

	clientID valueobject.ClientID

	brlCurrencyID   valueobject.CurrencyID
	brlCurrencyCode valueobject.CurrencyCode
	brlCurrency     *entity.Currency

	usdCurrencyID   valueobject.CurrencyID
	usdCurrencyCode valueobject.CurrencyCode
	usdCurrency     *entity.Currency

	currencies []*entity.Currency
}

func (s *RegisteredCurrenciesTestSuite) SetupSuite() {
	s.clientID = valueobject.GenerateClientID()

	s.brlCurrencyID = valueobject.GenerateCurrencyID()
	s.brlCurrencyCode, _ = valueobject.NewCurrencyCode("BRL")
	s.brlCurrency, _ = entity.NewCurrency(s.brlCurrencyID, s.brlCurrencyCode)
	s.usdCurrencyID = valueobject.GenerateCurrencyID()
	s.usdCurrencyCode, _ = valueobject.NewCurrencyCode("USD")
	s.usdCurrency, _ = entity.NewCurrency(s.usdCurrencyID, s.usdCurrencyCode)
	s.currencies = []*entity.Currency{s.brlCurrency}
}

func TestRegisteredCurrenciesTestSuite(t *testing.T) {
	suite.Run(t, new(RegisteredCurrenciesTestSuite))
}

func (s *RegisteredCurrenciesTestSuite) TestRegisteredCurrencies_RegisterCurrency() {
	type fields struct {
		clientID   valueobject.ClientID
		currencies []*entity.Currency
	}

	type args struct {
		currency *entity.Currency
	}

	tests := []struct {
		name           string
		fields         fields
		args           args
		wantCurrencies []*entity.Currency
		wantErr        bool
	}{
		{
			name: "When I try to add a currency that has not been registered, no errors should be returned",
			fields: fields{
				clientID:   s.clientID,
				currencies: []*entity.Currency{s.brlCurrency},
			},
			args: args{
				currency: s.usdCurrency,
			},
			wantCurrencies: []*entity.Currency{s.brlCurrency, s.usdCurrency},
			wantErr:        false,
		},
		{
			name: "When I try to add a currency that has already been registered, an error should be returned",
			fields: fields{
				clientID:   s.clientID,
				currencies: s.currencies,
			},
			args: args{
				currency: s.brlCurrency,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			a := &RegisteredCurrencies{
				clientID:   tt.fields.clientID,
				currencies: tt.fields.currencies,
			}

			if err := a.RegisterCurrency(tt.args.currency); tt.wantErr {
				s.Error(err)
			} else {
				s.ElementsMatch(a.currencies, tt.wantCurrencies)
				s.NoError(err)
			}
		})
	}
}

func (s *RegisteredCurrenciesTestSuite) TestRegisteredCurrencies_FindCurrencyByCode() {
	type fields struct {
		clientID   valueobject.ClientID
		currencies []*entity.Currency
	}

	type args struct {
		code valueobject.CurrencyCode
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Currency
		wantErr bool
	}{
		{
			name: "When I try to find a currency with a code that has already been registered, no errors should be returned",
			fields: fields{
				clientID:   s.clientID,
				currencies: s.currencies,
			},
			args: args{
				code: s.brlCurrencyCode,
			},
			want:    s.brlCurrency,
			wantErr: false,
		},
		{
			name: "When I try to find a currency with a code that has not been registered yet, an error should be returned",
			fields: fields{
				clientID:   s.clientID,
				currencies: s.currencies,
			},
			args: args{
				code: s.usdCurrencyCode,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			a := &RegisteredCurrencies{
				clientID:   tt.fields.clientID,
				currencies: tt.fields.currencies,
			}

			currency, err := a.FindCurrencyByCode(tt.args.code)
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
			}

			s.Equal(tt.want, currency)
		})
	}
}

func (s *RegisteredCurrenciesTestSuite) TestRegisteredCurrencies_HasCurrencyWithCode() {
	type fields struct {
		clientID   valueobject.ClientID
		currencies []*entity.Currency
	}

	type args struct {
		code valueobject.CurrencyCode
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "When I try to find a currency with a code that has already been registered, HasCurrencyWithCode should return true",
			fields: fields{
				clientID:   s.clientID,
				currencies: s.currencies,
			},
			args: args{
				code: s.brlCurrencyCode,
			},
			want: true,
		},
		{
			name: "When I try to find a currency with a code that has not been registered yet, HasCurrencyWithCode should return false",
			fields: fields{
				clientID:   s.clientID,
				currencies: s.currencies,
			},
			args: args{
				code: s.usdCurrencyCode,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			a := &RegisteredCurrencies{
				clientID:   tt.fields.clientID,
				currencies: tt.fields.currencies,
			}

			s.Equal(tt.want, a.HasCurrencyWithCode(tt.args.code))
		})
	}
}

func (s *RegisteredCurrenciesTestSuite) TestRegisteredCurrencies_UnregisterCurrency() {
	type fields struct {
		clientID   valueobject.ClientID
		currencies []*entity.Currency
	}

	type args struct {
		currencyID valueobject.CurrencyID
	}

	tests := []struct {
		name           string
		fields         fields
		args           args
		wantCurrencies []*entity.Currency
		wantErr        bool
	}{
		{
			name: "When I try to find a currency with an ID that does not exist, an error should be returned",
			fields: fields{
				clientID:   s.clientID,
				currencies: s.currencies,
			},
			args: args{
				currencyID: s.usdCurrencyID,
			},
			wantCurrencies: s.currencies,
			wantErr:        true,
		},
		{
			name: "When I try to find a currency with an ID that has been registered, no errors should be returned",
			fields: fields{
				clientID:   s.clientID,
				currencies: []*entity.Currency{s.brlCurrency},
			},
			args: args{
				currencyID: s.brlCurrencyID,
			},
			wantCurrencies: []*entity.Currency{},
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			a := &RegisteredCurrencies{
				clientID:   tt.fields.clientID,
				currencies: tt.fields.currencies,
			}

			if err := a.UnregisterCurrency(tt.args.currencyID); tt.wantErr {
				s.Error(err)
			} else {
				s.ElementsMatch(a.currencies, tt.wantCurrencies)
				s.NoError(err)
			}
		})
	}
}
