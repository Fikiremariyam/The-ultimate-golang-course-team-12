/*
Write a function RemoveDuplicates that takes a slice of integers
and returns a new slice without any duplicates. Ensure that the order
of the original elements is maintained.
*/

package main

import (
	"fmt"
	"slices"
)

func RemoveDuplicates(numbers []int) []int {
	var final []int
	for _, value := range numbers {
		if !slices.Contains(final, value) {
			final = append(final, value)
		}
	}
	return final
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 4}
	fmt.Println(RemoveDuplicates(slice))
}
