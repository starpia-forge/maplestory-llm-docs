package crawler

import (
	"math"
	"testing"
	"time"
)

func TestBackoff_NoJitter_ExponentiallyIncreasesAndCaps(t *testing.T) {
	b := NewBackoff(100*time.Millisecond, 1*time.Second, 2.0, 0.0)

	got := []time.Duration{b.Next(), b.Next(), b.Next(), b.Next(), b.Next(), b.Next()}
	want := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 400 * time.Millisecond, 800 * time.Millisecond, 1 * time.Second, 1 * time.Second}
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %d want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("index %d: got %v want %v", i, got[i], want[i])
		}
	}
}

func TestBackoff_FactorLessThanOne_TreatedAsOne(t *testing.T) {
	b := NewBackoff(150*time.Millisecond, 1*time.Second, 0.5, 0.0)
	d1 := b.Next()
	d2 := b.Next()
	if d1 != 150*time.Millisecond || d2 != 150*time.Millisecond {
		t.Fatalf("expected constant base when factor<1: got %v %v", d1, d2)
	}
}

func TestBackoff_JitterWithinBounds(t *testing.T) {
	base := 200 * time.Millisecond
	b := NewBackoff(base, 10*time.Second, 1.0, 0.25) // ±25%
	d := b.Next()
	min := time.Duration(float64(base)*(1-0.25) - 1) // tolerance for rounding
	max := time.Duration(float64(base)*(1+0.25) + 1)
	if d < min || d > max {
		t.Fatalf("jittered duration %v not within [%v, %v]", d, min, max)
	}
}

func TestBackoff_Reset(t *testing.T) {
	b := NewBackoff(100*time.Millisecond, 1*time.Second, 3.0, 0.0)
	_ = b.Next()
	_ = b.Next()
	b.Reset()
	d := b.Next()
	if math.Abs(float64(d-100*time.Millisecond)) > 0 {
		t.Fatalf("expected base after reset, got %v", d)
	}
}
