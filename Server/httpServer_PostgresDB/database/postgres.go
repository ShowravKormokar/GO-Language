package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() {
	connStr := "host=localhost port=5432 user=postgres password=89@Sk#90 dbname=phoneBook sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	fmt.Println("PostgreSQL connected successfully")
}
