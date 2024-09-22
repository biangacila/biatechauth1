package services

type EmailService interface {
	SendOpt(emailAddress, name, otp, systemName string) error
}
