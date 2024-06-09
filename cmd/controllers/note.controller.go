package controllers

import (
	"net/http"
	"notes-app-api/cmd/adapters"
)

func ControllersNotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		adapters.GetNotes(w, r)
	case "POST":
		adapters.CreateNote(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
