package e_test

import (
	"fmt"
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/biangacila/biatechauth1/infrastructure/adapters/cassandradb"
	"github.com/biangacila/biatechauth1/interfaces/https/controllers"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController(t *testing.T) {
	cassandradb.InitSession(constants.DbHost, constants.DbName)
	repo := cassandradb.NewCassandraUserRepository()
	service := services.NewUserServiceImpl(repo)
	controller := controllers.NewUserController(service)

	t.Run("create user", func(t *testing.T) {
		params, err := utils.ObjectToBufferReader(map[string]interface{}{
			"given_name":   "Merveilleux",
			"family_name":  "Biangacila",
			"email":        "Biangacila@gmail.com",
			"phone_number": "0729139504",
			"password":     "Nathan010309*",
		})
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		req, _ := http.NewRequest("POST", "/api/users", params)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		resp := httptest.NewRecorder()
		controller.Create(resp, req)

		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	})

}
