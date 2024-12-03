/*
Implement a higher-order function ApplyToEach that takes a
slice of integers and a function as arguments. The function should
apply the provided function to each element of the slice and return a
new slice with the results.
*/
package main

import "fmt"

func ApplyToEach(input1 []int, operation func([]int) []int) []int {
	return operation(input1)

}

func add(input []int) []int {
	for index, _ := range input {
		input[index] += 2

	}
	return input

}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(ApplyToEach(slice, add))

}
