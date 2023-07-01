package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say basic hello world ", func(t *testing.T) {
		got := Perimeter(4, 4)
		want := 16.0
		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want float64) {

	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
