package golimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/streamdp/microcache"
)

type Cache interface {
	Get(key string) (any, error)
	Set(key string, value any, expiration time.Duration) error
}

type LimitCache struct {
	c   Cache
	ttl time.Duration
}

func NewLimitCache(ctx context.Context, ttl time.Duration) *LimitCache {
	return &LimitCache{
		c:   microcache.New(ctx, nil),
		ttl: ttl,
	}
}

const (
	hitsPrefix     = "hits:"
	deadlinePrefix = "deadline:"
)

func (a *LimitCache) Set(key string, hits int, deadline int64) (err error) {
	if err = a.c.Set(hitsPrefix+key, hits, a.ttl); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	if err = a.c.Set(deadlinePrefix+key, deadline, a.ttl); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	return nil
}

func (a *LimitCache) Get(key string) (hits int, deadline int64, err error) {
	rawHits, err := a.c.Get(hitsPrefix + key)
	if err != nil {
		return
	}

	rawDeadline, err := a.c.Get(deadlinePrefix + key)
	if err != nil {
		return
	}

	return rawHits.(int), rawDeadline.(int64), nil
}
