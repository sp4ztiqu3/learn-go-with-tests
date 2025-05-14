package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		want := 6

		if sum != want {
			t.Errorf("got %d, want %d, given %v", sum, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("sum some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 3, 6})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
