package usecases

import "convercy/domain/valueobject"

//go:generate mockery --name CurrenciesService
type CurrenciesService interface {
	ListCurrencyCodes() (valueobject.CurrencyCodes, error)
}
