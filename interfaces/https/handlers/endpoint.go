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

	router.HandleFunc(prefix+"/logins-google/login", r.loginWithGoogleController.Login).Methods("GET")
	router.HandleFunc(prefix+"/logins-google/callback", r.loginWithGoogleController.Callback).Methods("GET")

	router.HandleFunc(prefix+"/forget-password/send", r.forgetPassword.SendOpt).Methods("POST")
	router.HandleFunc(prefix+"/forget-password/send/{email}/{system_name}", r.forgetPassword.SendOpt).Methods("GET")
	router.HandleFunc(prefix+"/forget-password/verify", r.forgetPassword.VerifyOpt).Methods("POST")
	router.HandleFunc(prefix+"/forget-password/verify/{email}/{opt}", r.forgetPassword.VerifyOpt).Methods("GET")
	router.HandleFunc(prefix+"/forget-password/reset", r.forgetPassword.ResetPassword).Methods("POST")
	router.HandleFunc(prefix+"/forget-password/reset/{email}/{opt}/{password}", r.forgetPassword.ResetPassword).Methods("GET")

}
