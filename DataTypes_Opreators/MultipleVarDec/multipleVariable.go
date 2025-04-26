package main

import "fmt"

func main() {
	// Declare multiple variable way-01
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println("a=", a, ", b=", b, ", c=", c, ", d=", d)

	// Declare multiple variable way-02
	var x, y = 77, "Showrav"
	fmt.Println("ID:", x, ", Name:", y)

	// Declare multiple variable way-03
	var (
		m int
		n int    = 1
		o string = "GO"
	)

	fmt.Println("m=", m)
	fmt.Println("n=", n)
	fmt.Println("o=", o)
}
