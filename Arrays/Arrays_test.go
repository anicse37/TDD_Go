package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/anicse37/TDD_Go/Arrays"
)

func TestArray(t *testing.T) {
	t.Run("Some Error Occured", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		got := arrays.Sum(arr)
		want := 15
		if got != want {
			t.Errorf("got %d | want %d", got, want)
		}
	})
}

func TestMultipleArray(t *testing.T) {

	got := arrays.MultipleArray([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
