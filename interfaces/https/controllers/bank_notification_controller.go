package controllers

import "net/http"

type BankNotificationController interface {
	ReceiveNotificationGet(w http.ResponseWriter, r *http.Request)
	ReceiveNotificationPost(w http.ResponseWriter, r *http.Request)
}
