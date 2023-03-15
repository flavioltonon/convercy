package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

type ExchangeRateUnit struct {
	base   CurrencyCode
	target CurrencyCode
}

func NewExchangeRateUnit(base CurrencyCode, target CurrencyCode) (ExchangeRateUnit, error) {
	exchangeRateUnit := ExchangeRateUnit{
		base:   base,
		target: target,
	}

	if err := exchangeRateUnit.Validate(); err != nil {
		return ExchangeRateUnit{}, err
	}

	return exchangeRateUnit, nil
}

func (v ExchangeRateUnit) Equal(ref ExchangeRateUnit) bool {
	return v.base.Equal(ref.base) && v.target.Equal(ref.target)
}

func (v ExchangeRateUnit) BaseCurrencyCode() CurrencyCode {
	return v.base
}

func (v ExchangeRateUnit) TargetCurrencyCode() CurrencyCode {
	return v.target
}

func (v ExchangeRateUnit) Validate() error {
	return ozzo.ValidateStruct(&v,
		ozzo.Field(&v.base, ozzo.Required),
		ozzo.Field(&v.target, ozzo.Required),
	)
}
