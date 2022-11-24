package better

import (
	Caching "CachingPatterns"
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"time"
)

type SingleflightCache struct {
	Caching.Cache
	Caching.DBLoader
	ttl   time.Duration
	group *singleflight.Group
}

func (s *SingleflightCache) Get(ctx context.Context, key string) (any, error) {
	val, err := s.Cache.Get(ctx, key)
	if err != nil && err != Caching.ErrKeyNotFound {
		return nil, err
	}
	//cache missed
	if err == Caching.ErrKeyNotFound {
		defer func() {
			s.group.Forget(key)
		}()
		val, err, _ = s.group.Do(key, func() (interface{}, error) {
			v, e := s.DBLoader.Load(ctx, key)
			if e != nil {
				return nil, fmt.Errorf("DB: Cannot load data , %w", e)
			}
			e = s.Cache.Set(ctx, key, v, s.ttl)
			if e != nil {
				log.Fatalln(err)
			}
			return v, nil
		})
	}
	return val, nil
}
