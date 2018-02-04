package list

// Problem 21
// Insert a given element into the given position of a list
// Eg: [1,2,3,4], 99, 2 => [1, 2, 99, 3, 4]

func ListInsert(inputList []int, value int, index int) []int {

	output := make([]int, 0, len(inputList)+1)

	if index < 0 || index > len(inputList) {
		output = append(output, inputList...)
	} else {
		output = append(output, inputList[:index]...)
		output = append(output, value)
		output = append(output, inputList[index:]...)
	}

	return output
}
