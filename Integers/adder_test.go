package main

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	AssertCorrectMessage(t, want, got)
}

func AssertCorrectMessage(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}
