package usecases

import "convercy/domain/valueobject"

// CurrenciesService is a service responsible for providing all known currencies
//
//go:generate mockery --name CurrenciesService
type CurrenciesService interface {
	// ListCurrencyCodes should return a list with all known currency codes
	ListCurrencyCodes() (valueobject.CurrencyCodes, error)
}
