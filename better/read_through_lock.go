package better

import (
	"CachingPatterns"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type ReadThroughCache struct {
	Caching.Cache
	Caching.DBLoader
	ttl time.Duration

	mutex sync.RWMutex
}

func (r *ReadThroughCache) Get(ctx context.Context, key string) (any, error) {
	r.mutex.RLock()
	val, err := r.Cache.Get(ctx, key)
	r.mutex.RUnlock()
	if err != nil && err != Caching.ErrKeyNotFound {
		return nil, err
	}
	//cache missed
	if err == Caching.ErrKeyNotFound {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		val, err = r.DBLoader.Load(ctx, key)
		if err != nil {
			return nil, fmt.Errorf("DB: Cannot load data , %w", err)
		}
		err = r.Cache.Set(ctx, key, val, r.ttl)
		if err != nil {
			log.Fatalln(err)
		}
		return val, nil
	}
	return val, nil
}

func (r *ReadThroughCache) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.Cache.Set(ctx, key, val, ttl)
}

type LoadFunc func(ctx context.Context, key string) (any, error)

func (f LoadFunc) Load(ctx context.Context, key string) (any, error) {
	return f(ctx, key)
}

func NewReadThroughCache(cache Caching.Cache) *ReadThroughCache {
	return &ReadThroughCache{
		Cache: cache,
		DBLoader: LoadFunc(func(ctx context.Context, key string) (any, error) {
			return "abc", nil
		}),
		ttl: 30 * time.Second,
	}
}
