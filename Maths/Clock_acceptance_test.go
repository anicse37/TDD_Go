package clockFace_test

import (
	clockFace "aniket/Maths"
	"testing"
	"time"
)

func TestSecondHandAtMidNight(t *testing.T) {
	tm := time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)
	want := clockFace.Point{X: 150, Y: 150 - 90}
	got := clockFace.SecodHand(tm)
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}
