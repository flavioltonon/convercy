package domain

import (
	"fmt"
)

type Error struct {
	message string
}

type ErrAlreadyExists Error

func (e ErrAlreadyExists) Error() string {
	return e.message
}

func ErrCurrencyAlreadyExists() error {
	return ErrNotFound{
		message: "currency already exists",
	}
}

func ErrExchangeRateAlreadyExists() error {
	return ErrNotFound{
		message: "exchange rate already exists",
	}
}

type ErrNotFound Error

func (e ErrNotFound) Error() string {
	return e.message
}

func ErrCurrencyNotFound() error {
	return ErrNotFound{
		message: "currency not found",
	}
}

func ErrCurrencyCodeNotFound() error {
	return ErrNotFound{
		message: "currency not found",
	}
}

func ErrCurrencyNotRegistered() error {
	return ErrNotFound{
		message: "currency not registered",
	}
}

func ErrExchangeRateNotFound() error {
	return ErrNotFound{
		message: "exchange rate not found",
	}
}

func ErrRegisteredCurrenciesNotFound() error {
	return ErrNotFound{
		message: "registered currencies not found",
	}
}

type ErrValidationFailure Error

func (e ErrValidationFailure) Error() string {
	return e.message
}

func ErrIncompatibleExchangeRates() error {
	return ErrValidationFailure{
		message: "both exchange rates should share the same target currency (e.g. USD)",
	}
}

func ErrUnexpectedExchangeRateBaseCurrencyCode(expected, actual string) error {
	return ErrValidationFailure{
		message: fmt.Sprintf("base currency in exchange rate should be %s but is %s", expected, actual),
	}
}

func ErrInvalidCurrencyAmount(cause error) error {
	return ErrValidationFailure{
		message: fmt.Sprintf("invalid currency amount: %v", cause),
	}
}

func ErrInvalidCurrencyCode(cause error) error {
	return ErrValidationFailure{
		message: fmt.Sprintf("invalid currency code: %v", cause),
	}
}

func ErrInvalidCurrency(cause error) error {
	return ErrValidationFailure{
		message: fmt.Sprintf("invalid currency: %v", cause),
	}
}

func ErrInvalidMoney(cause error) error {
	return ErrValidationFailure{
		message: fmt.Sprintf("invalid money: %v", cause),
	}
}
