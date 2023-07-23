package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying named hello", func(t *testing.T) {
		got := Hello("Shubham")
		want := "Hello, Shubham"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying default hello when empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
