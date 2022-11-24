package naive

import (
	"CachingPatterns"
	"context"
	"time"
)

type WriteThoughCache struct {
	Caching.Cache
	Caching.DBStore
}

func (w *WriteThoughCache) Set(ctx context.Context,
	key string, val any, expiration time.Duration) error {
	err := w.Cache.Set(ctx, key, val, expiration)
	if err != nil {
		return err
	}
	return w.DBStore.Save(ctx, key, val)
}

type SaveFunc func(ctx context.Context, key string, val any) error

func (f SaveFunc) Save(ctx context.Context, key string, val any) error {
	return f(ctx, key, val)
}
