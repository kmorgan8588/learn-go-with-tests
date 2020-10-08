package property

import "testing"

func TestRoman(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"converts 1", 1, "I"},
		{"converts 2", 2, "II"},
		{"converts 3", 3, "III"},
		{"converts 4", 4, "IV"},
		{"converts 5", 5, "V"},
		{"converts 6", 6, "VI"},
		{"converts 6", 6, "VI"},
		{"converts 7", 7, "VII"},
		{"converts 8", 8, "VIII"},
		{"converts 9", 9, "IX"},
		{"converts 10", 10, "X"},
		{"converts 14", 14, "XIV"},
		{"converts 18", 18, "XVIII"},
		{"converts 20", 20, "XX"},
		{"converts 39", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Want {
				t.Errorf(`got "%s", want "%s"`, got, test.Want)
			}
		})
	}
}
