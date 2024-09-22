package e_test

import (
	"fmt"
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/biangacila/biatechauth1/infrastructure/adapters/cassandradb"
	"github.com/biangacila/biatechauth1/interfaces/https/controllers"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/biangacila/biatechauth1/store"
	"github.com/gorilla/mux"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForgetPasswordController(t *testing.T) {
	cassandradb.InitSession(constants.DbHost, constants.DbName)
	store.InitTokens()
	store.InitialOtpStore()

	repoUser := cassandradb.NewCassandraUserRepository()
	service := services.NewForgetPasswordServiceImpl(repoUser)
	controller := controllers.NewForgetPasswordController(service)

	t.Run("create login", func(t *testing.T) {
		email := "Biangacila@gmail.com"
		systemName := "Chela"

		params, err := utils.ObjectToBufferReader(map[string]interface{}{
			"email":       email,
			"system_name": systemName,
		})
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// step 1
		req, _ := http.NewRequest("POST", "/api/forgetpassword/send", params)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		resp := httptest.NewRecorder()
		controller.SendOpt(resp, req)
		b, _ := io.ReadAll(resp.Body)
		fmt.Println("step 1 result> ", string(b))

		// Step 2: Prompt for OTP input and verify
		fmt.Println("Please enter the OTP sent to your email:")
		var inputOTP string
		scanln, err := fmt.Scanln(&inputOTP)
		if err != nil {
			fmt.Println("&>Error scanln:", err, scanln)
			return
		}

		params, err = utils.ObjectToBufferReader(map[string]interface{}{
			"email": email,
			"opt":   inputOTP,
		})
		req, _ = http.NewRequest("POST", "/api/forgetpassword/verify", params)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		resp = httptest.NewRecorder()
		controller.VerifyOpt(resp, req)
		b, _ = io.ReadAll(resp.Body)
		fmt.Println("step 2 result> ", string(b))

		// Step 3: Prompt for new password and reset
		fmt.Println("Please enter your new password:")
		var newPassword string

		params, err = utils.ObjectToBufferReader(map[string]interface{}{
			"email":    email,
			"opt":      inputOTP,
			"password": newPassword,
		})
		req, _ = http.NewRequest("POST", "/api/forgetpassword/verify", params)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		resp = httptest.NewRecorder()
		controller.VerifyOpt(resp, req)
		b, _ = io.ReadAll(resp.Body)
		fmt.Println("step 3 result> ", string(b))

	})
}
