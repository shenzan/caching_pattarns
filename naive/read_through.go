package naive

import (
	"CachingPatterns"
	"context"
	"fmt"
	"log"
	"time"
)

type ReadThroughCache struct {
	Caching.Cache
	Caching.DBLoader
	ttl time.Duration
}

func (r *ReadThroughCache) Get(ctx context.Context, key string) (any, error) {
	val, err := r.Cache.Get(ctx, key)
	if err != nil && err != Caching.ErrKeyNotFound {
		return nil, err
	}
	//cache missed
	if err == Caching.ErrKeyNotFound {
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

type LoadFunc func(ctx context.Context, key string) (any, error)

func (f LoadFunc) Load(ctx context.Context, key string) (any, error) {
	return f(ctx, key)
}
