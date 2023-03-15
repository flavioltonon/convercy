package entity

import (
	"convercy/domain"
	"convercy/domain/valueobject"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type Currency struct {
	id   valueobject.CurrencyID
	code valueobject.CurrencyCode
}

func NewCurrency(
	id valueobject.CurrencyID,
	code valueobject.CurrencyCode) (*Currency, error) {
	currency := &Currency{
		id:   id,
		code: code,
	}

	if err := currency.Validate(); err != nil {
		return nil, err
	}

	return currency, nil
}

func (a *Currency) ID() valueobject.CurrencyID {
	return a.id
}

func (a *Currency) Code() valueobject.CurrencyCode {
	return a.code
}

func (a *Currency) Validate() error {
	if err := ozzo.ValidateStruct(a,
		ozzo.Field(&a.id, ozzo.Required),
		ozzo.Field(&a.code, ozzo.Required),
	); err != nil {
		return domain.ErrInvalidCurrency(err)
	}

	return nil
}
