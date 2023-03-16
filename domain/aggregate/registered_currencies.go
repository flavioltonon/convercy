package aggregate

import (
	"convercy/domain"
	"convercy/domain/entity"
	"convercy/domain/valueobject"
)

// RegisteredCurrencies is an aggregate intended to ensure the business invariants related to the currency registration bounded context
type RegisteredCurrencies struct {
	clientID   valueobject.ClientID
	currencies []*entity.Currency
}

func NewRegisteredCurrencies(clientID valueobject.ClientID, currencies ...*entity.Currency) *RegisteredCurrencies {
	return &RegisteredCurrencies{
		clientID:   clientID,
		currencies: currencies,
	}
}

func (a *RegisteredCurrencies) ClientID() valueobject.ClientID {
	return a.clientID
}

func (a *RegisteredCurrencies) Currencies() []*entity.Currency {
	if a == nil || a.currencies == nil {
		return make([]*entity.Currency, 0)
	}

	return a.currencies
}

func (a *RegisteredCurrencies) RegisterCurrency(newCurrency *entity.Currency) error {
	if a.HasCurrencyWithCode(newCurrency.Code()) {
		return domain.ErrCurrencyAlreadyExists()
	}

	a.currencies = append(a.currencies, newCurrency)
	return nil
}

func (a *RegisteredCurrencies) UnregisterCurrency(currencyID valueobject.CurrencyID) error {
	for i, registeredCurrency := range a.currencies {
		if registeredCurrency.ID().Equal(currencyID) {
			a.currencies = append(a.currencies[:i], a.currencies[i+1:]...)
			return nil
		}
	}

	return domain.ErrCurrencyNotFound()
}

func (a *RegisteredCurrencies) FindCurrencyByCode(code valueobject.CurrencyCode) (*entity.Currency, error) {
	for _, currency := range a.currencies {
		if currency.Code().Equal(code) {
			return currency, nil
		}
	}

	return nil, domain.ErrCurrencyNotFound()
}

func (a *RegisteredCurrencies) HasCurrencyWithCode(code valueobject.CurrencyCode) bool {
	for _, currency := range a.currencies {
		if currency.Code().Equal(code) {
			return true
		}
	}

	return false
}
