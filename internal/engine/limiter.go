package engine

import (
	"context"
	"time"
)

type RateLimiter struct {
	ticker  *time.Ticker
	limiter chan struct{}
}

func NewRateLimiter(rate int) *RateLimiter {
	ch := make(chan struct{}, rate)
	for i := 0; i < rate; i++ {
		ch <- struct{}{}
	}
	return &RateLimiter{limiter: ch}
}

func (r *RateLimiter) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case token := <-r.limiter:
		go func() {
			time.Sleep(1 * time.Second)
			r.limiter <- token
		}()
		return nil
	}
}
