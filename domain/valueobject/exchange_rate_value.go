package valueobject

import (
	"math"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type ExchangeRateValue struct {
	value float64
}

func NewExchangeRateValue(value float64) (ExchangeRateValue, error) {
	currencyAmount := ExchangeRateValue{
		value: value,
	}

	if err := currencyAmount.Validate(); err != nil {
		return ExchangeRateValue{}, err
	}

	return currencyAmount, nil
}

func (v ExchangeRateValue) Equal(ref ExchangeRateValue) bool {
	return math.Abs(v.Value()-ref.Value()) <= v.epsilon()
}

func (v ExchangeRateValue) epsilon() float64 {
	return 1e-6
}

func (v ExchangeRateValue) Value() float64 {
	return v.value
}

func (v ExchangeRateValue) Validate() error {
	return ozzo.ValidateStruct(&v,
		ozzo.Field(&v.value, ozzo.Min(0.0)),
	)
}
