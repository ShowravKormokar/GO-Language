package userinput

import "fmt"

func UserInputInt() int {
	var num int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return -1 // Return -1 to indicate an error
	}

	return num
}

func UserInputString() string {
	var str string

	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return "" // Return an empty string to indicate an error
	}

	return str
}
