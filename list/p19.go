package list

// Problem 19
// Rotate a list N places to the left
// Eg: [1, 2, 3, 4, 5, 6, 7], 3 => [4, 5, 6, 7, 1, 2, 3]

func RotateListLeft(input []int, N int) []int {

	if N < 0 {
		panic("N cannot be less than zero")
	}

	length := len(input)
	output := make([]int, length)

	if length == 0 {
		return output
	}

	N = N % length
	for i, v := range input {
		output[(length+i-N)%length] = v
	}

	return output
}
