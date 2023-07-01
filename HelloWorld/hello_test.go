package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say basic hello world ", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		AssertCorrectMessage(t, got, want)
	})
	t.Run("Check for empty string that is passed, ", func(t *testing.T) {
		got := Hello("")
		want := "Hello, sWorld"
		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {

	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
