package main

import "fmt"

// Variadic Function (Not a specific no. of parameters recive)
// Recive as slice
func printArr(nums ...int) {
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}

func main() {
	printArr(1, 2, 3, 4, 5)
}
