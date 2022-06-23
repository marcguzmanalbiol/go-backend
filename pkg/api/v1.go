package v1

import (
	"cantor/pkg/api/posts"
	"cantor/pkg/api/users"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	post_router := router.PathPrefix("/posts/").Subrouter()
	posts.AddPostRouters(post_router)

	users_router := router.PathPrefix("/users").Subrouter()
	users.AddUsersRouters(users_router)
}
