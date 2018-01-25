package list

// Problem 10
// Run-length encoding of a list
// Eg: [1, 1, 1, 2, 2, 3, 3, 3, 3, 4] => [[3, 1], [2, 2], [4, 3], [1, 4]]

func RunlengthEncode(input []int) [][]int {
	output := make([][]int, 0)

	if len(input) == 0 {
		return output
	}

	currentNumber := input[0]
	currentNumberCount := 1
	for i := 1; i < len(input); i++ {
		if currentNumber == input[i] {
			currentNumberCount++
			continue
		}
		output = append(output, []int{currentNumberCount, currentNumber})
		currentNumber = input[i]
		currentNumberCount = 1
	}
	output = append(output, []int{currentNumberCount, currentNumber})
	return output
}
