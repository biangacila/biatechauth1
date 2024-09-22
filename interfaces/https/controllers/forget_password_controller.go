package controllers

import "net/http"

type ForgetPasswordController interface {
	SendOpt(w http.ResponseWriter, r *http.Request)
	VerifyOpt(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}
