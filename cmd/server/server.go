package server

import (
	"net/http"
	"notes-app-api/cmd/routes"
)

type Server struct {
	*http.Server
}

func NewServer() *Server {
	router := routes.NewRoutes()

	return &Server{
		&http.Server{Addr: ":8080", Handler: router},
	}
}
