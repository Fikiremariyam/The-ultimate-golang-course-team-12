/*
 Write a function GCD (Greatest Common Divisor) using recursion
in Go. Explain the base case and the recursive case in your solution.*/

package main

import "fmt"

func recursive(num1 int, num2 int, gcf int, iteration int) int {
	if iteration > num1 || iteration > num2 {
		return gcf

	}
	if num1%iteration == 0 && num2%iteration == 0 {
		gcf = iteration
	}
	return recursive(num1, num2, gcf, iteration+1)

}

func main() {

	varibe := recursive(6, 9, 1, 1)
	fmt.Println(varibe)

}
5 =5
4 3 =3 2
3 2 1 2 = 2 1 1 1 
2 1   1 0    1  1 0 =2 1 1 1 =4


0-1
1-1
2 -1
3-2
4-3
5 -5