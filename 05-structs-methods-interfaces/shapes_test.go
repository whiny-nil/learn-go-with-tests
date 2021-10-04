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
			t.Errorf("%#v expected %g, got %g", shape, expected, got)
		}
	}

	tests := []struct {
		Name     string
		Shape    Shape
		Expected float64
	}{
		{"Rectangle", Rectangle{Width: 8.0, Height: 10.0}, 80.0},
		{"Circle", Circle{Radius: 8.0}, 201.06192982974676},
		{"Triangle", Triangle{Base: 12.0, Height: 8.0}, 48.0},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			checkArea(t, test.Shape, test.Expected)
		})
	}
}
