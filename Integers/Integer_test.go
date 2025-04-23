package integers_test

import (
	"testing"

	integers "github.com/anicse37/TDD_Go/Integers"
)

func TestInteger(t *testing.T) {
	got := integers.Add(1, 3)
	want := 4
	if got != want {
		t.Errorf("got %d | want %d", got, want)
	}
}
