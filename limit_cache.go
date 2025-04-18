package golimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/streamdp/microcache"
)

type Cache interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
}

type LimitCache struct {
	c   Cache
	ttl time.Duration
}

func NewLimitCache(ctx context.Context, ttl time.Duration) *LimitCache {
	return &LimitCache{
		c:   microcache.New(ctx, -1),
		ttl: ttl,
	}
}

const (
	hitsPrefix     = "hits:"
	deadlinePrefix = "deadline:"
)

func (a *LimitCache) Set(ctx context.Context, key string, hits int, deadline int64) (err error) {
	if err = a.c.Set(ctx, hitsPrefix+key, hits, a.ttl); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	if err = a.c.Set(ctx, deadlinePrefix+key, deadline, a.ttl); err != nil {
		return fmt.Errorf("cache: %w", err)
	}

	return nil
}

func (a *LimitCache) Get(ctx context.Context, key string) (hits int, deadline int64, err error) {
	rawHits, err := a.c.Get(ctx, hitsPrefix+key)
	if err != nil {
		return
	}

	rawDeadline, err := a.c.Get(ctx, deadlinePrefix+key)
	if err != nil {
		return
	}

	return rawHits.(int), rawDeadline.(int64), nil
}
