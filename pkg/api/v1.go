package v1

import (
	"cantor/pkg/api/posts"
	"cantor/pkg/api/users"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	postRouter := router.PathPrefix("/posts/").Subrouter()
	posts.AddPostRouters(postRouter)

	usersRouter := router.PathPrefix("/users").Subrouter()
	users.AddUsersRouters(usersRouter)
}
