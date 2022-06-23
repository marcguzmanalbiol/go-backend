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

	/* ROUTER */

	r := mux.NewRouter()
	v1router := r.PathPrefix("/v1").Subrouter()

	v1.AddRoutes(v1router)

	graphQLrouter := r.PathPrefix("/graphQL").Subrouter()

	graphQLrouter.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling graphQL request")
	}))

	log.Fatal(http.ListenAndServe(":8080", r))

}
