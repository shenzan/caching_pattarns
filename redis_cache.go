package Caching

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v9"
	"time"
)

type RedisCache struct {
	Client redis.Cmdable
}

func (r *RedisCache) Get(ctx context.Context, key string) (any, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, ErrKeyNotFound
	}
	return val, err
}

func (r *RedisCache) Set(ctx context.Context, key string, val any,
	ttl time.Duration) error {
	msg, err := r.Client.Set(ctx, key, val, ttl).Result()
	if err != nil {
		return fmt.Errorf("%w, errMessage %s", err, msg)
	}
	return nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	_, err := r.Client.Del(ctx, key).Result()
	return err
}

func NewRedisCache(client redis.Cmdable) *RedisCache {
	return &RedisCache{
		Client: client,
	}
}
