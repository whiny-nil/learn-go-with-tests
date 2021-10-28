package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	expected := `3
2
1
Go!`

	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("expected 4 calls, got %d", spySleeper.Calls)
	}
}
