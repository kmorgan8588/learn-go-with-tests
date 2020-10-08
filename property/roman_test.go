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
