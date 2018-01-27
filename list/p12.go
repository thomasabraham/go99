package list

// Problem 12
// Decode a run-length encoded list.
// Eg: [[2, 1], [3, 2], 3, [2, 4], 5, 6, [3, 7]] => [1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 7, 7, 7]

func RunlengthDecode(input []interface{}) []int {
	output := make([]int, 0)

	for _, iv := range input {
		switch v := iv.(type) {
		case int:
			output = append(output, v)
			break
		case []int:
			for v[0] > 0 {
				output = append(output, v[1])
				v[0]--
			}
		default:
			panic("Unexpected type in input")
		}
	}
	return output
}
