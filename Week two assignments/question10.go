/*
Define a custom error type SliceError that includes the
operation performed (e.g., "append", "remove"), the index involved (if
applicable), and a descriptive message. Then, implement a function
SafeRemove that takes a slice of integers and an index, removes the
element at that index, and returns the new slice or an appropriate
SliceError if the index is out of bounds

*/

package main

import (
	"errors"
	"fmt"
)

func SafeRemove(input []int, k int) ([]int, error) {
	if k > len(input) {
		return nil, errors.New("index out of range")
	}
	input = append(input[:k], input[k+1:]...)
	return input, nil
}
func main() {

	input1 := []int{1, 2, 3, 4, 5, 6}

	value, err := SafeRemove(input1, 4)

	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err.Error())
	}
}

//
