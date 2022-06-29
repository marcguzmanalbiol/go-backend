package v1

import (
	"cantor/pkg/api/posts"
	"cantor/pkg/api/todos"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	postRouter := router.PathPrefix("/posts").Subrouter()
	posts.AddPostRouters(postRouter)

	todosRouter := router.PathPrefix("/todos").Subrouter()
	todos.AddTodosRouters(todosRouter)
}
