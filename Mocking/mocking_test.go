package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer, 3)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("Got %v || Want %v", got, want)
	}

}

func TestSleepCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	SleepCountdown(buffer, 3)
	got := buffer.String()
	want := `Write
Sleep
Write
Sleep
Write
Sleep
`
	if got != want {
		t.Errorf("Got %v || Want %v", got, want)
	}
}
