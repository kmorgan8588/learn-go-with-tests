package structs

import (
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

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Length: 12, Width: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
