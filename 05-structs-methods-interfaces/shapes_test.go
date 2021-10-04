package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{8.0, 10.0}
	expected := 36.0
	got := Perimeter(rectangle)

	if got != expected {
		t.Errorf("expected %g, got %g", expected, got)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{8.0, 10.0}
		expected := 80.0
		got := rectangle.Area()

		if got != expected {
			t.Errorf("expected %g, got %g", expected, got)
		}

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{8.0}
		expected := 201.06192982974676
		got := circle.Area()

		if got != expected {
			t.Errorf("expected %g, got %g", expected, got)
		}

	})
}
