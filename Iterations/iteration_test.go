package main

import "testing"

func TestRepeatCharacter(t *testing.T) {

	got := Repeat("A")
	want := "AAAAA"

	AssertCorrectMessage(t, want, got)
}

func AssertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
