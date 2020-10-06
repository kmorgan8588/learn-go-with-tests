package iteration

import "strings"

//Repeat takes in a character and returns a string with multiple copies of the character
func Repeat(character string, repeatCount int) string {

	return strings.Repeat(character, repeatCount)
}
