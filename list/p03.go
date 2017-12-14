package list

import "strconv"

func GetValueAt(list []int, index int) int {
	if 0 <= index && index < len(list) {
		return list[index]
	} else {
		panic("Given index is out of range 0 and " + strconv.Itoa(len(list)))
	}
}
