package mappers

import (
	"convercy/domain/aggregate"
	"convercy/domain/valueobject"
	"convercy/infrastructure/repository/mongodb/dao"
)

type RegisteredCurrenciesMapper struct {
	currencyMapper *CurrencyMapper
}

func NewRegisteredCurrenciesMapper(currencyMapper *CurrencyMapper) *RegisteredCurrenciesMapper {
	return &RegisteredCurrenciesMapper{currencyMapper: currencyMapper}
}

func (m *RegisteredCurrenciesMapper) ToDAO(model *aggregate.RegisteredCurrencies) dao.RegisteredCurrencies {
	return dao.RegisteredCurrencies{
		ClientID:   model.ClientID().String(),
		Currencies: m.currencyMapper.ToDAOs(model.Currencies()),
	}
}

func (m *RegisteredCurrenciesMapper) ToModel(dao dao.RegisteredCurrencies) (*aggregate.RegisteredCurrencies, error) {
	clientID, err := valueobject.NewClientID(dao.ClientID)
	if err != nil {
		return nil, err
	}

	currencies, err := m.currencyMapper.ToModels(dao.Currencies)
	if err != nil {
		return nil, err
	}

	return aggregate.NewRegisteredCurrencies(clientID, currencies...), nil
}
