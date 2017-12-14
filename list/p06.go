package main

import "fmt"

// Problem 6
// Is the list palindrome

func main() {
	fmt.Println("Problem-6: Palindrome or not.")
	numbers := readList()
	fmt.Println("The numbers you entered are", numbers)
	if isPalindRome(numbers) {
		fmt.Println("The numbers you entered are palindrome")
	} else {
		fmt.Println("The numbers you entered are not palindrome")
	}
}

func isPalindRome(numbers []int) bool {
	var palindRome bool = true
	length := len(numbers)
	for i := 0; i < length/2; i++ {
		if numbers[i] != numbers[length-1-i] {
			palindRome = false
			break
		}
	}
	return palindRome
}

func readList() []int {
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
	return numbers
}
