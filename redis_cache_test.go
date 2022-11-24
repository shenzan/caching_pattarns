package Caching

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisCache(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	cache := &RedisCache{
		client: client,
	}
	ctx := context.Background()
	cache.Delete(ctx, "name")
	val, _ := cache.Get(ctx, "name")
	assert.Empty(t, val)
	cache.Set(ctx, "name", "Thomas", time.Minute)
	val, _ = cache.Get(ctx, "name")
	assert.Equal(t, "Thomas", val)
}
