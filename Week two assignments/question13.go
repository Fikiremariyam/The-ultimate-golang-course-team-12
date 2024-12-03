/*
Write a function Swap that takes two integers as pointers and
swaps their values. Explain the concept of pointers in Go and why
they are used in this context.
*/
package main

import "fmt"

func swap(input1 *int, input2 *int) {
	temp := *input1
	*input1 = *input2
	*input2 = temp
}

func main() {
	input1 := 15
	input2 := 25

	fmt.Println("before igniting swap functino \n", input1, input2)

	swap(&input1, &input2)

	fmt.Println("after clling th swap function ", input1, input2)

}
