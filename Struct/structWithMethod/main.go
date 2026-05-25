package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// Reciver Function
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 10.50, height: 8.25}
	fmt.Println("Area:", rect.Area())
}
