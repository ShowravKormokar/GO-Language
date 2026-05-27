package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	describe(42)
	describe("Hello")
	describe(3.14)
}
