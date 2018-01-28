package list

// Problem 11
// Modified run-length encoding.
// Eg: [1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 7, 7, 7] => [[2, 1], [3, 2], 3, [2, 4], 5, 6, [3, 7]]

func RunlengthEncode2(input []int) []interface{} {

	encodedList := RunlengthEncode(input)

	output := make([]interface{}, 0)

	for _, v := range encodedList {
		if v[0] > 1 {
			output = append(output, v)
		} else {
			output = append(output, v[1])
		}
	}

	return output
}
