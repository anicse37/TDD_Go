package main_test

import (
	custom "aniket/HelloWorld"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := custom.Hello("Chris")
		want := "Hello, Chris"
		printMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := custom.Hello("")
		want := "Hello, World"
		printMessage(t, got, want)

	})

}
func printMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
