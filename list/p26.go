package list

// Problem 26
// Find all the combinations of length K from the given list
// Eg: [1, 2, 3, 4, 5], 4 => [[1, 2, 3, 4], [1, 2, 3, 5], [1, 2, 4, 5], [1, 3, 4, 5], [2, 3, 4, 5]]

func ListCombinations(input []int, K int) [][]int {
	output := make([][]int, 0)

	if K == 1 {
		for _, v := range input {
			output = append(output, []int{v})
		}
	}

	for i, v := range input {
		smallerSubsets := ListCombinations(input[i+1:], K-1)
		for _, smallerSubset := range smallerSubsets {
			output = append(output, append(smallerSubset, v))
		}
	}

	return output
}
