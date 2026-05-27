package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		fmt.Printf("%d\t", i*2+1)

		i++
	}
}
