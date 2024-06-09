package routes

import (
	"net/http"
	"notes-app-api/cmd/controllers"
)

func NotesRoutes() {
	http.HandleFunc("/notes", controllers.ControllersNotes)
}
