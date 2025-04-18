package golimiter

import (
	"context"
	"log"
	"time"
)

type Limiter struct {
	c *LimitCache
}

func New(ctx context.Context, ttl time.Duration) *Limiter {
	return &Limiter{NewLimitCache(ctx, ttl)}
}

func (l *Limiter) Allow(ctx context.Context, key string, rate int, period time.Duration) bool {
	now := time.Now().UnixMicro()

	if hits, deadline, err := l.c.Get(ctx, key); err == nil && deadline >= now {
		if hits >= rate {
			return false
		}
		if err = l.c.Set(ctx, key, hits+1, deadline); err != nil {
			log.Println(err)
		}
		return true
	}

	if err := l.c.Set(ctx, key, 1, now+period.Microseconds()); err != nil {
		log.Println(err)
	}

	return true
}
