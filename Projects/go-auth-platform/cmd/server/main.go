package main

import (
	"fmt"
	"go-auth-platform/routes"
	"net/http"
)

func main() {
	r := routes.RegisterRouter()

	fmt.Println("Server connected successfully on port:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Server couldn't connected!", err)
	}
}
