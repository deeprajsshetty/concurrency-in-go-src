package basicgoroutine

import "fmt"

/*
	In Golang, we can create functions in THREE different ways. Created below.
*/
func BasicGoroutine() {
	// Using Function
	go sayHello()

	// Using Anonymous Function
	go func() {
		fmt.Println("hello anonymous function")
	}()

	// Using assign Function to Variable
	sayHelloUsingVariable := func() {
		fmt.Println("hello function assigning to variable")
	}
	go sayHelloUsingVariable()
}

func sayHello() {
	fmt.Println("hello normal function")
}
