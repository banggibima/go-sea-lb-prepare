package besttimetobuyandsellstock

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	got := MaxProfit(prices)

	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
