package main

import "fmt"

var a = 10

func main() {
	age := 30

	if age > 18 {
		a := 20
		fmt.Println("Under if block: ", a)// 20, because a is under local block that shadow of global a
	}

	fmt.Println("Outer if block: ", a)//10, beacuse a comes from global(original)
}
