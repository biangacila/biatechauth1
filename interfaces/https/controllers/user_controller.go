package controllers

import (
	"github.com/biangacila/biatechauth1/domain/entities"
	"net/http"
)

type UserController interface {
	GenericController[entities.User]
	Create(w http.ResponseWriter, r *http.Request)
	Lock(w http.ResponseWriter, r *http.Request)
	Unlock(w http.ResponseWriter, r *http.Request)
	Exist(w http.ResponseWriter, r *http.Request)
}
