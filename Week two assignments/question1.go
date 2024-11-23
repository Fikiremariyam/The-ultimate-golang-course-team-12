/* qustion1
1. Write a function ResizeSlice that takes a slice of integers and a
target length as arguments. If the slice is smaller than the target
length, append zeros until it reaches the target length. If it's larger,
truncate the slice. Ensure that the function handles the edge cases
correctly (e.g., when the target length is zero or negative).

*/

package main

import (
	"fmt"
	"slices"
)

func ResizeSlice(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return numbers
	}
	if len(numbers) < target {
		for i := 0; i < target-len(numbers); i++ {
			numbers = slices.Insert(numbers, len(numbers), 0)
		}
	}
	return numbers[:target]
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 8, 8, 9}
	final := ResizeSlice(slice, 4)

	fmt.Println(final)
}
