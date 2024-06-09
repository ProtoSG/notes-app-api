package utils

import (
	"net/http"
	"notes-app-api/cmd/routes"
)

type Server struct {
	*http.Server
}

func NewServer() *Server {
	routes.NotesRoutes()

	return &Server{
		&http.Server{Addr: ":8080"},
	}
}
