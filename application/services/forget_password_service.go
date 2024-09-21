package services

type ForgetPasswordService interface {
	SendOtp(email, systemName string) error
	VerifyOtp(email, opt string) error
	ResetPassword(email, opt, password string) error
	GenerateOpt() (string, error)
}
