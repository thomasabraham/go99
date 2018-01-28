package list

// Problem 16
// Drop every Nth element from a list
// Eg: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 3 => [1, 2, 4, 5, 7, 8, 10]

func DropEveryNth(input []int, N int) []int {
	if N <= 1 {
		panic("N should be greater than 1")
	}
	output := make([]int, 0, len(input)-(len(input)/N))
	counter := 1
	for _, v := range input {
		if counter == N {
			counter = 1
		} else {
			counter++
			output = append(output, v)
		}
	}
	return output
}
