package jumpgame

import (
	"testing"
)

func TestCanReach(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	expected := true
	got := CanReach(nums)

	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
