package tokenbucket

import (
	"sync"
	"time"
)

type TokenContext struct {
	tokenBuckets map[string]*TokenBucket
	capacity     float64
	refillRate   float64
}

type TokenBucket struct {
	capacity   float64
	tokens     float64
	refillRate float64
	lastRefill time.Time
	mu         sync.Mutex
}

var tokenctx *TokenContext

func InitRateLimiter(capacity float64, refillRate float64) {
	tokenctx = &TokenContext{
		tokenBuckets: make(map[string]*TokenBucket),
		capacity:     capacity,
		refillRate:   refillRate,
	}
}

func NewTokenBucket(ip string) *TokenBucket {
	if tokenctx == nil {
		InitRateLimiter(5, 2)
	}

	if bucket, exists := tokenctx.tokenBuckets[ip]; exists {
		return bucket
	}

	return &TokenBucket{
		capacity:   tokenctx.capacity,
		tokens:     tokenctx.capacity,
		refillRate: tokenctx.refillRate,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()

	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastRefill = now
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1 {
		tb.tokens -= 1
		return true
	}

	return false
}
