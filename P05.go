package main

import "fmt"

// Problem 5
// Reverse a list

func main() {
	fmt.Println("Problem-5: Reverse a list")
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
	fmt.Println("The numbers you entered are", numbers)
	fmt.Println("The numbers you entered in reverse are", reverse(numbers))
}

func reverse(numbers []int) []int {
	reverse := []int{}
	length := len(numbers)
	for i := range numbers {
		reverse = append(reverse, numbers[length-i-1]);
	}
	return reverse
}
