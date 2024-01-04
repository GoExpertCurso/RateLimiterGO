package slidingwindow

import (
	"sync"
	"time"
)

type Slidingwindow struct {
	mu      sync.Mutex
	count   int
	window  time.Duration
	history []int
}

func (sw *Slidingwindow) increment() {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	sw.count++
	sw.history = append(sw.history, sw.count)
}

func (sw *Slidingwindow) removedExpired(now time.Time) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	for len(sw.history) > 0 && now.Sub(time.Unix(0, int64(sw.history[0]))) >= sw.window {
		sw.history = sw.history[1:]
	}
}

func (sw *Slidingwindow) countRequests(now time.Time) int {
	sw.removedExpired(now)
	return len(sw.history)
}

func (sw *Slidingwindow) Allow() bool {
	now := time.Now()
	count := sw.countRequests(now)
	if count > sw.count {
		return false
	}
	sw.increment()
	return true
}

func (sw *Slidingwindow) SetRate(rate float64) {
	sw.count = int(rate * float64(sw.window/time.Second))
}

func (sw *Slidingwindow) SetWindow(window time.Duration) {
	sw.window = window
}
