package list

// Problem 24
// Randomly choose N numbers from given range
// Eg: 3, 7, 2 => [6, 5]

func RandomInRange(start int, end int, count int) []int {
	rangeList := RangeAsList(start, end)
	resultList := []int{}
	if count <= len(rangeList) && count >= 0 {
		resultList = RandomN(rangeList, count)
	}
	return resultList
}
