package list

func Last(list []int) int {
	if len(list) > 0 {
		return list[len(list)-1]
	} else {
		panic("cannot get last item from empty list")
	}
}
