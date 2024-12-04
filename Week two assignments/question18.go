/*
Write a function ReverseString that takes a string as input and
returns its reverse. Use rune slices to handle Unicode characters
properly
*/
package main

import (
	"fmt"
)

func REverseString(input string) string {
	runes := []rune(input)

	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string
	return string(runes)

}

func main() {
	str1 := "hi how are you"
	str2 := REverseString(str1)
	fmt.Println(str2)
}
