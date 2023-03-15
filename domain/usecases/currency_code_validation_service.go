package usecases

import "convercy/domain/valueobject"

//go:generate mockery --name CurrencyCodeValidationService
type CurrencyCodeValidationService interface {
	ValidateCurrencyCode(code valueobject.CurrencyCode) error
}
