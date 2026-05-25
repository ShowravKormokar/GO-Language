package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

func main() {
	// Pass factor
	double := multiplier(2)
	triple := multiplier(3)

	// Pass value
	fmt.Println(double(5))  // 10 [2 * 5 = 10]
	fmt.Println(double(10)) // 20 [2 * 10 = 20]

	fmt.Println(triple(5))  // 15 [3 * 5 = 15]
	fmt.Println(triple(10)) // 30 [3 * 10 = 30]
}
