/*
	Create a function FilterSlice that takes a slice of integers and

a predicate function (a function that returns a boolean). The function
should return a new slice containing only the elements that satisfy the
predicate function.
*/
package main

import (
	"fmt"
)

func Filterslice(input1 []int, operation func(int) bool) []int {
	var final []int

	for _, item := range input1 {
		if operation(item) {
			final = append(final, item)
		}
	}
	return final

}

func predicate(n int) bool {
	if n > 0 {
		return true
	}

	return false

}

func main() {
	numbers := []int{-1, 2, 3, 4, 5, 6}
	final := Filterslice(numbers, predicate)
	fmt.Println(final)
}
