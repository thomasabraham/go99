package list

// Problem 7
// Flatten a nested list structure
// Eg: [1 [2 [3 4] 5]] => [1 2 3 4 5]
// Currently to keep things simple, we'll use a hardcoded slice

func Flatten(inputSlice []interface{}) []int {
	output := []int{}
	for _, iv := range inputSlice {
		switch v := iv.(type) {
		case int:
			output = append(output, v)
			break
		case []interface{}:
			output = append(output, Flatten(v)...)
		default:
			// This will probably be an error case.
			// Ignoring it as our input is fixed
		}
	}
	return output
}
