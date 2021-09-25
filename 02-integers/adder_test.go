package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	expected := 4
	got := Adder(1, 3)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func ExampleAdder() {
	sum := Adder(6, 7)
	fmt.Println(sum)
	// Output: 13
}
