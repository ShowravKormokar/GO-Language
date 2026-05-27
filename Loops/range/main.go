package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	for i, ele := range arr {
		fmt.Printf("%d -> %d\n", i, ele)
	}
}
