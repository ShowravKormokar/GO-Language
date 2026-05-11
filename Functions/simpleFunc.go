package main

import "fmt"

func add(x int, y int) {
	var sum int = x + y

	fmt.Println("Sum:", sum)
}

func main() {
	var a, b int = 3, 4
	// Call the function passing arguments
	add(a, b)
}
