package main

import (
	"fmt"
	"net/http"
	"server_postgres/database"
	"server_postgres/routes"
)

func main() {
	var err error

	// Database Connect
	database.ConnectPostgres()
	_, err = database.DB.Exec(`CREATE TABLE IF NOT EXISTS contacts(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	phone TEXT NOT NULL,
	description TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}

	// Routes
	r := routes.RegisterRoutes()

	// Connect server
	fmt.Println("Server running on port :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Couldn't connect to server:", err)
	}
}
