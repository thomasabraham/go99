package list

import (
	"sort"
	"strconv"
	"testing"
)

func Test26(t *testing.T) {
	cases := []struct {
		input  []int
		K      int
		output [][]int
	}{
		{[]int{}, 0, [][]int{}},
		{[]int{}, -1, [][]int{}},
		{[]int{}, 1, [][]int{}},
		{[]int{1}, 0, [][]int{}},
		{[]int{1}, -1, [][]int{}},
		{[]int{1}, 2, [][]int{}},
		{[]int{1}, 1, [][]int{[]int{1}}},
		{[]int{1, 2}, 0, [][]int{}},
		{[]int{1, 2}, 1, [][]int{[]int{1}, []int{2}}},
		{[]int{1, 2}, 2, [][]int{[]int{1, 2}}},
		{[]int{1, 2}, 3, [][]int{}},
		{[]int{1, 2, 3}, 2, [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}},
		{[]int{1, 2, 3, 4}, 3, [][]int{[]int{1, 2, 3}, []int{1, 2, 4}, []int{1, 3, 4}, []int{2, 3, 4}}},
	}
	for _, c := range cases {
		result := ListCombinations(c.input, c.K)
		if !checkResultList(c.output, result) {
			t.Errorf("Actual result %v don't match expected result %v", result, c.output)
		}
	}
}

func checkResultList(expected [][]int, actual [][]int) bool {

	if len(actual) != len(expected) {
		return false
	}

	// Convert each subset to string after sorting
	expectedList := make([]string, len(expected))
	for k, v := range expected {
		expectedList[k] = sortConvertIntListToString(v)
	}
	actualList := make([]string, len(actual))
	for k, v := range actual {
		actualList[k] = sortConvertIntListToString(v)
	}

	expectedStr := sortConvertStringListToString(expectedList)
	actualStr := sortConvertStringListToString(actualList)

	return expectedStr == actualStr
}

// Sort a list of integers and convert it to a string
func sortConvertIntListToString(input []int) string {
	sort.Ints(input)
	output := ""
	for _, v := range input {
		output = output + strconv.Itoa(v)
	}
	return output
}

// Sort a list of Strings and convert it to a single string
func sortConvertStringListToString(input []string) string {
	sort.Strings(input)
	output := ""
	for _, v := range input {
		output = output + v
	}
	return output
}
