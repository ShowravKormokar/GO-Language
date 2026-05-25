package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	inc := counter()
	inc2 := counter()

	// Each closure has its own isolated state.

	fmt.Println(inc()) // 1
	fmt.Println(inc()) // 2
	fmt.Println(inc()) // 3

	fmt.Println(inc2()) // 1
}
