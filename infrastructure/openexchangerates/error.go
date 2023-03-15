package openexchangerates

import "fmt"

type ErrNotFound struct {
	message string
}

func (e ErrNotFound) Error() string {
	return e.message
}

func errExchangeRateNotFound(code string) error {
	return ErrNotFound{
		message: fmt.Sprintf("exchange rate not found for currency code %s", code),
	}
}
