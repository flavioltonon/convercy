package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

type ExchangeRateUnit struct {
	base   CurrencyCode
	target CurrencyCode
}

func NewExchangeRateUnit(baseCurrencyCode string, targetCurrencyCode string) (ExchangeRateUnit, error) {
	b, err := NewCurrencyCode(baseCurrencyCode)
	if err != nil {
		return ExchangeRateUnit{}, err
	}

	t, err := NewCurrencyCode(targetCurrencyCode)
	if err != nil {
		return ExchangeRateUnit{}, err
	}

	exchangeRateUnit := ExchangeRateUnit{
		base:   b,
		target: t,
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
