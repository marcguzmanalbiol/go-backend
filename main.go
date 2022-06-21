package main

import (
	"cantor/pkg/db"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	var (
		id   int
		name string
	)

	connStr := "postgresql://admin:admin@localhost:5432/cantor?sslmode=disable"

	db.QueryDB()
	db.OpenDB()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM knowledge_areas")
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
