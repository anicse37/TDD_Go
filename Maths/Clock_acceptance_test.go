package clockFace_test

import (
	"testing"
	"time"

	clockFace "github.com/anicse37/TDD_Go/Maths"
)

func TestSecondHandAtMidNight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := clockFace.Point{X: 150, Y: 150 + 90}
	got := clockFace.SecondHand(tm)
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}
