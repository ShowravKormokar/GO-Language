package main

import "fmt"

func getInputs() (int, int) {
	var p, q int
	fmt.Print("Enter 1st num:")
	fmt.Scanln(&p)
	fmt.Print("Enter 2nd num:")
	fmt.Scanln(&q)

	return p, q
}

func addMul(x int, y int) (int, int) {
	var sum int = x + y
	var mul int = x * y

	return sum, mul
}

func main() {
	var a, b int

	a, b = getInputs()
	// Call the function passing arguments
	add, mul := addMul(a, b)

	fmt.Println("Sum:", add, "Mul:", mul)
}
