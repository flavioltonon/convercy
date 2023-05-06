package repositories

import "convercy/domain/aggregate"

//go:generate mockery --name RegisteredCurrenciesRepository
type RegisteredCurrenciesRepository interface {
	GetRegisteredCurrencies() (*aggregate.RegisteredCurrencies, error)
	SaveRegisteredCurrencies(registeredCurrencies *aggregate.RegisteredCurrencies) error
}
