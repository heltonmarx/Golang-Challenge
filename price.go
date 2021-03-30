package sample1

import (
	"time"
)

// Price represents the price and the timestamp of
// last time of value was retrieved.
type Price struct {
	value    float64
	lastUsed time.Time
}

// NewPrice returns a new price with a new timestamp.
func NewPrice(price float64) Price {
	return Price{
		value:    price,
		lastUsed: time.Now(),
	}
}

// IsValid returns true if the price was retrieved less than maxAge.
func (p *Price) IsValid(maxAge time.Duration) bool {
	now := time.Now()
	if now.Sub(p.lastUsed) > maxAge {
		return false
	}
	p.lastUsed = now
	return true
}

// Value returns the price value.
func (p *Price) Value() float64 {
	return p.value
}
