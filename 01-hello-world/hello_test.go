package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, got string) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, Marc"
		got := Hello("Marc")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("saying hello to everyone", func(t *testing.T) {
		expected := "Hello, World"
		got := Hello("")
		assertCorrectMessage(t, expected, got)
	})
}
