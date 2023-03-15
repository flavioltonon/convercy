package valueobject

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type ExchangeRate struct {
	rate ExchangeRateValue
	unit ExchangeRateUnit
}

func NewExchangeRate(rate float64, baseCurrencyCode string, targetCurrencyCode string) (ExchangeRate, error) {
	r, err := NewExchangeRateValue(rate)
	if err != nil {
		return ExchangeRate{}, err
	}

	b, err := NewCurrencyCode(baseCurrencyCode)
	if err != nil {
		return ExchangeRate{}, err
	}

	t, err := NewCurrencyCode(targetCurrencyCode)
	if err != nil {
		return ExchangeRate{}, err
	}

	u, err := NewExchangeRateUnit(b, t)
	if err != nil {
		return ExchangeRate{}, err
	}

	exchangeRate := ExchangeRate{
		rate: r,
		unit: u,
	}

	if err := exchangeRate.Validate(); err != nil {
		return ExchangeRate{}, err
	}

	return exchangeRate, nil
}

func (v ExchangeRate) Equal(ref ExchangeRate) bool {
	return v.rate.Equal(ref.rate) && v.unit.Equal(ref.unit)
}

func (v ExchangeRate) Rate() ExchangeRateValue {
	return v.rate
}

func (v ExchangeRate) Unit() ExchangeRateUnit {
	return v.unit
}

func (v ExchangeRate) Validate() error {
	return ozzo.ValidateStruct(&v,
		ozzo.Field(&v.rate, ozzo.Required),
		ozzo.Field(&v.unit, ozzo.Required),
	)
}
