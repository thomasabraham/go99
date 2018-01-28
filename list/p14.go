package list

// Problem 14
// Duplicate every element in the input list
// Eg: [1 2 3 4 5 2] => [1 1 2 2 3 3 4 4 5 5 2 2]

func DuplicateElements(input []int) []int {
	output := make([]int, 0, len(input)*2)
	for _, v := range input {
		output = append(output, v)
		output = append(output, v)
	}
	return output
}
