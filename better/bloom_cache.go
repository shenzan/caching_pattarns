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
	if b.BloomFilter.Exist(key) {
		return b.Cache.Get(ctx, key)
	}
	return nil, nil
}

func NewBloomCache(cache Caching.Cache, filter BloomFilter) *BloomCache {
	return &BloomCache{
		Cache:       cache,
		BloomFilter: filter,
	}
}
