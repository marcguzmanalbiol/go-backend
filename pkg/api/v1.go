package v1

import (
	"cantor/pkg/api/posts"
	"cantor/pkg/api/users"

	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) {
	post_router := r.PathPrefix("/posts/").Subrouter()
	posts.AddPostRouters(post_router)

	users_router := r.PathPrefix("/users").Subrouter()
	users.AddUsersRouters(users_router)
}
