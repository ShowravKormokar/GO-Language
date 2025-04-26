package main

import "fmt"

const MAX float32 = 4.00

func main() {
	//Way 01
	const PI = 3.1416
	fmt.Println("This is untyped constant value of PI= ", PI)

	//Way 02
	const MONTH int = 12
	fmt.Println("This is typed constant value of MONTH= ", MONTH)

	//Way 03
	fmt.Println("This constant out side of block, MAX= ", MAX)

	//Way 04
	const (
		MIN          = 0
		LEARN string = "Constant"
	)
	fmt.Println("Multiple Constant= ", MIN, LEARN)
}
