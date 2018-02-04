package list

// Problem 23
// Randomly choose N items from given list
// Eg: [2, 3, 4, 5, 6], 3 => [5, 2, 4]

import (
	"math/rand"
)

func RandomN(inputList []int, length int) []int {

	if length > len(inputList) {
		panic("InputList is too small")
	}

	if length < 0 {
		panic("Length cannot be negative")
	}

	result := make([]int, length)

	for k, _ := range result {
		index := rand.Intn(len(inputList))
		result[k] = inputList[index]
		_, inputList = RemoveKthFromList(inputList, index)
	}
	return result
}
