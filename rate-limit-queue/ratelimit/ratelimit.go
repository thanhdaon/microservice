package ratelimit

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func New(redisUri string) (*RateLimit, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &RateLimit{rdb: rdb}, nil
}

type RateLimit struct {
	rdb *redis.Client
}

func (r *RateLimit) DefineRateLimits(arena string, limits []Limit) error {
	return nil
}

func (r *RateLimit) Enqueue(arena, queue, payload string) error {
	return nil
}

func (r *RateLimit) Dequeue(arena string) (payload, txID string, err error) {
	return "", "", nil
}

func (r *RateLimit) Commit(txID string) error {
	return nil
}

func (r *RateLimit) Reap() error {
	return nil
}
