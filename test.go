package main

import (
	"errors"
	"fmt"
)

// factorial calculates the factorial of a given number.
// It returns an error if the number is negative.
func factorial(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("cannot calculate factorial of a negative number")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func main() {
	var num int
	fmt.Print("Enter a number to calculate its factorial: ")
	fmt.Scan(&num)

	result, err := factorial(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Factorial of %d is %d\n", num, result)
	}
}
