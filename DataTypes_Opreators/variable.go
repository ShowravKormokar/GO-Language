// int- stores integers (whole numbers), such as 123 or -123
// float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
// string - stores text, such as "Hello World". String values are surrounded by double quotes
// bool- stores values with two states: true or false

package main

import "fmt"

var isOK bool = true // Var is working in both in/out side of function, but inferrend is only block scope work

func main() {

	// variable declaration type-01 : (Syntax: var variableName type = value)
	var a int = 5
	var b int = 5

	fmt.Println("1st sum= ", a+b)

	// variable declaration type-02 : (Syntax: variablename := value)
	x := 10 // the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value)
	y := 15 //It is not possible to declare a variable using :=, without assigning a value to it.

	fmt.Println("2nd sum= ", x+y)

	// variable declaration example
	var fName string
	fName = "Showrav"
	lName := "Kormokar"

	fmt.Println("Full Name: ", fName+lName)
	// Example 2
	if isOK {
		fmt.Println(isOK)
	}

	// Example 3
	var result float32 = 3.97
	fmt.Println("Your result is: ", result)
}
