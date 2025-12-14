package crawler

import (
	"math"
	"math/rand"
	"time"
)

// Backoff implements an exponential backoff with optional jitter.
type Backoff struct {
	base   time.Duration
	max    time.Duration
	factor float64
	jitter float64 // 0..1 fraction for +/- jitter range
	tries  int
	r      *rand.Rand
}

// NewBackoff creates a backoff with the given parameters.
// factor < 1 is treated as 1 (no growth). jitter is in [0,1].
func NewBackoff(base, max time.Duration, factor, jitter float64) *Backoff {
	if factor < 1 {
		factor = 1
	}
	if jitter < 0 {
		jitter = 0
	}
	if jitter > 1 {
		jitter = 1
	}
	return &Backoff{
		base:   base,
		max:    max,
		factor: factor,
		jitter: jitter,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Next returns the next backoff duration.
func (b *Backoff) Next() time.Duration {
	// compute base * factor^tries
	pow := math.Pow(b.factor, float64(b.tries))
	d := time.Duration(float64(b.base) * pow)
	if d > b.max {
		d = b.max
	}
	// apply jitter if requested
	if b.jitter > 0 {
		// multiplier in [1-jitter, 1+jitter]
		min := 1 - b.jitter
		max := 1 + b.jitter
		m := min + b.r.Float64()*(max-min)
		d = time.Duration(float64(d) * m)
	}
	b.tries++
	return d
}

// Reset sets the backoff sequence back to initial.
func (b *Backoff) Reset() { b.tries = 0 }
