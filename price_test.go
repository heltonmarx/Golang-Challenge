package sample1

import (
	"testing"
	"time"
)

func TestPrice_IsValid(t *testing.T) {
	tests := []struct {
		price    float64
		maxAge   time.Duration
		wait     time.Duration
		expected bool
	}{
		{
			price:    5,
			maxAge:   time.Millisecond * 100,
			wait:     time.Millisecond * 50, // 50pct
			expected: true,
		},
		{
			price:    5,
			maxAge:   time.Millisecond * 100,
			wait:     time.Millisecond * 70, // 70pct
			expected: true,
		},
		{
			price:    5,
			maxAge:   time.Millisecond * 100,
			wait:     time.Millisecond * 120, // 120pct
			expected: false,
		},
	}

	for _, test := range tests {
		price := NewPrice(test.price)
		time.Sleep(test.wait)
		assertBool(t, test.expected, price.IsValid(test.maxAge), "wrong price retrieved validation")
		assertFloat(t, test.price, price.Value(), "wrong price returned")
	}
}
