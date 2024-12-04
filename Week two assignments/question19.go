/* Define a function ExecuteWithRecovery that takes a function
as an argument, executes it, and recovers from any panics.
Demonstrate how to use this function with an example that
intentionally causes a panic
*/

package main

import "fmt"

func ExecuteWithRecovery(f func()) {

	defer func() { // the exuted just before return statement so if f() caused an err we handle it here
		if r := recover(); r != nil {
			fmt.Println("recveed from panic")
		}
	}()

	f()

}

func main() {

	panicfunction := func() { // we cause the panic here
		fmt.Println("started the panic funtion")
		panic("ops problem occured!")
	}

	ExecuteWithRecovery(panicfunction)

}
