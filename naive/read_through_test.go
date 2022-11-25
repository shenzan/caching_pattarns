package naive

import (
	"CachingPatterns"
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReadThroughCache(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	for client.Ping(context.Background()).Err() != nil {

	}

	redis := &Caching.RedisCache{
		Client: client,
	}
	cache := &ReadThroughCache{
		Cache: redis,
		DBLoader: LoadFunc(func(ctx context.Context, key string) (any, error) {
			return "abc", nil
		}),
		ttl: 30 * time.Second,
	}

	val, err := cache.Get(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, "abc", val)
	val, err = redis.Get(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, "abc", val)
	cache.Delete(context.Background(), "123")
}
