package services

import (
	"errors"
	"github.com/biangacila/biatechauth1/domain/repositories"
	"github.com/biangacila/biatechauth1/infrastructure/adapters/communitions"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/biangacila/biatechauth1/store"
)

type ForgetPasswordServiceImpl struct {
	EmailService
	serviceUser *UserServiceImpl
}

func NewForgetPasswordServiceImpl(repoUser repositories.UserRepository) *ForgetPasswordServiceImpl {
	var userService = NewUserServiceImpl(repoUser)
	var emailService = communitions.NewCommunicationEmailService() // make sure this implements EmailService
	return &ForgetPasswordServiceImpl{
		EmailService: emailService, // properly initialize EmailService
		serviceUser:  userService,
	}
}

func (f ForgetPasswordServiceImpl) SendOtp(email, systemName string) error {
	// let check if the user email exists
	user, err := f.serviceUser.UserExists(email)
	if err != nil {
		return err
	}

	// Change password allow for local provider only
	if user.Provider != "local" {
		var msg = "password change allow only for local provider and not Google or other third party"
		return errors.New(msg)
	}

	otp, err := f.GenerateOpt()
	if err != nil {
		return err
	}

	err = store.GetStoreOtp().Set(email, otp)
	if err != nil {
		return err
	}

	err = f.EmailService.SendOpt(email, user.String(), otp, systemName)
	if err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		return err
	}
	return nil
}
func (f ForgetPasswordServiceImpl) VerifyOtp(email, opt string) error {
	o, err := store.GetStoreOtp().Get(opt)
	if err != nil {
		return err
	}
	if o.Email != email {
		return errors.New("invalid email associate with Otp")
	}
	return nil
}
func (f ForgetPasswordServiceImpl) ResetPassword(email, opt, password string) error {
	if err := f.VerifyOtp(email, opt); err != nil {
		return err
	}
	if err := f.serviceUser.ResetPassword(email, password); err != nil {
		return err
	}
	return store.GetStoreOtp().Remove(opt)

}
func (f ForgetPasswordServiceImpl) GenerateOpt() (string, error) {
	otp := store.GetStoreOtp().Generate()
	if !store.GetStoreOtp().InStore(otp) {
		return otp, nil
	}
	return "", errors.New("failed to generate unique OTP")
}
