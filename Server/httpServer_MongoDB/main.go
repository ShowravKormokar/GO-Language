package main

import (
	"fmt"
	"net/http"
	"server_MongoDB/database"
)

func main() {
	database.ConnctMongo()

	fmt.Println("Server running on port:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Couldn't connect:", err)
	}
}
