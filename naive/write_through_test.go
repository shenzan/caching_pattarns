package naive

import (
	"CachingPatterns"
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWriteThroughCache(t *testing.T) {
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
	m := make(map[string]any)
	cache := &WriteThoughCache{
		Cache: redis,
		DBStore: SaveFunc(func(ctx context.Context, key string, val any) error {
			m[key] = val
			return nil
		}),
	}
	cache.Set(context.Background(), "write", "asdf", 10*time.Second)
	assert.Equal(t, "asdf", m["write"])
}
