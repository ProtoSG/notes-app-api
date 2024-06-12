package routes

import (
	"net/http"
	"notes-app-api/cmd/controllers"

	"github.com/gorilla/mux"
)

func NotesRoutes(r *mux.Router) {
	r.HandleFunc("/notes", controllers.GetNotes).Methods(http.MethodGet)
	r.HandleFunc("/notes", controllers.CreateNote).Methods(http.MethodPost)
	r.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods(http.MethodPut)
	r.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods(http.MethodDelete)
}
