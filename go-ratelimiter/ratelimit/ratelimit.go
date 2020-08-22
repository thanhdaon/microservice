package ratelimit

import (
	"time"

	"github.com/segmentio/ksuid"
)

type RateLimiter interface {
	Acquire() (*Token, error)
	Release(*Token)
}

type Config struct {
	// Limit determines how many rate limit tokens can be active at a time
	Limit int

	// Throttle is the min time between requests for a Throttle Rate Limiter
	Throttle time.Duration

	// TokenResetsAfter is the maximum amount of time a token can live before being
	// forcefully released - if set to zero time then the token may live forever
	TokenResetsAfter time.Duration
}

type tokenFactory func() *Token

type Token struct {
	ID        string
	CreatedAt time.Time
}

func NewToken() *Token {
	return &Token{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
	}
}
