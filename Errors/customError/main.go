package main

import (
	"fmt"
)

type CustomErr struct {
	msg        string
	statusCode int
}

func (err CustomErr) Error() string {
	return fmt.Sprintf("Caught Error: %d : %s", err.statusCode, err.msg)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, CustomErr{msg: "Not divide by zero", statusCode: 0}
	}

	if a%b != 0 {
		return 0, CustomErr{msg: "Not divisible", statusCode: 1}
	}

	return a / b, nil
}

func main() {
	res, err := divide(5, 2)
	if err != nil {
		fmt.Println(err.Error())
		// return
	} else {
		fmt.Println("Result:", res)
	}
}
