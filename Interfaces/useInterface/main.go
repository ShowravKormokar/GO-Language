package main

import "fmt"

type Walker interface {
	walk()
}

type dog struct {
	name  string
	color string
}

type cat struct {
	name, color, food string
}

// Reciver Function
func (d dog) walk() {
	fmt.Println("Dog name is:", d.name)
}

func (c cat) walk() {
	fmt.Println("Cat name is:", c.name)
}

func caller(w Walker) {
	w.walk()
}

func main() {
	ct := cat{name: "Mew", color: "white", food: "fish"}
	dg := dog{name: "Buck", color: "black"}

	caller(ct)
	caller(dg)
}
