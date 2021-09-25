package iteration

import "testing"

func TestRepeat(t *testing.T) {
	expected := "aaaaaa"
	got := Repeat("a", 6)

	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
