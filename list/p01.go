package main

import "fmt"

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

	lastNumber := last(numbers)
	fmt.Println("Last number is", lastNumber)
}

func last(list []int) int {
	if len(list) > 0 {
		return list[len(list)-1]
	} else {
		panic("cannot get last item from empty list")
	}
}
