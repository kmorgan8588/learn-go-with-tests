package structs

import "testing"

func checkExpected(t *testing.T, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f and wanted %.2f", got, want)
	}
}
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Length: 10.0, Width: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	checkExpected(t, got, want)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{Length: 2.0, Width: 4.0}
	got := Area(rectangle)
	want := 8.0

	checkExpected(t, got, want)

}
