package slices

//Sum takes an array of 5 integers and finds the sum
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

//SumAll takes a variable number of integer slices, and returns a slice of their sums
func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, set := range numbers {
		if len(set) == 0 {
			sums = append(sums, 0)
		} else {
			tail := set[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
