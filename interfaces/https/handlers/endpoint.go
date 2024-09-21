package handlers

func (r *Endpoint) RegisterRoutes() {
	prefix := "/backend-biatechauth1/api"
	router := r.router

	router.HandleFunc(prefix+"/users", r.userController.Create).Methods("POST")
	router.HandleFunc(prefix+"/users/exist", r.userController.Exist).Methods("POST")
	router.HandleFunc(prefix+"/users/exist/{username}", r.userController.ExistGet).Methods("GET")

	router.HandleFunc(prefix+"/logins", r.loginController.NewLogin).Methods("POST")
	router.HandleFunc(prefix+"/logins/has-login/{username}", r.loginController.HasLogin).Methods("GET")
	router.HandleFunc(prefix+"/logins/valid-token", r.loginController.IsValidToken).Methods("POST")
	router.HandleFunc(prefix+"/logins/valid-token/{token}", r.loginController.IsValidTokenGet).Methods("GET")

}
