package main

import (
	v1 "cantor/pkg/api"
	"cantor/pkg/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	db.OpenDB()

	r := mux.NewRouter()
	v1router := r.PathPrefix("/v1").Subrouter()

	v1.AddRoutes(v1router)

	graphqlRouter := r.PathPrefix("/graphQL").Subrouter()

	graphqlRouter.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling graphQL request")
	}))

	log.Fatal(http.ListenAndServe(":8080", r))

}
