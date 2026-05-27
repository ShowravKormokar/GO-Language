package main

import (
	"errors"
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
		// return 0, CustomErr{msg: "Not divide by zero", statusCode: 0}
		return 0, errors.New("Not divisible by zero") // By Error Interface
	}

	if a%b != 0 {
		return 0, CustomErr{msg: "Not divisible", statusCode: 1} // By custom error
	}

	return a / b, nil
}

func main() {
	res, err := divide(5, 0)
	if err != nil {
		// Unwraping error [Seperate custom and non-custom errors]
		var cusErr CustomErr
		if errors.As(err, &cusErr) {
			fmt.Print("Custom error occured -> ")
			fmt.Println("msg:", cusErr.msg, ", error code:", cusErr.statusCode)
		} else {
			fmt.Printf("Non-custom error: %s\n", err.Error())
		}
	} else {
		fmt.Println("Result:", res)
	}
}
