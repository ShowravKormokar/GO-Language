package main

import "fmt"

type Address struct {
	city string
	zip  uint
}

type Employee struct {
	name    string
	salary  float64
	Address // Embedded struct
}

func main() {
	emp := Employee{
		name:   "Ram",
		salary: 50000.00,
		Address: Address{
			city: "Dhaka",
			zip:  1200,
		},
	}
	fmt.Println(emp.name, emp.salary, emp.city, emp.zip) // Ram 50000 Dhaka 1200
}
