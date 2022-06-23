package main

import (
	v1 "cantor/pkg/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	/* 	var (
	   		id      int
	   		user_id int
	   		name    string
	   		checked bool
	   	)

	   	connStr := "postgresql://admin:admin@localhost:5432/cantor?sslmode=disable"

	   	db.QueryDB()
	   	db.OpenDB()

	   	db, err := sql.Open("postgres", connStr)
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   	rows, err := db.Query("SELECT * FROM todos")
	   	if err != nil {
	   		log.Fatalln(err)
	   	}

	   	defer rows.Close()

	   	for rows.Next() {
	   		err := rows.Scan(&id, &user_id, &name, &checked)
	   		if err != nil {
	   			log.Fatal(err)
	   		}
	   		fmt.Println(id, user_id, name, checked)
	   	}
	   	err = rows.Err()
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	*/

	r := mux.NewRouter()
	v1router := r.PathPrefix("/v1").Subrouter()

	v1.AddRoutes(v1router)

	graphQLrouter := r.PathPrefix("/graphQL").Subrouter()

	graphQLrouter.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling graphQL request")
	}))

	log.Fatal(http.ListenAndServe(":8080", r))

}
