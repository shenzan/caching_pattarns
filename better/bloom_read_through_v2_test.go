package better

import (
	Caching "CachingPatterns"
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBloomFilterReadThroughV2(t *testing.T) {
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

	bfCache := NewBloomCache(
		NewReadThroughCache(
			Caching.NewRedisCache(client)), filter)

	val, err := bfCache.Get(context.Background(), "Monday")
	assert.NoError(t, err)
	assert.Equal(t, "abc", val)
	val, err = bfCache.Get(context.Background(), "Friday")
	assert.Nil(t, val, err)
}
