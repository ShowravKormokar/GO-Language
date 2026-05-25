package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println(x)
	fmt.Println(p)  // Show Address
	fmt.Println(*p) // Pointer dereference (get value from the address)
	*p++            // Change value of p, also change the x value
	fmt.Println(x)

	var y int // Not assign any value automatically hold 0 value
	q := &y
	fmt.Println(y)
	fmt.Println(q)
	fmt.Println(*q)
}
