package list

// Problem 25
// Generate random permutation of elements of a list
// Eg: [1,2,3,4,5] => [3, 2, 5, 4, 1]

func RandomizeList(input []int) []int {
	return RandomN(input, len(input))
}
