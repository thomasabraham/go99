package list

// Problem 11
// Modified run-length encoding.
// Eg: [1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 7, 7, 7] => [[2, 1], [3, 2], 3, [2, 4], 5, 6, [3, 7]]

func RunlengthEncode2(input []int) []interface{} {
	output := make([]interface{}, 0)

	if len(input) == 0 {
		return output
	}

	currentNumber := input[0]
	currentNumberCount := 1
	for _, v := range input[1:] {
		if currentNumber == v {
			currentNumberCount++
			continue
		}
		if currentNumberCount > 1 {
			output = append(output, []int{currentNumberCount, currentNumber})
		} else {
			output = append(output, currentNumber)
		}
		currentNumber = v
		currentNumberCount = 1
	}
	if currentNumberCount > 1 {
		output = append(output, []int{currentNumberCount, currentNumber})
	} else {
		output = append(output, currentNumber)
	}
	return output
}
