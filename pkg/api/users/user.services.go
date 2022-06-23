package users

import (
	"cantor/pkg/db"
	"database/sql"
	"fmt"
	"log"
)

func getAllTodos() []todo {

	var todos []todo

	var (
		id      int
		userId  int
		name    string
		checked bool
	)

	fmt.Println("Writing all the TODOS:")
	database := db.GetPool()
	rows, err := database.Query("SELECT * FROM todos")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("An error has occurred. ")
		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&id, &userId, &name, &checked)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo{id, userId, name, checked})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return todos
}
