package Caching

import (
	"context"
	"errors"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Cache interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, val any, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

type DBLoader interface {
	Load(ctx context.Context, key string) (any, error)
}

type DBStore interface {
	Save(ctx context.Context, key string, val any) error
}
