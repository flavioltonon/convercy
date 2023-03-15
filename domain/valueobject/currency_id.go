package valueobject

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type CurrencyID struct {
	value string
}

func GenerateCurrencyID() CurrencyID {
	return CurrencyID{
		value: uuid.NewString(),
	}
}

func NewCurrencyID(value string) (CurrencyID, error) {
	currencyID := CurrencyID{
		value: value,
	}

	if err := currencyID.Validate(); err != nil {
		return CurrencyID{}, err
	}

	return currencyID, nil
}

func (v CurrencyID) Equal(ref CurrencyID) bool {
	return v.String() == ref.String()
}

func (v CurrencyID) String() string {
	return v.value
}

func (v CurrencyID) Validate() error {
	return is.UUIDv4.Validate(v.String())
}
