package redis

import (
	"context"
	"convercy/domain"
	"convercy/domain/aggregate"
	"convercy/domain/valueobject"
	"convercy/infrastructure/repository/redis/dao"
	"convercy/infrastructure/repository/redis/mappers"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CurrencyExchangeRatesCache struct {
	mapper *mappers.CurrencyExchangeRatesMapper
	cache  *Cache
	ttl    time.Duration
}

func NewCurrencyExchangeRatesCache(mapper *mappers.CurrencyExchangeRatesMapper, cache *Cache, ttl time.Duration) *CurrencyExchangeRatesCache {
	return &CurrencyExchangeRatesCache{
		mapper: mapper,
		cache:  cache,
		ttl:    ttl,
	}
}

func (c *CurrencyExchangeRatesCache) generateKey(currencyCode valueobject.CurrencyCode) string {
	return fmt.Sprintf("currencies-exchange-rates:%s", currencyCode)
}

func (c *CurrencyExchangeRatesCache) GetCurrencyExchangeRates(currencyCode valueobject.CurrencyCode) (*aggregate.CurrencyExchangeRates, error) {
	var dao dao.CurrencyExchangeRates

	ctx, cancel := context.WithTimeout(context.Background(), c.cache.options.ConnectionTimeout)
	defer cancel()

	key := c.generateKey(currencyCode)

	if err := c.cache.client.Get(ctx, key).Scan(&dao); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, domain.ErrCurrencyExchangeRatesNotFound()
		}

		return nil, err
	}

	return c.mapper.ToModel(dao)
}

func (c *CurrencyExchangeRatesCache) SaveCurrencyExchangeRates(currencyExchangeRates *aggregate.CurrencyExchangeRates) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.cache.options.ConnectionTimeout)
	defer cancel()
	key := c.generateKey(currencyExchangeRates.CurrencyCode())
	dao := c.mapper.ToDAO(currencyExchangeRates)
	return c.cache.client.Set(ctx, key, dao, c.ttl).Err()
}

func (c *CurrencyExchangeRatesCache) DeleteCurrencyExchangeRates(currencyCode valueobject.CurrencyCode) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.cache.options.ConnectionTimeout)
	defer cancel()
	key := c.generateKey(currencyCode)
	return c.cache.client.Del(ctx, key).Err()
}
