package structs

import "testing"

func checkExpected(t *testing.T, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f and wanted %.2f", got, want)
	}
}
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	checkExpected(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(2, 4)
	want := 8.0

	checkExpected(t, got, want)

}
