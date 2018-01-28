package list

// Problem 18
// Extract a slice from a list
// Eg: [1, 1, 2, 3, 3, 5, 6, 7] 3, 5 => [2, 3, 3]

func SliceList(input []int, start int, end int) []int {
	if len(input) < 1 {
		return []int{}
	}
	if start > end {
		temp := start
		start = end
		end = temp
	}

	if start > len(input) {
		start = len(input) + 1
	}
	if start < 1 {
		start = 1
	}
	if end < 0 {
		end = 0
	}
	if end > len(input) {
		end = len(input)
	}
	return input[start-1 : end]
}
