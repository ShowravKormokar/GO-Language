package main

import "fmt"

// Anonymous fun assign in variable
var printSum = func(total int) {
	fmt.Println("Sum:", total)
}

func main() {
	// Function expression or assign function in variable
	add := func(a int, b int) {
		c := a + b
		printSum(c)
	}

	add(5, 10)
}
