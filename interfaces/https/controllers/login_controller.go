package controllers

import "net/http"

type LoginController interface {
	NewLogin(w http.ResponseWriter, r *http.Request)
	HasLogin(w http.ResponseWriter, r *http.Request)
	IsValidToken(w http.ResponseWriter, r *http.Request)
	IsValidTokenGet(w http.ResponseWriter, r *http.Request)
}
