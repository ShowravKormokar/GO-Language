package main

import "fmt"

func main() {
	funcTypes := []string{"Standard or Named Func", "Anonymous Function", "Closure Function", "Recursive Function", "Variadic Function", "Higher-Order Function", "Generator Function", "Callback Function", "Lambda Function", "Method Function", "Interface Method", "Goroutine Function", "Defer Function", "Panic and Recover Function", "Error Handling Function", "Testing Function", "Benchmarking Function", "Initialization Function", "Main Function", "Package-Level Function", "Function Literals", "First-Class Functions", "IIFE (Immediately Invoked Function Expression)", "Reciver Function"}

	for i, v := range funcTypes {
		fmt.Println(i, ".", v)
	}
}
