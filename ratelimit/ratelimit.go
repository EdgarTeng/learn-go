package ratelimit

import (
	"time"

	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

var (
	tokenBucket *rate.Limiter
	leakyBucket ratelimit.Limiter
)

func init() {
	tokenBucket = rate.NewLimiter(2, 5)
	leakyBucket = ratelimit.New(1)
}

func Allow() bool {
	return tokenBucket.Allow()
}

func Take() time.Time {
	return leakyBucket.Take()
}
