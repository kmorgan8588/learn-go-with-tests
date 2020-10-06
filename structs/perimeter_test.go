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
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		checkExpected(t, got, tt.want)
	}

}
