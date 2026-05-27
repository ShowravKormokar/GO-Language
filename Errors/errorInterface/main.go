package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Not divisible by zero")
	}
	return a / b, nil
}

func main() {
	res, err := divide(2, 0)
	if err != nil {
		fmt.Println("Caught error:", err.Error())
		// return
	} else {
		fmt.Println("Result:", res)
	}
}
