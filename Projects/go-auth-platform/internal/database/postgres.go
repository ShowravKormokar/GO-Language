package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() {
	conn_str := "host=localhost port=5432 user=postgres password=89@Sk#90 dbname=GOAuth sslmode=disable"

	DB, err := sql.Open("postgres", conn_str)
	if err != nil {
		log.Fatal("Failed to connect:", err)

	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	fmt.Println("PostgreSQL connected successfully")
}
