package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
}

// Struct
type Rectangle struct {
	width, height float64
}

// Method implement interface for Rectangle struct to satisfy the Shape interface
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var s Shape
	s = Rectangle{10, 5}
	fmt.Println("Area:", s.Area()) // Output: 50
}
