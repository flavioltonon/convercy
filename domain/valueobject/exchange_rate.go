package valueobject

import ozzo "github.com/go-ozzo/ozzo-validation/v4"

type ExchangeRate struct {
	rate ExchangeRateValue
	unit ExchangeRateUnit
}

func NewExchangeRate(rate float64, baseCurrencyCode string, targetCurrencyCode string) (ExchangeRate, error) {
	r, err := NewExchangeRateValue(rate)
	if err != nil {
		return ExchangeRate{}, err
	}

	u, err := NewExchangeRateUnit(baseCurrencyCode, targetCurrencyCode)
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

func (e ExchangeRate) Equal(ref ExchangeRate) bool {
	return e.rate.Equal(ref.rate) && e.unit.Equal(ref.unit)
}

func (e ExchangeRate) Rate() ExchangeRateValue {
	return e.rate
}

func (e ExchangeRate) Unit() ExchangeRateUnit {
	return e.unit
}

func (e ExchangeRate) Validate() error {
	return ozzo.ValidateStruct(&e,
		ozzo.Field(&e.rate, ozzo.Required),
		ozzo.Field(&e.unit, ozzo.Required),
	)
}

// Inverse returns the inverse version of an exchange rate (e.g. BRL/USD -> USD/BRL)
func (e ExchangeRate) Inverse() ExchangeRate {
	return ExchangeRate{
		rate: ExchangeRateValue{
			value: 1 / e.rate.value,
		},
		unit: ExchangeRateUnit{
			base:   e.unit.target,
			target: e.unit.base,
		},
	}
}
