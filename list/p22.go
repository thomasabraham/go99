package list

// Problem 22
// List of all integers in the given range
// Eg: 2, 5 => [2, 3, 4, 5]

func RangeAsList(start int, end int) []int {
	if start > end {
		return []int{}
	}
	result := make([]int, end-start+1)
	for k, _ := range result {
		result[k] = start
		start++
	}
	return result
}
