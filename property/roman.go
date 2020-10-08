package property

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}
type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
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

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		//look ahead to next symbol if we can and the current symbol is base 10 ()
		if i+1 < len(roman) && symbol == 'I' {
			nextSymbol := roman[i+1]

			//build two character string
			potentialNumber := string([]byte{symbol, nextSymbol})

			//get value of the potentialNumber
			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++ //used the character, so we iterate past it
			} else {
				total++
			}
		} else {
			total++
		}
	}
	return total
}
