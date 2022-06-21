package posts

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AddPostRouters(r *mux.Router) {
	r.HandleFunc("/", HelloWorld)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}
