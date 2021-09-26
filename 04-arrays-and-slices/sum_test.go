package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("sums a slice numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		expected := 6
		got := Sum(numbers)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
