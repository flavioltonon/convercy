package mongodb

import (
	"context"
	"errors"

	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/infrastructure/repository/mongodb/dao"
	"convercy/infrastructure/repository/mongodb/mappers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RegisteredCurrenciesRepository struct {
	mapper     *mappers.RegisteredCurrenciesMapper
	repository *Repository
}

func NewRegisteredCurrenciesRepository(mapper *mappers.RegisteredCurrenciesMapper, repository *Repository) *RegisteredCurrenciesRepository {
	return &RegisteredCurrenciesRepository{
		mapper:     mapper,
		repository: repository,
	}
}

func (r *RegisteredCurrenciesRepository) collection() *mongo.Collection {
	return r.repository.database.Collection("currencies")
}

func (r *RegisteredCurrenciesRepository) GetRegisteredCurrencies() (*aggregate.RegisteredCurrencies, error) {
	var registeredCurrencies dao.RegisteredCurrencies

	ctx := context.Background()

	err := r.collection().FindOne(ctx, bson.D{}).Decode(&registeredCurrencies)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, domain.ErrRegisteredCurrenciesNotFound()
	}
	if err != nil {
		return nil, err
	}

	return r.mapper.ToModel(registeredCurrencies)
}

func (r *RegisteredCurrenciesRepository) SaveRegisteredCurrencies(registeredCurrencies *aggregate.RegisteredCurrencies) error {
	_, err := r.collection().UpdateOne(
		context.Background(),
		primitive.M{"_id": registeredCurrencies.ClientID()},
		primitive.M{"$set": r.mapper.ToDAO(registeredCurrencies)},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}

	return nil
}
