package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("Enter a list of numbers.")
	fmt.Println("Enter one number per line.")
	fmt.Println("Enter the number of numbers in the list in the first line, followed by actual numbers in subsequent lines.")

	var count int
	fmt.Scanf("%d", &count)

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	fmt.Println("You have entered", count, "numbers.")
	fmt.Println("You entered", numbers)
	fmt.Println("Choose an index between 0 and", count-1)
	var index int
	fmt.Scanf("%d", &index)

	numberAtIndex := getValue(numbers, index)
	fmt.Println("The number at index", index, "is", numberAtIndex)
}

func getValue(list []int, index int) int {
	if 0 <= index && index < len(list) {
		return list[index]
	} else {
		panic("Given index is out of range 0 and " + strconv.Itoa(len(list)))
	}
}
