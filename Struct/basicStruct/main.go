package main

import (
	"fmt"
)

// Structure name as Users
type Users struct {
	_id  uint8 // 0 - 255
	name string
	age  uint // Only positive
}

func main() {
	// Instance of struct, name as "user"
	user1 := Users{_id: 255, name: "Showrav", age: 24}

	fmt.Println(user1._id, user1.name, user1.age) // 255 Showrav 24
}
