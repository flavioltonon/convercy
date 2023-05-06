package repositories

import "convercy/domain/valueobject"

// CurrenciesRepository is a service responsible for providing all known currencies
//
//go:generate mockery --name CurrenciesRepository
type CurrenciesRepository interface {
	// ListCurrencyCodes should return a list with all known currency codes
	ListCurrencyCodes() (valueobject.CurrencyCodes, error)
}
