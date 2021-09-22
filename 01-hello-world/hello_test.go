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
		got := Hello("Marc", "en")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("saying hello to everyone", func(t *testing.T) {
		expected := "Hello, World"
		got := Hello("", "en")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		expected := "Bonjour, Marc"
		got := Hello("Marc", "fr")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("saying hello to everyone in French", func(t *testing.T) {
		expected := "Bonjour, World"
		got := Hello("", "fr")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		expected := "Hola, Marc"
		got := Hello("Marc", "es")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("saying hello to everyone in Spanish", func(t *testing.T) {
		expected := "Hola, World"
		got := Hello("", "es")
		assertCorrectMessage(t, expected, got)
	})
}
