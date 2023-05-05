package mongodb

import (
	"context"
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Repository is a MongoDB repository
type Repository struct {
	client   *mongo.Client
	database *mongo.Database
}

type Options struct {
	Database DatabaseOptions
}

func (o Options) Validate() error {
	return ozzo.ValidateStruct(&o,
		ozzo.Field(&o.Database, ozzo.Required),
	)
}

type DatabaseOptions struct {
	Name string
	URI  string
}

func (o DatabaseOptions) Validate() error {
	return ozzo.ValidateStruct(&o,
		ozzo.Field(&o.Name, ozzo.Required),
		ozzo.Field(&o.URI, ozzo.Required),
	)
}

// NewRepository creates a new Repository
func NewRepository(o *Options) (*Repository, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	client, err := mongo.NewClient(options.Client().
		ApplyURI(o.Database.URI).
		SetConnectTimeout(5 * time.Second).
		SetReadConcern(readconcern.Available()).
		SetReadPreference(readpref.SecondaryPreferred()),
	)
	if err != nil {
		return nil, err
	}

	return &Repository{
		client:   client,
		database: client.Database(o.Database.Name),
	}, nil
}

// Connect initializes the Repository
func (r *Repository) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.client.Connect(ctx); err != nil {
		return err
	}

	return r.client.Ping(ctx, readpref.Primary())
}

// Disconnect interrupts the connection of the Repository with MongoDB servers
func (r *Repository) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.client.Disconnect(ctx)
}
