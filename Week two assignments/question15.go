/*
 Write a function GCD (Greatest Common Divisor) using recursion
in Go. Explain the base case and the recursive case in your solution.*/

package main

import "fmt"

func recursive(num1 int, num2 int, gcf int, iteration int) int {
	if iteration > num1 || iteration > num2 {
		return gcf

	}
	if num1%iteration == 0 && num2%iteration == 0 {
		gcf = iteration
	}
	return recursive(num1, num2, gcf, iteration+1)

}

func main() {

	varibe := recursive(6, 9, 1, 1)
	fmt.Println(varibe)

}







package main

import (
	"fmt"
)

func fibonachimomo() func(int) int {
	memo := make(map[int]int)

	var fab func(int) int
	fab = func(n int) int {
		if n <= 1 {
			return n
		}
		if val, found := memo[n]; found {
			return val
		}
		memo[n] = fab(n-1) + fab(n-2)
		return memo[n]
	}

	return fab
}

func main() {
	// Call fibonachimomo to get the Fibonacci function
	fib := fibonachimomo() // No arguments here

	// Use the returned function
	result := fib(5)
	fmt.Println(result) // Output: 5
}

