package list

// Problem 17
// Split the given list into two parts, length of first list is given
// Eg: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] 4 => [1, 2, 3, 4], [5, 6, 7, 8, 9, 10]

func SplitList(input []int, pivot int) ([]int, []int) {
	if pivot < 0 {
		pivot = 0
	}
	if pivot > len(input) {
		pivot = len(input)
	}
	a := input[:pivot]
	b := input[pivot:]
	return a, b
}
