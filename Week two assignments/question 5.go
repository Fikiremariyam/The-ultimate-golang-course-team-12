/*
 Write a function PartitionEvenOdd that takes an array of integers
and returns two slices: one containing all even numbers and the other
containing all odd numbers.
*/

package main

import (
	"fmt"
)

func PartitionEvenOdd(numbers []int) ([]int, []int) {
	var even []int
	var odd []int
	for _, value := range numbers {
		if value%2 == 0 {
			even = append(even, value)
		} else {
			odd = append(odd, value)
		}
	}
	return even, odd

}

func main() {
	slices := []int{1, 2, 3, 4, 5}
	fmt.Println(PartitionEvenOdd(slices))
}
