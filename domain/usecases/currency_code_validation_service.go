package usecases

import "convercy/domain/valueobject"

// CurrencyCodeValidationService is a service capable of validating currency codes by cross-validating it with all known currency
// codes
//
//go:generate mockery --name CurrencyCodeValidationService
type CurrencyCodeValidationService interface {
	// ValidateCurrencyCode validates a given CurrencyCode
	ValidateCurrencyCode(code valueobject.CurrencyCode) error
}
