/*
Implement a simple memoization function in Go called
MemoizeFib for calculating Fibonacci numbers. It should use a map
to store already computed Fibonacci values and reduce redundant
calculations.
*/
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
	fib := fibonachimomo()
	result := fib(5)
	fmt.Println(result)

}
