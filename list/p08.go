package list

// Problem 8
// Remove consecutive duplicates from a list
// Eg: [1 2 2 2 3 3 4 5 5 5 5 5 2 2] => [1 2 3 4 5 2]

func RemoveDuplicates(input []int) []int {
	if len(input) <= 0 {
		return input
	}
	output := make([]int, 0, len(input))
	output = append(output, input[0])
	for _, x := range input {
		if x != output[len(output)-1] {
			output = append(output, x)
		}
	}
	return output
}
