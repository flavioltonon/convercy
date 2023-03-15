package mappers

import (
	"convercy/domain/entity"
	"convercy/domain/valueobject"
	"convercy/infrastructure/repository/mongodb/dao"
)

type CurrencyMapper struct{}

func NewCurrencyMapper() *CurrencyMapper {
	return new(CurrencyMapper)
}

func (m *CurrencyMapper) ToDAO(model *entity.Currency) dao.Currency {
	return dao.Currency{
		ID:   model.ID().String(),
		Code: model.Code().String(),
	}
}

func (m *CurrencyMapper) ToDAOs(models []*entity.Currency) []dao.Currency {
	daos := make([]dao.Currency, 0, len(models))

	for _, model := range models {
		daos = append(daos, m.ToDAO(model))
	}

	return daos
}

func (m *CurrencyMapper) ToModel(dao dao.Currency) (*entity.Currency, error) {
	id, err := valueobject.NewCurrencyID(dao.ID)
	if err != nil {
		return nil, err
	}

	code, err := valueobject.NewCurrencyCode(dao.Code)
	if err != nil {
		return nil, err
	}

	return entity.NewCurrency(id, code)
}

func (m *CurrencyMapper) ToModels(daos []dao.Currency) ([]*entity.Currency, error) {
	models := make([]*entity.Currency, 0, len(daos))

	for _, dao := range daos {
		model, err := m.ToModel(dao)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return models, nil
}
