package redis

import (
	"context"
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/redis/go-redis/v9"
)

// Cache is a Redis cache
type Cache struct {
	client  *redis.Client
	options *Options
}

type Options struct {
	Address           string
	ConnectionTimeout time.Duration
}

func (o Options) Validate() error {
	return ozzo.ValidateStruct(&o,
		ozzo.Field(&o.Address, ozzo.Required),
		ozzo.Field(&o.ConnectionTimeout, ozzo.Required, ozzo.Min(1*time.Second)),
	)
}

// NewCache creates a new Cache
func NewCache(options *Options) (*Cache, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr: options.Address,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Cache{
		client:  client,
		options: options,
	}, nil
}
