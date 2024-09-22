package controllers

import (
	"encoding/json"
	"github.com/biangacila/biatechauth1/application/dtos"
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type ForgetPasswordControllerImpl struct {
	service services.ForgetPasswordService
}

func NewForgetPasswordController(
	service services.ForgetPasswordService,
) ForgetPasswordControllerImpl {
	return ForgetPasswordControllerImpl{
		service: service,
	}
}

/*func NewAuthGoogleController(
	service services.LoginService,
) *AuthGoogleControllerImpl {
	return &AuthGoogleControllerImpl{
		service: service,
	}
}*/

func (f ForgetPasswordControllerImpl) SendOpt(w http.ResponseWriter, r *http.Request) {
	var payload dtos.ForgetPasswordSendDto
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		payload = dtos.ToEntity(vars, dtos.ForgetPasswordSendDto{})
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
			return
		}
	}
	// let validate our payload
	if err := dtos.ValidateAnyWithAnyDto(payload, dtos.ForgetPasswordSendDto{}); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	err := f.service.SendOtp(payload.Email, payload.SystemName)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"status": "sent",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}
func (f ForgetPasswordControllerImpl) VerifyOpt(w http.ResponseWriter, r *http.Request) {
	var payload dtos.ForgetPasswordVerifyDto
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		payload = dtos.ToEntity(vars, dtos.ForgetPasswordVerifyDto{})
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
			return
		}
	}
	// let validate our payload
	if err := dtos.ValidateAnyWithAnyDto(payload, dtos.ForgetPasswordVerifyDto{}); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	err := f.service.VerifyOtp(payload.Email, payload.Otp)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"status": "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
func (f ForgetPasswordControllerImpl) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var payload dtos.ForgetPasswordResetDto
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		payload = dtos.ToEntity(vars, dtos.ForgetPasswordResetDto{})
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
			return
		}
	}
	// let validate our payload
	if err := dtos.ValidateAnyWithAnyDto(payload, dtos.ForgetPasswordResetDto{}); err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusBadRequest)
		return
	}

	err := f.service.ResetPassword(payload.Email, payload.Otp, payload.Password)
	if err != nil {
		http.Error(w, utils.HttpResponseError(err), http.StatusInternalServerError)
	}
	response := map[string]interface{}{
		"status": "reset",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
