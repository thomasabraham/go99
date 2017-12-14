package list

// Problem 6
// Is the list palindrome
func IsPalindRome(numbers []int) bool {
	var palindRome bool = true
	length := len(numbers)
	for i := 0; i < length/2; i++ {
		if numbers[i] != numbers[length-1-i] {
			palindRome = false
			break
		}
	}
	return palindRome
}
