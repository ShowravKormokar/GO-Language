package main

import "fmt"

// A closure is an anonymous function that captures variables from its outer scope.
// This counter func return closure func
func counter() func() int {
	n := 0 // Outer scope variable, that held closure by it's own state

	// This anonymous function is actually closure function
	return func() int {
		n++ // Outer scope variable that modify
		return n
	}
}

func main() {
	inc := counter() // Get closure function to call the counter func

	// Here remember the 'n' variable everytime when call the inc
	fmt.Println(inc()) // 1
	fmt.Println(inc()) // 2
	fmt.Println(inc()) // 3
}
