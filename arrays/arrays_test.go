package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	want := 15

	if sum != want {
		t.Errorf("got %d, want %d, given %v", sum, want, numbers)
	}
}
