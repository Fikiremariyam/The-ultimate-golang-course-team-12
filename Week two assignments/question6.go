/* question 6
Implement a function MergeAndSort that takes two sorted slices of
integers as input and returns a single sorted slice with all elements
from both slices. The function should return an error if any of the
slices contain duplicate eleme
*/

package main

import (
	"errors"
	"fmt"
)

func MergeAndSort(slice1 []int, slice2 []int) ([]int, error) {
	helper := func(slice []int) bool {
		for i := 1; i < len(slice1); i++ {
			if slice[i-1] == slice[i] {
				return true
			}
		}
		return false
	}

	if helper(slice1) || helper(slice2) {

		return nil, errors.New("input slice have duplicates")
	}

	var final []int
	pointer1 := 0
	pointer2 := 0

	for pointer1 < len(slice1) && pointer2 < len(slice2) {

		if slice1[pointer1] == slice2[pointer2] {
			final = append(final, slice1[pointer1])
			pointer1++
			pointer2++

		} else if slice1[pointer1] <= slice2[pointer2] {
			final = append(final, slice1[pointer1])
			pointer1++
		} else {
			final = append(final, slice2[pointer2])
			pointer2++
		}
	}

	if pointer1 < len(slice1) {
		final = append(final, slice1[pointer1:]...)
	}
	if pointer2 < len(slice2) {
		final = append(final, slice2[pointer2:]...)
	}

	return final, nil
}

func main() {

	sorted1 := []int{1, 4, 4, 7, 8, 11}
	sorted2 := []int{3, 4, 5, 6, 7}
	print, err := MergeAndSort(sorted1, sorted2)
	if err != nil {

		fmt.Println(err)

	} else {
		fmt.Println(print)
	}
}
