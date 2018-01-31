package list

// Program 20
// Remove Kth element from the list
// Eg: [2, 3, 4, 5, 6, 7, 8], 3 => 5, [2, 3, 4, 6, 7, 8]

func RemoveKthFromList(input []int, K int) (value int, result []int) {
	output := make([]int, 0, len(input)-1)
	output = append(output, input[:K]...)
	output = append(output, input[K+1:]...)
	return input[K], output
}
