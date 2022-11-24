package better

import (
	Caching "CachingPatterns"
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBloomFilterReadThrough(t *testing.T) {
	filter := &filter{}
	filter.Add("Monday")
	filter.Add("Tuesday")
	filter.Add("Wednesday")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	for client.Ping(context.Background()).Err() != nil {

	}

	redisCache := &Caching.RedisCache{
		Client: client,
	}

	rtCache := &ReadThroughCache{
		Cache: redisCache,
		DBLoader: LoadFunc(func(ctx context.Context, key string) (any, error) {
			return "abc", nil
		}),
		ttl: 30 * time.Second,
	}

	bfCache := &BloomCache{
		Cache:       rtCache,
		BloomFilter: filter,
	}

	val, err := bfCache.Get(context.Background(), "Monday")
	assert.NoError(t, err)
	assert.Equal(t, "abc", val)
	val, err = bfCache.Get(context.Background(), "Friday")
	assert.Nil(t, val, err)
}
