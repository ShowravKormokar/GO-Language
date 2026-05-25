package main

import "fmt"

func updateValue(val *int) {
	*val = *val + 5
}

func main() {
	num := 10
	updateValue(&num)            // send address
	fmt.Println("Updated:", num) // 15
}
