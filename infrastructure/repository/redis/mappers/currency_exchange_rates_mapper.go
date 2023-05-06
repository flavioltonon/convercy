package mappers

import (
	"convercy/domain/aggregate"
	"convercy/domain/valueobject"
	"convercy/infrastructure/repository/redis/dao"
)

type CurrencyExchangeRatesMapper struct {
	exchangeRatesMapper *ExchangeRatesMapper
}

func NewCurrencyExchangeRatesMapper(exchangeRatesMapper *ExchangeRatesMapper) *CurrencyExchangeRatesMapper {
	return &CurrencyExchangeRatesMapper{exchangeRatesMapper: exchangeRatesMapper}
}

func (m *CurrencyExchangeRatesMapper) ToDAO(model *aggregate.CurrencyExchangeRates) dao.CurrencyExchangeRates {
	return dao.CurrencyExchangeRates{
		CurrencyCode:  model.CurrencyCode().String(),
		ExchangeRates: m.exchangeRatesMapper.ToDAOs(model.ExchangeRates()),
	}
}

func (m *CurrencyExchangeRatesMapper) ToModel(dao dao.CurrencyExchangeRates) (*aggregate.CurrencyExchangeRates, error) {
	code, err := valueobject.NewCurrencyCode(dao.CurrencyCode)
	if err != nil {
		return nil, err
	}

	exchangeRates, err := m.exchangeRatesMapper.ToModels(dao.ExchangeRates)
	if err != nil {
		return nil, err
	}

	return aggregate.NewCurrencyExchangeRates(code, exchangeRates...)
}

type ExchangeRatesMapper struct {
	exchangeRateUnitsMapper *ExchangeRateUnitsMapper
}

func NewExchangeRatesMapper(exchangeRateUnitsMapper *ExchangeRateUnitsMapper) *ExchangeRatesMapper {
	return &ExchangeRatesMapper{exchangeRateUnitsMapper: exchangeRateUnitsMapper}
}

func (m *ExchangeRatesMapper) ToDAO(model valueobject.ExchangeRate) dao.ExchangeRate {
	return dao.ExchangeRate{
		Rate: model.Rate().Value(),
		Unit: m.exchangeRateUnitsMapper.ToDAO(model.Unit()),
	}
}

func (m *ExchangeRatesMapper) ToDAOs(models valueobject.ExchangeRates) []dao.ExchangeRate {
	daos := make([]dao.ExchangeRate, 0, len(models))

	for _, model := range models {
		daos = append(daos, m.ToDAO(model))
	}

	return daos
}

func (m *ExchangeRatesMapper) ToModel(dao dao.ExchangeRate) (valueobject.ExchangeRate, error) {
	return valueobject.NewExchangeRate(dao.Rate, dao.Unit.Base, dao.Unit.Target)
}

func (m *ExchangeRatesMapper) ToModels(daos []dao.ExchangeRate) (valueobject.ExchangeRates, error) {
	models := make([]valueobject.ExchangeRate, 0, len(daos))

	for _, dao := range daos {
		model, err := m.ToModel(dao)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return models, nil
}

type ExchangeRateUnitsMapper struct{}

func NewExchangeRateUnitsMapper() *ExchangeRateUnitsMapper {
	return new(ExchangeRateUnitsMapper)
}

func (m *ExchangeRateUnitsMapper) ToDAO(model valueobject.ExchangeRateUnit) dao.ExchangeRateUnit {
	return dao.ExchangeRateUnit{
		Base:   model.BaseCurrencyCode().String(),
		Target: model.TargetCurrencyCode().String(),
	}
}
