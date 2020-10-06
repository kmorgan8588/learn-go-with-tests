package structs

import (
	"math"
	"testing"
)

func checkExpected(t *testing.T, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %g and wanted %g", got, want)
	}
}
func TestRectanglePerimeter(t *testing.T) {
	rectangle := Rectangle{Length: 10.0, Width: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	checkExpected(t, got, want)
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		checkExpected(t, got, want)
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Length: 2.0, Width: 4.0}
		checkArea(t, rectangle, 8.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := math.Pi * 100

		checkArea(t, circle, want)
	})
}
