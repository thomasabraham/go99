package list

// Problem 9
// Pack consecutive duplicates in a list to sublist
// Eg: [1 2 2 2 3 3 4 5 5 5 5 5 2 2] => [[1] [2 2 2] [3 3] [4] [5 5 5 5 5] [2 2]]

func PackDuplicates(input []int) [][]int {

	output := make([][]int, 0)

	if len(input) == 0 {
		return output
	} else {
		var firstSub []int
		firstSub = append(firstSub, input[0])
		output = append(output, firstSub)
	}

	for _, v := range input[1:] {
		if output[len(output)-1][0] == v {
			output[len(output)-1] = append(output[len(output)-1], v)
		} else {
			newSub := make([]int, 1)
			newSub[0] = v
			output = append(output, newSub)
		}
	}
	return output
}
