package slidingwindow

import "time"

type RateLimiter interface {
	SetRate(rate float64)
	SetWindow(window time.Duration)
	Allow() bool
}
