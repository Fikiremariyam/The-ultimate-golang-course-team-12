/*
Write a function Retry that takes another function and a number
of retries as arguments. If the provided function returns an error, the
Retry function should attempt to call it again up to the specified
number of retries before returning the error
*/

package main

import (
	"errors"
	"fmt"
)

func Retry(func1 func() error, tries int) {

	if tries >= 0 {
		error := func1()
		if error != nil {
			Retry(func1, tries-1)
		}
	}

}

func tempo() error {
	fmt.Println("called for")
	final := errors.New("this function eonsistantly cause error ")
	return final
}

func main() {

	Retry(tempo, 5)

}
