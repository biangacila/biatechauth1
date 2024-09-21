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

func TestLoginController(t *testing.T) {
	cassandradb.InitSession(constants.DbHost, constants.DbName)

	repoUser := cassandradb.NewCassandraUserRepository()
	repo := cassandradb.NewCassandraLoginRepository()
	service := services.NewLoginServiceImpl(repo, repoUser)
	controller := controllers.NewLoginController(service)

	t.Run("create login", func(t *testing.T) {
		params, err := utils.ObjectToBufferReader(map[string]interface{}{
			"username": "Biangacila@gmail.com",
			"password": "Nathan010309*",
		})
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		req, _ := http.NewRequest("POST", "/api/logins", params)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		resp := httptest.NewRecorder()
		controller.NewLogin(resp, req)

		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	})

}
