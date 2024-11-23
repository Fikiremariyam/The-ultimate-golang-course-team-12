/*
	task

Write a function FindEmployeeAge that takes a map of employee
names (string) to ages (int) and an employee name as input. Return
the age of the employee if found, or an error if the employee does not
exist in the map.
*/
package main

import (
	"errors"
	"fmt"
)

func FindEmployeeAge(employee_names map[string]int, name string) (int, error) {
	for key, age := range employee_names {
		if key == name {
			return age, nil // Return the age and no error
		}
	}
	// Return 0 and an error if the name is not found
	return 0, errors.New("employee not found")
}

func main() {

	emplyee_data := map[string]int{
		"abebe":  12,
		"kebede": 13,
		"alemu":  14,
	}

	age, err := FindEmployeeAge(emplyee_data, "kebede")

	if err != nil {
		fmt.Println("Error ", err)

	} else {

		fmt.Println(age)

	}

}
