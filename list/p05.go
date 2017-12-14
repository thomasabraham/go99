package list

// Problem 5
// Reverse a list
func Reverse(numbers []int) []int {
	reverse := []int{}
	length := len(numbers)
	for i := range numbers {
		reverse = append(reverse, numbers[length-i-1])
	}
	return reverse
}
