// Package ratelimiter provides per-client request throttling.
package ratelimiter

import "time"

type Limiter interface {
	Allow(ip string) (bool, time.Duration)
}

type Config struct {
	RequestsPerTimeFrame int
	TimeFrame            time.Duration
	Enabled              bool
}
