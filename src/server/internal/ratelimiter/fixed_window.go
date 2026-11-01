package ratelimiter

import (
	"sync"
	"time"
)

// FixedWindowRateLimiter allows up to `limit` requests per client within a
// fixed time window that starts at the client's first request.
type FixedWindowRateLimiter struct {
	mu      sync.Mutex
	clients map[string]int
	limit   int
	window  time.Duration
}

// NewFixedWindowLimiter creates a limiter allowing `limit` requests per `window`.
func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		clients: make(map[string]int),
		limit:   limit,
		window:  window,
	}
}

// Allow records a request from ip and reports whether it is permitted. When
// it is not, the returned duration is how long the client should wait.
//
// The check and the increment happen under a single lock so concurrent
// requests cannot exceed the limit.
func (rl *FixedWindowRateLimiter) Allow(ip string) (bool, time.Duration) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	count, exists := rl.clients[ip]
	if !exists {
		time.AfterFunc(rl.window, func() { rl.reset(ip) })
	}

	if count >= rl.limit {
		return false, rl.window
	}

	rl.clients[ip] = count + 1
	return true, 0
}

func (rl *FixedWindowRateLimiter) reset(ip string) {
	rl.mu.Lock()
	delete(rl.clients, ip)
	rl.mu.Unlock()
}
