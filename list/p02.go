package list

func LastButOne(list []int) int {
	if len(list) > 1 {
		return list[len(list)-2]
	} else {
		panic("Cannot get second last item from list with less than 2 items")
	}
}
