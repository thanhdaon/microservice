package ratelimit

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func New(redisUri string) (*RateLimit, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &RateLimit{
		rdb:                rdb,
		reservationSeconds: 10,
	}, nil
}

type RateLimit struct {
	rdb                *redis.Client
	cycleDuration      time.Duration
	reservationSeconds int
}

func (r *RateLimit) DefineRateLimits(arena string, limits []Limit) error {
	// { queue_name -> allowed calls per cycleDuration }
	m := map[string]interface{}{}
	for _, l := range limits {
		m[fmt.Sprintf("%s:%s", arena, l.ID)] = l.Limit
	}
	return r.rdb.HSet(context.Background(), "rate_limits", m).Err()
}

func (r *RateLimit) Enqueue(arena, queue, payload string) error {
	return enqueueScript.Run(
		context.Background(), r.rdb, []string{},
		arena, queue, payload,
	).Err()
}

func (r *RateLimit) Dequeue(arena string) (payload, txID string, err error) {
	txid := uuid.New().String()
	v, err := dequeueScript.Run(
		context.Background(), r.rdb, []string{},
		arena, r.getTime(), txid, r.reservationSeconds,
	).Result()

	if err != nil {
		if err == redis.Nil {
			return "", "", ERR_QUEUE_EMPTY
		}
		return "", "", err
	}

	return v.(string), txid, nil
}

func (r *RateLimit) Commit(txid string) error {
	return commitScript.Run(
		context.Background(), r.rdb, []string{},
		txid,
	).Err()
}

func (r *RateLimit) Reap() error {
	return nil
}

func (r *RateLimit) Prune() error {
	return r.rdb.FlushAll(context.Background()).Err()
}

func (r *RateLimit) Close() {
	r.rdb.Close()
}

func (r *RateLimit) getTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
