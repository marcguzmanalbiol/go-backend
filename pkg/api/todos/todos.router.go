package todos

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func AddTodosRouters(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World!")
	})

	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		todos := getAllTodos()
		fmt.Println(todos)
	})
}
