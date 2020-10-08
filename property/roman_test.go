package property

import "testing"

func TestRoman(t *testing.T) {
	t.Run("converts 1", func(t *testing.T) {
		want := "I"
		got := ConvertToRoman(1)

		if got != want {
			t.Errorf(`got "%s", want "%s"`, got, want)
		}
	})

	t.Run("converts 2", func(t *testing.T) {
		want := "II"
		got := ConvertToRoman(2)

		if got != want {
			t.Errorf(`got "%s", want "%s"`, got, want)
		}
	})

}
