package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d and wanted %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d and wanted %d", got, want)

		}
	}
	t.Run("multiple collections", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		checkSums(t, got, want)
	})

	t.Run("empty collection", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
