package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRoutes() *mux.Router {
	r := mux.NewRouter()

	NotesRoutes(r)
	UserRoutes(r)

	http.Handle("/", r)

	return r
}
