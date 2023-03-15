package valueobject

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type ClientID struct {
	value string
}

func GenerateClientID() ClientID {
	return ClientID{
		value: uuid.NewString(),
	}
}

func NewClientID(value string) (ClientID, error) {
	currencyID := ClientID{
		value: value,
	}

	if err := currencyID.Validate(); err != nil {
		return ClientID{}, err
	}

	return currencyID, nil
}

func (v ClientID) Equal(ref ClientID) bool {
	return v.String() == ref.String()
}

func (v ClientID) String() string {
	return v.value
}

func (v ClientID) Validate() error {
	return is.UUIDv4.Validate(v.String())
}
