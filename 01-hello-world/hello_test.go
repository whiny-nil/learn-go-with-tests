package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Marc")
	want := "Hello, Marc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
