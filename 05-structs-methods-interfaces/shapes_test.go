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
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("expected %g, got %g", expected, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{8.0, 10.0}
		checkArea(t, rectangle, 80.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{8.0}
		checkArea(t, circle, 201.06192982974676)
	})
}
