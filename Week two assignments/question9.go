/*
Write a function ReplaceMapKeys that takes a map of string keys to
integer values and replaces every key by reversing its string. The
function should perform the replacement in-place (modifying the
original map without creating a new one).

*/

package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s) // Handle Unicode characters properly
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReplaceMapKeys(input1 map[string]int) {
	for key, value := range input1 {
		newkey := reverseString(key)
		input1[newkey] = value

		delete(input1, key)
	}

}
func main() {
	input := map[string]int{
		"one": 1,
		"two": 2,
	}
	ReplaceMapKeys(input)
	fmt.Println(input)
}

//note :- i have done it but its dosent iterate the all over the  map
