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
	got := SumAll([]int{1, 2}, []int{9, 0})

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("it adds the tails of given slices", func(t *testing.T) {
		expected := []int{2, 0}
		got := SumAllTails([]int{1, 2}, []int{9, 0})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("it handles empty slices", func(t *testing.T) {
		expected := []int{2, 0}
		got := SumAllTails([]int{1, 2}, []int{})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
