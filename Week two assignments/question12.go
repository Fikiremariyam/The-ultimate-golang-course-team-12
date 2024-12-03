/*
Create a function IsPalindrome that takes a string as input and
checks whether it is a palindrome. Make the function case-insensitive
and ignore spaces or punctuation.

Approach :
first remove all the spaces and punuations
second check if it is a palandrome
*/

package main

import "fmt"

func clean(input1 string) []string {

	var cleaned []string

	for i := 0; i < len(input1); i++ {
		if input1[i] > 65 && input1[i] < 122 {

			cleaned = append(cleaned, string(input1[i]))

		}

	}
	return cleaned

}
func IsPalindrome(input1 []string) bool {

	for i := 0; i < len(input1)/2; i++ {
		if input1[i] != input1[len(input1)-1-i] {
			return false

		}

	}
	return true

}

func main() {
	text := "Mmmm"

	val := IsPalindrome(clean(text))
	if val == true {
		fmt.Println("it is palandrome")

	} else {
		fmt.Println("it is not palandrom")
	}

}
