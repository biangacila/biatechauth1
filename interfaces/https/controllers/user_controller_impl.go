package controllers

import (
	"encoding/json"
	"github.com/biangacila/biatechauth1/application/dtos"
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/domain/entities"
	"github.com/biangacila/biatechauth1/internal/utils"
	"net/http"
)

type UserControllerImpl struct {
	GenericController[entities.User]
	service services.UserService
}

func NewUserController(service services.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		GenericController: NewGenericController[entities.User](service), // Reuse the generic controller
		service:           service,
	}
}

func (u *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	payload := dtos.UserPayloadDto{}

	// This will reject any fields that are not part of the DTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	if err := dtos.ValidateAnyWithAnyDto(payload, dtos.UserPayloadDto{}); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	user, err := u.service.Create(payload.GivenName, payload.FamilyName, payload.Email, payload.Phone, payload.Password)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusInternalServerError)
		return
	}
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func (u *UserControllerImpl) Lock(w http.ResponseWriter, r *http.Request) {
	payload := dtos.UserPayloadLockDto{}

	// This will reject any fields that are not part of the DTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	err := u.service.Lock(payload.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status": "locked",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func (u *UserControllerImpl) Unlock(w http.ResponseWriter, r *http.Request) {
	payload := dtos.UserPayloadLockDto{}

	// This will reject any fields that are not part of the DTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	err := u.service.UnLock(payload.Username)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status": "locked",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func (u *UserControllerImpl) Exist(w http.ResponseWriter, r *http.Request) {
	payload := dtos.UserPayloadLockDto{}

	// This will reject any fields that are not part of the DTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	user, err := u.service.UserExists(payload.Username)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"user": user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(response)
}
