package controllers

import "net/http"

type AuthGoogleController interface {
	Login(http.ResponseWriter, *http.Request)
	Callback(http.ResponseWriter, *http.Request)
	ValidateToken(w http.ResponseWriter, r *http.Request)
	LoginWithGoogle(w http.ResponseWriter, r *http.Request)
}
