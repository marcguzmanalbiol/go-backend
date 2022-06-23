package db

import (
	"database/sql"
	"fmt"
	"log"
)

func GetPool() *sql.DB {
	return DB
}

var DB *sql.DB

func OpenDB() *sql.DB {
	fmt.Println("Function to open database!")

	connStr := "postgresql://admin:admin@localhost:5432/cantor?sslmode=disable"

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	DB = database
	return database
}
