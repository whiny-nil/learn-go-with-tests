package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Marc")

	got := buffer.String()
	want := "Hello, Marc\n"

	if got != want {
		t.Errorf("expected \"%s\", got \"%s\"", want, got)
	}
}
