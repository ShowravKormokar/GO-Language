package main

import "fmt"

// Static Array

func main() {
	// Fixed size declaration
	var arr1 [3]int = [3]int{10, 20, 30}

	fmt.Println("Array-1:")
	for i, val := range arr1 {
		fmt.Println(i, "->", val)
	}

	// Short declaration
	arr2 := [3]int{1, 2, 3}

	fmt.Println("Array-2:")
	for _, val := range arr2 {
		fmt.Println(val)
	}

	// Compiler inferred size
	arr3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println("Array-3:")
	for _, val := range arr3 {
		fmt.Println(val)
	}

	fmt.Println("Matrix(2D array):")
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
	for i := range 2 {
		for j := range 3 {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}
