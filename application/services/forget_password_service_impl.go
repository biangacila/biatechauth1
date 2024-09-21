package services

import (
	"errors"
	"github.com/biangacila/biatechauth1/domain/repositories"
	"github.com/biangacila/biatechauth1/infrastructure/adapters/communitions"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/biangacila/biatechauth1/store"
	"time"
)

type ForgetPasswordServiceImpl struct {
	EmailService
	serviceUser *UserServiceImpl
}

func NewForgetPasswordServiceImpl(repoUser repositories.UserRepository) *ForgetPasswordServiceImpl {
	var userService = NewUserServiceImpl(repoUser)
	var emailService = communitions.NewCommunicationEmailService()
	return &ForgetPasswordServiceImpl{
		EmailService: emailService,
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
	return f.serviceUser.ResetPassword(email, password)
}
func (f ForgetPasswordServiceImpl) GenerateOpt() (string, error) {
	var otp string
	var hub = store.GetStoreOtp()
	for {
		otp = hub.Generate()
		if !hub.InStore(otp) {
			break
		}
		<-time.After(2 * time.Second)
	}
	return otp, nil
}
