package property

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(num int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}
	return result.String()
}
