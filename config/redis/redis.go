package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimit struct {
	client *redis.Client
	window time.Duration
	limit  int
	ctx    context.Context
}

func NewRateLimiter(client *redis.Client, window time.Duration, limit int) *RateLimit {
	return &RateLimit{
		client: client,
		window: window,
		limit:  limit,
		ctx:    context.Background(),
	}
}

func (rl *RateLimit) Allow(key string) bool {
	pipe := rl.client.TxPipeline()

	incr := pipe.Incr(rl.ctx, key)
	pipe.Expire(rl.ctx, key, rl.window)

	_, err := pipe.Exec(rl.ctx)
	if err != nil {
		return false
	}

	return incr.Val() <= int64(rl.limit)
}
