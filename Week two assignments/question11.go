/*
Write a function SumVariadic that takes a variadic parameter of
integers and returns their sum. Demonstrate how to call this function
with both a slice of integers and individual integer arguments.


Apprach  :
ussagee of  int...
*/

package main

import "fmt"

func SumVariadic(input1 ...int) int {
	var sum int
	for _, i := range input1 {
		sum += i

	}
	return sum

}

func main() {

	indivsualNumbers := SumVariadic(1, 2, 3, 4, 5, 5)
	fmt.Println("the sum of the indicusual numbers in sumVaridiac funciton is  \n", indivsualNumbers)
	sliceOfNumbers := []int{1, 2, 3, 4, 5}

	sliceOfNumberss := SumVariadic(sliceOfNumbers...)
	fmt.Println("the sum of  slice of numbers using the same function is", sliceOfNumberss)

}
