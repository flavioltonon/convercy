package valueobject

import (
	"math"

	"convercy/domain"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type CurrencyAmount struct {
	value float64
}

func NewCurrencyAmount(value float64) (CurrencyAmount, error) {
	currencyAmount := CurrencyAmount{
		value: value,
	}

	if err := currencyAmount.Validate(); err != nil {
		return CurrencyAmount{}, err
	}

	return currencyAmount, nil
}

func (v CurrencyAmount) Equal(ref CurrencyAmount) bool {
	return math.Abs(v.Value()-ref.Value()) <= v.epsilon()
}

func (v CurrencyAmount) epsilon() float64 {
	return 1e-2
}

func (v CurrencyAmount) Value() float64 {
	return v.value
}

func (v CurrencyAmount) Validate() error {
	if err := ozzo.ValidateStruct(&v, ozzo.Field(&v.value, ozzo.Required, ozzo.Min(0.0))); err != nil {
		return domain.ErrInvalidCurrencyAmount(err)
	}

	return nil
}
