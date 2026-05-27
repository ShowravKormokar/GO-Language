package main

import "fmt"

func funcDeferTest() {
	fmt.Println("Defer function.")
	if r := recover(); r != nil {
		fmt.Println("Recover from", r)
	}
}

func main() {
	defer funcDeferTest() // Defer func

	fmt.Println("Start")
	defer fmt.Println("Deferred 1")
	// panic("something wrong")
	defer fmt.Println("Deferred 2")
	fmt.Println("End")
}

/*
1. Multiple Defer function execute in LIFO order
2. Panic [Identify some critical bug/unrecoverable error like invalid memory access]
*/
