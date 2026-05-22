package main

import "fmt"

func main() {
	// Anonymous Function
	// If this func. call immediatly, it called as Immediatly Invocked Function Expression (IIFE)
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(10, 10)
}