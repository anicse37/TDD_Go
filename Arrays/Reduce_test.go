package arrays_test

import (
	"testing"

	arrays "github.com/anicse37/TDD_Go/Arrays"
)

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, arrays.Reduce([]int{1, 2, 3}, multiply, 1), 6)

	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, arrays.Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}
