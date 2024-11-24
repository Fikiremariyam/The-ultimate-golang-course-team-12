/*
	 qestion 2

	Implement a function WordFrequency that takes a slice of strings

representing words and returns a map with each unique word as a
key and its count as the value. Ensure the function is case-insensitive
and ignores punctuation.
*/
package main

import (
	"fmt"
	"strings"
)

func WordFrequency(words []string) map[string]int {

	count := make(map[string]int)

	for _, element := range words {
		count[strings.ToLower(element)]++
	}

	return count

}

func main() {

	string := "my name name Name MY "
	slice := strings.Split(string, " ")

	fmt.Println(WordFrequency(slice))
}
