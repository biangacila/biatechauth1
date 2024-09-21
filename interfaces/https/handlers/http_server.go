package handlers

import (
	"github.com/gorilla/mux"
)

func SetupServer() *mux.Router {
	router := mux.NewRouter()

	// todo Initialize repositories

	// todo Initialize services

	// todo Initialize controllers

	RegisterRoutes(router)
	return router
}
