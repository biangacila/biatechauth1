package services

import (
	"github.com/biangacila/biatechauth1/domain/aggregates"
	"github.com/biangacila/biatechauth1/domain/entities"
	"github.com/biangacila/biatechauth1/domain/repositories"
	"time"
)

type LoginServiceImpl struct {
	GenericServiceImpl[entities.Login]
	repo        repositories.LoginRepository
	serviceUser *UserServiceImpl
}

func NewLoginServiceImpl(repo repositories.LoginRepository, repoUser repositories.UserRepository) *LoginServiceImpl {
	var genericService = NewGenericServiceImpl(repo, entities.Login{})
	var userService = NewUserServiceImpl(repoUser)
	return &LoginServiceImpl{
		GenericServiceImpl: *genericService, // Use a pointer receiver for GenericServiceImpl
		repo:               repo,
		serviceUser:        userService,
	}
}
func (l LoginServiceImpl) NewLogin(username, password string) (user entities.User, token string, err error) {
	agg := aggregates.NewLoginAggregate()
	if err = agg.ValidUsernamePassword(username, password); err != nil {
		return user, token, err
	}

	user, err = l.serviceUser.UserExists(username)
	if err != nil {
		return user, token, err
	}
	if err = agg.Matches(user.Password, password); err != nil {
		return user, token, err
	}
	//let hid our password
	user.Password = ""

	token, _, err = agg.CreateToken(user.Email, user.GivenName, user.FamilyName, user.Phone)
	if err != nil {
		return entities.User{}, "", err
	}

	// let get our expired date time of token and aggregate our login for saving
	expiredAt, _, _ := l.IsValueToken(token)
	login, err := agg.New(username, token, expiredAt)
	if err != nil {
		return entities.User{}, "", err
	}

	// let save login now into db
	if _, err = l.repo.New(login); err != nil {
		return entities.User{}, "", err
	}
	return user, token, nil
}
func (l LoginServiceImpl) HasLogin(username string) (time.Time, bool, error) {
	agg := aggregates.NewLoginAggregate()
	login, err := l.repo.Find("logins", map[string]interface{}{"Username": username}, entities.Login{})
	if err != nil {
		return time.Time{}, false, err
	}
	return agg.GetTokenExpiryAndValidity(login.SignedToken)
}
func (l LoginServiceImpl) IsValueToken(token string) (time.Time, bool, error) {
	agg := aggregates.NewLoginAggregate()
	return agg.GetTokenExpiryAndValidity(token)
}
