package main

import "fmt"

func main() {
	// Dynamic array (slice)
	arr := []int{10, 20, 30}
	fmt.Println(arr) // [10 20 30]

	// Append new elements
	arr = append(arr, 40, 50)
	fmt.Println(arr) // [10 20 30 40 50]

	// Length andCapacity
	fmt.Println("Length:", len(arr))   // 5
	fmt.Println("Capacity:", cap(arr)) // capacity auto increase
}
