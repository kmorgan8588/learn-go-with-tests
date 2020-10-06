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

func TestRectangleArea(t *testing.T) {
	rectangle := Rectangle{Length: 2.0, Width: 4.0}
	got := Area(rectangle)
	want := 8.0

	checkExpected(t, got, want)
}

func TestCircleArea(t *testing.T) {
	circle := Circle{10}
	got := Area(circle)
	want := math.Pi * 10

	checkExpected(t, got, want)
}
