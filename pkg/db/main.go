package db

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func GetPool() *sql.DB {
	return DB
}

func OpenDB() {
	fmt.Println("Function to open database!")

	connStr := "postgresql://admin:admin@localhost:5432/cantor?sslmode=disable"

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	DB = database
}
