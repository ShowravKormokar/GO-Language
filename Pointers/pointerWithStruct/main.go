package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Birthday() {
	p.age++
}

func main() {
	person := Person{name: "Showrav", age: 25}
	person.Birthday()                    // Not need to send to address
	fmt.Println(person.name, person.age) // Showrav 26
}
