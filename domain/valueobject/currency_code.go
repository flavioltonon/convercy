package valueobject

import (
	"convercy/domain"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type CurrencyCode struct {
	value string
}

func NewCurrencyCode(value string) (CurrencyCode, error) {
	currencyCode := CurrencyCode{
		value: value,
	}

	if err := currencyCode.Validate(); err != nil {
		return CurrencyCode{}, err
	}

	return currencyCode, nil
}

func (v CurrencyCode) Equal(ref CurrencyCode) bool {
	return v.String() == ref.String()
}

func (v CurrencyCode) String() string {
	return v.value
}

func (v CurrencyCode) Validate() error {
	if err := ozzo.ValidateStruct(&v, ozzo.Field(&v.value, ozzo.Required, ozzo.Length(3, 3))); err != nil {
		return domain.ErrInvalidCurrencyCode(err)
	}

	return nil
}
