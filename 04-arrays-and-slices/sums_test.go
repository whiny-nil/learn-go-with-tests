package sums

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	expected := 6
	got := Sum(numbers)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumAll(t *testing.T) {
	expected := []int{3, 9}
	got := SumAll([]int{1, 2}, []int{0, 9})

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
