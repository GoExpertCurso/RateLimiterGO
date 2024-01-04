package slidingwindow

import "time"

type SlidingwindowRateLimiter struct {
	sw Slidingwindow
}

func NewSlidingwindowRateLimiter(rate float64, window time.Duration) *SlidingwindowRateLimiter {
	sw := Slidingwindow{
		count:   int(rate * float64(window/time.Second)),
		window:  window,
		history: []int{},
	}
	return &SlidingwindowRateLimiter{
		sw,
	}
}

func (rl *SlidingwindowRateLimiter) SetRate(rate float64) {
	rl.sw.SetRate(rate)
}

func (rl *SlidingwindowRateLimiter) SetWindow(window time.Duration) {
	rl.sw.SetWindow(window)
}

func (rl *SlidingwindowRateLimiter) Allow() bool {
	return rl.sw.Allow()
}
