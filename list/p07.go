package main

import "fmt"

// Problem 7
// Flatten a nested list structure
// Eg: [1 [2 [3 4] 5]] => [1 2 3 4 5]
// Currently to keep things simple, we'll use a hardcoded slice

func main() {
	fmt.Println("Problem-7: Flatten a nested list")
	inputSlice := []interface{}{
		[]interface{}{1, 2},
		3,
		[]interface{}{
			5,
			[]interface{}{6, 7},
		},
		8,
		9,
	}
	fmt.Println("The intput slice is", inputSlice)
	fmt.Println("The flattened slice is", flatten(inputSlice))
}

func flatten(inputSlice []interface{}) []int {
	output := []int{}
	for _, iv := range inputSlice {
		switch v := iv.(type) {
		case int:
			output = append(output, v)
			break
		case []interface{}:
			output = append(output, flatten(v)...)
		default:
			// This will probably be an error case.
			// Ignoring it as our input is fixed
		}
	}
	return output
}
