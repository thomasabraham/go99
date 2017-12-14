package main

import "fmt"

func main() {
	fmt.Println("Enter some numbers")
	exit := false
	numbers := []int{}
	for exit == false {
		var number int
		n, err := fmt.Scanf("%d", &number)
		if n == 0 || err != nil {
			exit = true
		} else {
			numbers = append(numbers, number)
		}
	}
	fmt.Printf("You have entered %d numbers.\n", count(numbers))
	fmt.Println("The numbers you entered are", numbers)
}
func count(numbers []int) int {
	return len(numbers)
}
