package better

import (
	Caching "CachingPatterns"
	"context"
)

type BloomCache struct {
	Caching.Cache
	BloomFilter
}

func (b *BloomCache) Get(ctx context.Context, key string) (any, error) {
	if exist := b.BloomFilter.Exist(key); exist {
		return b.Cache.Get(ctx, key)
	}
	return nil, nil
}

type BloomFilter interface {
	Exist(key string) bool
}

func NewBloomCache(cache Caching.Cache, filter BloomFilter) *BloomCache {
	return &BloomCache{
		Cache:       cache,
		BloomFilter: filter,
	}
}
