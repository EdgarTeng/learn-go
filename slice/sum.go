package slice

// Sum is use for add all items together
func Sum(numbers []int) int {
	var result int
	for _, item := range numbers {
		result += item
	}
	return result
}
