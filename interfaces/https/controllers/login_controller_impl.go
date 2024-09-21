package controllers

import (
	"encoding/json"
	"errors"
	"github.com/biangacila/biatechauth1/application/dtos"
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type LoginControllerImpl struct {
	service services.LoginService
}

func NewLoginController(
	service services.LoginService,
) *LoginControllerImpl {
	return &LoginControllerImpl{
		service: service,
	}
}
func (l LoginControllerImpl) NewLogin(w http.ResponseWriter, r *http.Request) {

	payload := dtos.LoginDto{}

	// This will reject any fields that are not part of the DTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}
	if err = dtos.Validate(payload, dtos.LoginDto{}); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	user, token, err := l.service.NewLogin(payload.Username, payload.Password)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (l LoginControllerImpl) HasLogin(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	username := vars["username"]
	if username == "" {
		http.Error(w, utils.HttpResponseError(errors.New("invalid username")), http.StatusBadRequest)
		return
	}
	expiredAt, ok, err := l.service.HasLogin(username)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"expired_at": expiredAt,
		"status":     ok,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(response)
}

func (l LoginControllerImpl) IsValidTokenGet(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	token := vars["token"]
	if token == "" {
		http.Error(w, utils.HttpResponseError(errors.New("empty token")), http.StatusBadRequest)
		return
	}
	expiredAt, ok, err := l.service.IsValueToken(token)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"expired_at": expiredAt,
		"status":     ok,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(response)
}
func (l LoginControllerImpl) IsValidTokenPost(w http.ResponseWriter, r *http.Request) {
	payload := dtos.LoginCheckTokenDto{}

	// This will reject any fields that are not part of the DTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	token := payload.Token
	if token == "" {
		http.Error(w, utils.HttpResponseError(errors.New("empty token")), http.StatusBadRequest)
		return
	}
	expiredAt, ok, err := l.service.IsValueToken(token)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"expired_at": expiredAt,
		"status":     ok,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(response)
}
