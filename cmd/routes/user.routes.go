package routes

import (
	"net/http"
	"notes-app-api/cmd/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user", controllers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user", controllers.GetUser).Methods(http.MethodGet)
}
