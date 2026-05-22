package main

import "fmt"

var (
	a = 10
	b = 20
)

func init() {
	fmt.Println("This is the first func that autometically run before main() func when program executed. And cannot be call manually.")
	fmt.Println(a)
	a = 30
}

func main() {
	fmt.Println(a)
	fmt.Println(b)
}
