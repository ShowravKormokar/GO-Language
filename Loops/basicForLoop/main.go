package main

import "fmt"

// In GO only one loop presents "for"
func main() {
	// initialization; condition; iteration
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\t", i+1)
	}
}
