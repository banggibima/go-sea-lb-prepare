package missingnumber

import (
	"testing"
)

func TestFindMissingNumber(t *testing.T) {
	nums := []int{3, 0, 1}
	expected := 2
	got := FindMissingNumber(nums)

	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
