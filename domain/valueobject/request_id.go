package valueobject

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type RequestID struct {
	value string
}

func GenerateRequestID() RequestID {
	return RequestID{
		value: uuid.NewString(),
	}
}

func NewRequestID(value string) (RequestID, error) {
	currencyID := RequestID{
		value: value,
	}

	if err := currencyID.Validate(); err != nil {
		return RequestID{}, err
	}

	return currencyID, nil
}

func (v RequestID) Equal(ref RequestID) bool {
	return v.String() == ref.String()
}

func (v RequestID) String() string {
	return v.value
}

func (v RequestID) Validate() error {
	return is.UUIDv4.Validate(v.String())
}
