/*
Write a function WindowSum that takes a slice of integers and a
window size k and returns a new slice where each element is the
sum of a sliding window of size k from the input slice. If k is greater
than the length of the slice or less than 1, return an error.
*/

package main

import (
	"errors"
	"fmt"
)

func WindowSum(Array []int, k int) ([]int, error) {
	if k > len(Array) {
		return nil, errors.New("the length of the arry is less than the size of k ")
	}
	var final []int
	left, right := 0, k
	for right < len(Array) {
		var tempo int
		for _, val := range Array[left:right] {
			tempo += val

		}
		final = append(final, tempo)

		left++
		right++

	}
	return final, nil

}

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	final, error := WindowSum(array, 3)
	if error != nil {
		fmt.Println(error.Error())

	} else {
		fmt.Println(final)
	}

}
