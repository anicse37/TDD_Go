package iteration_test

import (
	iteration "aniket/Iteration"
	"testing"
)

func TestFor(t *testing.T) {
	got := iteration.For("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q | want %q", got, want)
	}
}
func TestRepeat(t *testing.T) {
	got := iteration.Repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q | want %q", got, want)
	}
}

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.For("a")
	}
}
