package handlers

import (
	"github.com/biangacila/biatechauth1/interfaces/https/controllers"
	"github.com/gorilla/mux"
)

type Endpoint struct {
	router                    *mux.Router
	userController            controllers.UserController
	loginController           controllers.LoginController
	genericController         controllers.GenericController[any]
	loginWithGoogleController controllers.AuthGoogleController
	forgetPassword            controllers.ForgetPasswordController
}

func NewEndpoint(router *mux.Router, serv *ControllerHandlers) *Endpoint {
	return &Endpoint{
		router:                    router,
		userController:            serv.userController,
		loginController:           serv.loginController,
		genericController:         serv.genericController,
		loginWithGoogleController: serv.loginGoogleController,
		forgetPassword:            serv.forgetPasswordController,
	}
}

func SetupServer(controllerHandlers *ControllerHandlers) *mux.Router {
	router := mux.NewRouter()

	// Create endpoint
	endpoint := NewEndpoint(router, controllerHandlers)
	endpoint.router = router
	endpoint.RegisterRoutes()

	return router
}
