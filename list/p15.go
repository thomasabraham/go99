package list

// Problem 15
// Duplicate every element in the input list N times
// Eg: [1 2 3 4 5 2], 2 => [1 1 2 2 3 3 4 4 5 5 2 2]

func DuplicateElementsN(input []int, N int) []int {
	if N <= 0 {
		panic("N should not be <=0")
	}
	output := make([]int, 0, len(input)*N)
	for _, v := range input {
		for i := 0; i < N; i++ {
			output = append(output, v)
		}
	}
	return output
}
